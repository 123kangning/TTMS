package mw

/*
+---------------------------------------------------------------------------+---------------------------------------------------------------+
|							BuyTicketMsg									|						ReturnTicketMsg							|
|---------------------------------------------------------------------------+---------------------------------------------------------------+
|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|	Price	|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|
+---------------------------------------------------------------------------+---------------------------------------------------------------+
*/
import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
	"strings"
	"time"
)

var nc *nats.Conn
var JS nats.JetStreamContext

func InitNats() {
	// 连接到nats的服务器
	var err error
	nc, err = nats.Connect(consts.NatsAddress)
	if err != nil {
		log.Panic("1", err)
	}

	// 初始化JetStream功能
	JS, err = nc.JetStream(nats.Context(context.Background()))
	if err != nil {
		log.Panic("2", err)
	}
	//streamName, subject, subject1 := "stream", "order.buy", "order.return"
	streamName, subject, subject2, subject3 := "stream", "order.buy", "ticket.commit", "ticket.timeout"
	stream, err := JS.StreamInfo(streamName)
	if err != nil {
		log.Println("3", err) // 如果不存在，这里会有报错
	}
	if stream == nil {
		//fmt.Sprintf("%s.%s", streamName, subject), fmt.Sprintf("%s.%s", streamName, subject1)
		log.Printf("creating stream %q and subject %q", streamName, subject)
		_, err = JS.AddStream(&nats.StreamConfig{
			Name: streamName,
			Subjects: []string{fmt.Sprintf("%s.%s", streamName, subject),
				fmt.Sprintf("%s.%s", streamName, subject2), fmt.Sprintf("%s.%s", streamName, subject3)},
			MaxAge: 3 * 24 * time.Hour,
		})
		if err != nil {
			log.Panicln("4", err)
		}
	}
	consumer, err := JS.AddConsumer(streamName, &nats.ConsumerConfig{Durable: "buyConsumer"})
	fmt.Println("consumer = ", consumer)
	sub, err := JS.PullSubscribe(fmt.Sprintf("%s.%s", streamName, subject), "buyConsumer") //买票

	for {
		msgs, _ := sub.Fetch(10)
		for _, msg := range msgs {
			switch msg.Subject {
			case "stream.order.buy":
				AddOrderHandler(msg)
			}
		}
	}
}

func AddOrderHandler(msg *nats.Msg) {
	data := strings.Split(string(msg.Data), ";")
	log.Println("data = ", data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	d3, _ := strconv.Atoi(data[3])
	d5, _ := strconv.Atoi(data[5])

	ctx := context.Background()
	oldTicketCount := dao.GetOrderCount(ctx, d1, d2, d3, []int{1, 2})
	if oldTicketCount > 0 {
		//终止此消息，再次保证了一张票只能被一个人买到
		err := msg.Term()
		if err != nil {
			log.Println(ctx, "Term error", err)
		}
		return
	}

	t := float64(time.Now().Add(consts.OrderDelayTime).Unix())
	//fmt.Println("time = ", time.Now().Format("2006-01-02 15:04:05"), " unix = ", time.Now().Unix())
	//fmt.Println("time = ", time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"), " unix = ", t)
	orderInfo := strings.Join(data[:4], ";")
	//fmt.Println("orderInfo = ", orderInfo)
	ToDelayQueue(ctx, orderInfo, t)
	//fmt.Println(d0, d1, d2, d3, data[4])
	err := dao.AddOrder(d0, d1, d2, d3, data[4], d5)
	//暂时未出现这种情况，先不管
	if err != nil {
		log.Println("dao.AddOrder failed,err=", err, "msg is", string(msg.Data))
	}

	//业务消费消息完成之后，再返回ack告知消息队列跳过该消息，获取下一条消息
	err = msg.Ack()
	if err != nil {
		log.Println(ctx, "ack error", err)
		return
	}
}
