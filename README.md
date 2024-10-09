## 架构设计t

该系统采用微服务架构，拆分为多个服务，包括：

- `user` 用户服务，负责登陆注册以及用户信息
- `studio` 剧院服务，负责演出厅、座位等信息的管理
- `play` 剧目服务，负责电影和演出计划等信息的管理
- `ticket` 票务服务，负责售票和退票等
- `order` 订单服务，负责订单以及票房统计的管理和更新等

## 技术栈选择

- Gin：一个基于Go语言的Web框架，提供了快速构建API的功能。
- GORM：一个ORM库，提供了对数据库的操作。
- Kitex：一个高性能RPC框架，适合用于微服务架构。

## 开发步骤

1. 定义API接口和数据结构：根据需求，设计系统的API接口和数据结构，包括但不限于用户、剧院、演出、演员、订单、库存等。

2. 设计数据库模型：根据API接口和数据结构，设计数据库模型，选择合适的数据库，如MySQL。

3. 编写服务代码：根据API接口和数据库模型，编写服务代码，分别实现用户服务、剧院服务、订单服务、库存服务等。

4. 实现消息传递：通过消息服务，实现各个服务之间的通信和消息传递。

5. 部署和测试：将各个服务部署到不同的服务器上，并进行测试和调试。

## 目录结构设计

- configs：用于存放配置文件，如数据库配置、日志配置等。
- api:放置openApi文档以及.proto文件。
- internal：用于存放项目的内部代码，包括但不限于服务代码、数据库操作代码、工具函数等。
    - api：用于存放API接口相关的代码，如路由处理函数、参数校验等。
    - dao：用于存放数据访问对象相关的代码，如数据库模型定义、数据查询等。
    - service：用于存放服务相关的代码，如用户服务、剧院服务、订单服务、库存服务等。
- kitex_gen: 通过IDL文件生成的kitex库代码。
- scripts：用于存放项目的脚本文件。
- pkg：用于存放项目的公共库和工具函数等。

## 部署（单机模拟分布式环境）

- 首先确认go版本为1.16～1.20
- cd 到TTMS主目录下，执行`docker-compose up -d`，一键部署
- 在云服务器上部署时，需要用nginx从默认的80端口转发到web网关层的8080端口，并解决跨域问题，nginx配置在`configs/conf.d/default.conf`

```
location /ttms/ {
   	proxy_pass http://127.0.0.1:8080/ttms/;
	add_header Access-Control-Allow-Methods *;
        add_header Access-Control-Allow-Origin $http_origin always;
        add_header Access-Control-Allow-Max-Age 3600;
        add_header Access-Control-Allow-Credentials true;
        add_header Access-Control-Allow-Headers $http_access_control_request_headers;
        if ($request_method = OPTIONS) {
            return 200;
        }
	proxy_set_header Host $host;
}
```
- 前端压缩包有两个，都在script目录下，`build1.zip`是管理员界面，`build0.zip`是用户界面。解压之后，放在nginx配置对应的目录下就好
  - 用户前端放在 `/var/www/build0` ，系统前端放在 `/var/www/build1`
  - 管理员前端界面网址`http://localhost:81/`
  - 用户前端界面网址`http://localhost:82/`


## 运行（单机模拟分布式环境）

- cd 到script目录下
- 执行 `bash ./build.sh`（构建）
- 执行 `bash ./bootstrap.sh`（启动）

## 终止（单机模拟分布式环境）

- cd 到script目录下
- 执行 `bash ./stop.sh`

## 性能优化

- Kitex:
  - client端引入长连接以及连接池
  - client和server端同时开启多路复用
  - server端调高连接限制以及每秒请求数限制

- MySQL：
  - 增加 `max_connections`的值，之前默认151,增加至1000,以应对较高并发量
  - 为买票和订单服务数据库表中添加索引
    - ticket_index ON `tickets`(`schedule_id`,`seat_row`,`seat_col`)
    - order_index ON `orders`(`user_id`,`schedule_id`,`seat_row`,`seat_col`)
  - 优化部分SQL,对于判断是否存在的查找语句，一律返回`id`，避免回表

- 优化后情况：
  - 单组30个请求，分别模拟多用户抢票，抢票完成之后，再次发送退票请求
  - 开启200个线程进行压测，循环300次，共计180 0000次请求
  - 单机能承载大约3000+QPS(本机同时运行服务器及压测应用，服务器配置大约为4核8G8线程)
