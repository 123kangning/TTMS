package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/redis"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func NewKafkaEventLoop() {
	go ticketTimeOut()
	logrus.Info("kafka event start")
}

func ticketTimeOut() {
	tag := "ticketTimeOut"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{consts.KafkaBroker1Url, consts.KafkaBroker2Url, consts.KafkaBroker3Url},
		GroupID: consts.TicketGroupID,
		Topic:   consts.TicketTimeoutTopic,
	})
	defer r.Close()

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			logrus.Error(tag, "failed to fetch message:", err)
			break
		}
		timeoutTicketHandler(m)
		if err := r.CommitMessages(ctx, m); err != nil {
			logrus.Error(tag, "failed to commit message:", err)
		}
	}
}

func timeoutTicketHandler(msg kafka.Message) {
	data := strings.Split(string(msg.Value), ";")
	fmt.Println(data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	redis.ReturnTicket(context.Background(), fmt.Sprintf("%d;%d;%d", d0, d1, d2))
	dao.ReturnTicket(context.Background(), int64(d0), int32(d1), int32(d2))
}
