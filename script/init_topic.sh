docker exec -it kafka1 kafka-topics.sh --create --if-not-exists --zookeeper 172.17.0.1:2181/kafka --replication-factor 2 --partitions 3 --topic ticket-timeout
docker exec -it kafka1 kafka-topics.sh --create --if-not-exists --zookeeper 172.17.0.1:2181/kafka --replication-factor 2 --partitions 3 --topic order-buy