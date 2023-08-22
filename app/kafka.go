package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// the productTopic, sourceTopic and broker address are initialized as constants
const (
	productTopic = "products"
	broker       = "localhost:29092"
)

func createWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(broker),
		Topic:                  productTopic,
		AllowAutoTopicCreation: true,
	}
	return w
}

func sendProduct(ctx context.Context, w *kafka.Writer, p BaseProduct) {
	log.Println(p.ID)
	serializedProduct, _ := json.Marshal(p)
	message := kafka.Message{
		Key:   []byte(fmt.Sprintf("%s:%d", p.Source, p.ID)),
		Value: serializedProduct,
	}

	err := w.WriteMessages(ctx, message)
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func sendMessage(ctx context.Context, w *kafka.Writer, key string, data []byte) {
	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	err := w.WriteMessages(ctx, message)
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
