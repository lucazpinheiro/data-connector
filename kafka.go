package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// the productTopic, sourceTopic and broker address are initialized as constants
const (
	sourceTopic  = "sources"
	productTopic = "products"
	broker       = "localhost:29092"
)

func createReader() *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   sourceTopic,
		GroupID: "consumer-group-id",
	})
	return r
}

func readSource(ctx context.Context, r *kafka.Reader, ch chan ImportMessage) {
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}

		var sourcecMessage ImportMessage
		json.Unmarshal(m.Value, &sourcecMessage)
		if err != nil {
			panic("could not unmarshal message " + err.Error())
		}

		ch <- sourcecMessage

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

func createWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(broker),
		Topic:                  productTopic,
		AllowAutoTopicCreation: true,
	}
	return w
}

func sendProduct(ctx context.Context, w *kafka.Writer, products []BaseProduct) {
	var messages = make([]kafka.Message, len(products))
	for i, p := range products {
		serializedProduct, _ := json.Marshal(p)
		messages[i] = kafka.Message{
			Key:   []byte(fmt.Sprintf("%s:%d", p.Source, p.ID)),
			Value: serializedProduct,
		}
	}

	err := w.WriteMessages(ctx, messages...)
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
