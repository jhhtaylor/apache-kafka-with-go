# apache-kafka-with-go

# Prerequisite

Download Kafka from https://www.apache.org/dyn/closer.cgi?path=/kafka/3.5.0/kafka_2.13-3.5.0.tgz

Download kafka-go
```bash
go get github.com/segmentio/kafka-go
```

# Setup

In two different terminal tabs

```bash
cd Downloads/kafka_2.13-3.5.0
```

and then in each

```bash
bin/zookeeper-server-start.sh config/zookeeper.properties
```

```bash
bin/kafka-server-start.sh config/server.properties
```

## create a topic to store your events

(you only need to do this the first time you set up a new topic)

```bash
bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092
```

# Demos

## quickstart producer console app

```bash
bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092
```

run this with 

```bash
go run consumer.go 
```

in `../apache-kafka-with-go/consumer`

## quickstart consumer console app

```bash
bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092
```

run this with 

```bash
go run producer.go 
```

in `../apache-kafka-with-go/producer`

## quickstart producer and consumer console app

just run both bash scripts above in different terminals

You can also have all 4 running at the same time, it’s okay

# Resources

The tutorial came from this video: https://youtu.be/O69sz842sFs

Guide and commands came from here: https://kafka.apache.org/quickstart

kafka-go package: https://github.com/segmentio/kafka-go