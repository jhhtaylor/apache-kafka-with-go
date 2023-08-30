# **apache-kafka-with-go**

This repository contains a demonstration of using Apache Kafka with Go.

## **Prerequisites**

1. **Download Apache Kafka**: Fetch Kafka from **[here](https://www.apache.org/dyn/closer.cgi?path=/kafka/3.5.0/kafka_2.13-3.5.0.tgz)**.
2. **Download kafka-go**:
    
    ```bash
    go get github.com/segmentio/kafka-go
    ```
    

## **Setup**

### **Starting Kafka Servers**

1. Open your terminal and navigate to the Kafka directory:
    
    ```bash
    cd Downloads/kafka_2.13-3.5.0
    ```
    
2. **Terminal Tab 1**: Start the Zookeeper server:
    
    ```bash
    bin/zookeeper-server-start.sh config/zookeeper.properties
    ```
    
3. **Terminal Tab 2**: Start the Kafka server:
    
    ```bash
    bin/kafka-server-start.sh config/server.properties
    ```
    

### **Creating a Topic**

1. **Terminal Tab 3**: Create a topic named **`quickstart-events`** (Only the first time):
    
    ```bash
    bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092
    ```
    

## **Demos**

### **1. Quickstart Producer Console App**

- **Terminal Tab 4**: Start the Kafka console producer:
    
    ```bash
    bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092
    ```
    
- **Terminal Tab 5**: Navigate to the **`consumer`** directory and run the Go consumer:
    
    ```bash
    cd ../apache-kafka-with-go/consumer
    go run consumer.go
    ```
    

### **2. Quickstart Consumer Console App**

- **Terminal Tab 4**: Start the Kafka console consumer:
    
    ```bash
    bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092
    ```
    
- **Terminal Tab 5**: Navigate to the **`producer`** directory and run the Go producer:
    
    ```bash
    cd ../apache-kafka-with-go/producer
    go run producer.go
    ```
    

### **3. Quickstart Producer and Consumer Console App**

You can have both the producer and consumer bash scripts (from **Terminal Tab 4**) running simultaneously in different terminals.

*Note:* It's perfectly fine to have all the terminals (Tabs 1, 2, and both from Tab 4) running concurrently.

## **Resources**

- **Tutorial Video**: **[YouTube Link](https://youtu.be/O69sz842sFs)**
- **Official Kafka Quickstart Guide**: **[Apache Kafka Documentation](https://kafka.apache.org/quickstart)**
- **kafka-go Package**: **[GitHub Repository](https://github.com/segmentio/kafka-go)**