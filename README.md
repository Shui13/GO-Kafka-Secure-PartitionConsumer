# GO-Kafka-Secure-PartitionConsumer
Secure Kafka partition consumer in Go lang

create topic 
./kafka-topics.sh --create --zookeeper <ip>:2181 --replication-factor 1 --partitions 10 --topic t3

List topics
./kafka-topics.sh --zookeeper <ip>:2181 --list

Produce message to kafka topic
./kafka-console-producer.sh --broker-list <ip>:9093 --topic t3 --producer.config ssl.properties

Consume partition 
./kafka-console-consumer.sh --bootstrap-server <ip>:9093 --topic t3 --new-consumer --consumer.config ssl.properties
