package jwt

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/user"
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var MySecret = []byte(consts.JWTSecret)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddress,
		Password: consts.RedisPassword,
		DB:       consts.RedisDB, // use default DB
	})
}

func GenToken(userInfo *user.User) (string, error) {
	c := MyClaims{
		userInfo.Id,
		userInfo.Type,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(consts.JWTOverTime)),
			Issuer:    "kangning",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完成的编码后的字符串token
	tokenStr, err := token.SignedString(MySecret)
	if err != nil {
		log.Panicln(err)
	}
	//从黑名单中取出token,该想法不合适，已禁止
	//num, err1 := Client.HDel(context.Background(), "black", strconv.Itoa(int(userInfo.Id))).Result()
	//log.Println("Client.LRem return ", num, err1)
	return tokenStr, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i any, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		//判断用户是否在黑名单中
		t, _ := Client.HGet(context.Background(), "black", strconv.Itoa(int(claims.ID))).Result()
		if t == tokenString {
			return nil, errors.New("更改密码后需重新登录")
		}
		return claims, nil
	}
	return nil, errors.New("JWT 解析错误")
}

// DiscardToken 废弃JWT(更改密码后)
func DiscardToken(id int, tokenString string) {
	//用户更改密码之后强制用户重新登录-->右边为之前不太恰当的一个想法，原因请查看两行后的注释 （此时，将用户添加到黑名单中，用户重新登陆之后，把用户从黑名单中移除）
	_, err := Client.HSet(context.Background(), "black", strconv.Itoa(id), tokenString).Result()
	//为黑名单添加过期时间，因为用户重新登录之后，之前的token还不能从黑名单中取出来，否则之前过期的token就又可以重新使用了（这种情况不允许）
	err = Client.Expire(context.Background(), "black", consts.JWTOverTime).Err()
	if err != nil {
		log.Println(err)
	}
}

// SetVerification 生成验证码之后加入redis
func SetVerification(email, verification string) {
	err := Client.SetNX(context.Background(), email, verification, 5*time.Minute).Err()

	if err != nil {
		log.Println(err)
	}
}

// CheckVerification 验证失败时返回error
func CheckVerification(email, verification string) error {
	v, err := Client.Get(context.Background(), email).Result()
	if err != nil || v != verification {
		log.Println(err)
		return errors.New("验证码无效")
	}
	err = Client.Del(context.Background(), email).Err()
	if err != nil {
		log.Println(err)
	}
	return nil
}
