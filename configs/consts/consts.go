package consts

import "time"

// 基础设施配置
const (
	EtcdAddress     = "127.0.0.1:2379"
	NatsAddress     = "nats://127.0.0.1:4222"
	KafkaBroker1Url = "localhost:9091"
	KafkaBroker2Url = "localhost:9092"
	KafkaBroker3Url = "localhost:9093"
	MySQLDefaultDSN = "TTMS:TTMS@tcp(localhost:3307)/TTMS?charset=utf8mb4&parseTime=True"
	RedisAddress    = "localhost:6378"
	RedisPassword   = "redis"
	RedisDB         = 0
	RedisTicketDB   = 1
	WebServerPort   = "8080"

	UserServiceName   = "userSvr"
	StudioServiceName = "studioSvr"
	PlayServiceName   = "playSvr"
	TicketServiceName = "ticketSvr"
	OrderServiceName  = "orderSvr"
)

const (
	TicketGroupID      = "ticket"
	OrderGroupID       = "order"
	TicketTimeoutTopic = "ticket-timeout"
	OrderBuyTopic      = "order-buy"
)

// 超时配置
const (
	TicketCacheTime  = time.Hour * 24   //票缓存时间
	OrderDelayTime   = time.Minute * 30 //订单超时时间
	RedisLockTimeOut = time.Second * 5  //redis锁超时时间

	RPCTimeout     = 10 * time.Second
	ConnectTimeout = 5 * time.Second

	JWTSecret = "kangning"
	// JWTOverTime 超长时间测试
	JWTOverTime = time.Hour * 72000
)

// 令牌桶限流配置
const (
	LimitRate = 2000.0 //令牌桶限流速率
	LimitCap  = 3000   //令牌桶容量
)
