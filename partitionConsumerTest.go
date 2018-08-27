package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
)

func main() {
	tlsConfig, err := NewTLSConfig("/Users/pathToFile.pem",
		"/Users/pathToFile.key",
		"/Users/pathToFile.pem")

	if err != nil {
		fmt.Print(err)
	}

	tlsConfig.InsecureSkipVerify = true

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig

  //list broker ips
	master, err := sarama.NewConsumer([]string{"<ip>:9093"}, config)

	if err != nil {
		fmt.Print("Error", err)
	}

	fmt.Print("establishing connection...\n")

  //master.ConsumePartition(topic, partition number, offset)
	consumer, err := master.ConsumePartition("t3", 3, -1)
	if err != nil {
		fmt.Print("Error: ", err)
	}

	fmt.Print("listening to messages...\n")

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Print("Kafka Error: %s", err.Error())
		case msg := <-consumer.Messages():
			fmt.Print("Data received: topic:", msg.Topic, ", partition:", msg.Partition, ", offset:", msg.Offset, ", key:", msg.Key, ", value:", string(msg.Value))
		}
	}
}

func NewTLSConfig(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return &tlsConfig, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool

	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, err
}
