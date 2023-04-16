package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/user"
	"TTMS/pkg/gmail"
	"TTMS/pkg/jwt"
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	req := &user.CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//_, err := jwt.ParseToken(req.Token)
	//if err != nil {
	//	c.JSON(http.StatusOK, user.CreateUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
	//	return
	//}
	fmt.Println(req)
	resp, err := rpc.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UserLogin(c *gin.Context) {
	req := &user.UserLoginRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	resp, err := rpc.UserLogin(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	if resp.BaseResp.StatusCode == 0 {
		Token, err := jwt.GenToken(resp.UserInfo)
		if err != nil {
			log.Println("JWT 生成错误", err)
			resp.BaseResp.StatusCode = 1
			resp.BaseResp.StatusMessage = err.Error()
		} else {
			resp.Token = Token
		}
	}
	resp.UserInfo = nil
	c.JSON(http.StatusOK, resp)
}

func GetAllUser(c *gin.Context) {
	req := &user.GetAllUserRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	token := c.Query("Token")
	//token验证
	_, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, user.GetAllUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	//接收resp
	resp, err := rpc.GetAllUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type ChangeUserPasswordRequest struct {
	UserId      int64
	Password    string
	NewPassword string
	Token       string
}

func ChangeUserPassword(c *gin.Context) {
	receive := &ChangeUserPasswordRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//token验证
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, user.ChangeUserPasswordResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	//接收resp
	req := &user.ChangeUserPasswordRequest{
		UserId:      receive.UserId,
		Password:    receive.Password,
		NewPassword: receive.NewPassword,
	}
	resp, err := rpc.ChangeUserPassword(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	jwt.DiscardToken(int(req.UserId), receive.Token)
	c.JSON(http.StatusOK, resp)
}

type DeleteUserRequest struct {
	UserId int64
	Token  string
}

func DeleteUser(c *gin.Context) {
	receive := &DeleteUserRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//token验证
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, user.DeleteUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	//接收resp
	req := &user.DeleteUserRequest{
		UserId: receive.UserId,
	}
	resp, err := rpc.DeleteUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetUserInfoRequest struct {
	Token string
}

func GetUserInfo(c *gin.Context) {
	receive := &GetUserInfoRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//token验证
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, user.GetUserInfoResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	//接收resp
	req := &user.GetUserInfoRequest{
		UserId: claim.ID,
	}
	resp, err := rpc.GetUserInfo(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

/*
GetVerification 两个优化提升了用户体验：
1.将发送邮件的操作异步进行，使前端用户界面能快速得到反馈
2.当两次请求发送邮箱验证码的时间间隔在 1～5 分钟时，服用上一次的邮箱验证码，同时更新redis中存储 <email,verification> 的过期时间
*/
func GetVerification(c *gin.Context) {
	req := &user.GetVerificationRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	mail := req.Email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// 编译正则表达式
	re := regexp.MustCompile(pattern)
	resp := &user.GetVerificationResponse{BaseResp: &user.BaseResp{}}
	if re.MatchString(mail) {
		t, _ := jwt.Client.TTL(context.Background(), mail).Result()
		if t < 0 { //过期，允许再次发送
			verification := gmail.GetVerification()
			go gmail.Send(mail, verification)
			jwt.SetVerification(mail, verification)
		} else if t < time.Minute*4 { //距离上次请求发送邮箱验证码已经过去了1分钟,但还不足5分钟，复用之前的验证码
			verification, _ := jwt.Client.Get(context.Background(), mail).Result()
			go gmail.Send(mail, verification)
			jwt.Client.Expire(context.Background(), mail, 5*time.Minute)
		} else { //距离上次请求发送邮箱验证码不足了1分钟,拒绝
			resp.BaseResp.StatusCode = 1
			resp.BaseResp.StatusMessage = "距离上次发送不足一分钟"
		}
	} else {
		fmt.Println("不是合法的邮箱地址")
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = "不是合法的邮箱地址"
	}
	c.JSON(http.StatusOK, resp)
}

type BindEmailRequest struct {
	Email        string
	Verification string
	Token        string
}

func BindEmail(c *gin.Context) {
	receive := &BindEmailRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//token验证
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, user.DeleteUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &user.BindEmailRequest{
		Email:        receive.Email,
		Verification: receive.Verification,
		UserId:       claim.ID,
	}
	var resp *user.BindEmailResponse
	err = jwt.CheckVerification(receive.Email, receive.Verification)
	if err != nil {
		resp = &user.BindEmailResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}}
	} else { //验证码匹配成功
		req.UserId = claim.ID
		resp, err = rpc.BindEmail(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
			log.Fatalln(err)
		}
	}
	c.JSON(http.StatusOK, resp)
}

func ForgetPassword(c *gin.Context) {
	req := &user.ForgetPasswordRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	var resp *user.ForgetPasswordResponse
	err := jwt.CheckVerification(req.Email, req.Verification)
	if err != nil {
		resp = &user.ForgetPasswordResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}}
	} else { //验证码匹配成功
		resp, err = rpc.ForgetPassword(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, err)
			log.Fatalln(err)
		}
	}
	c.JSON(http.StatusOK, resp)
}
