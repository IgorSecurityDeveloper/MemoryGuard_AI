package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryEvent struct {
	Host              string  `json:"host"`
	TotalMemoryMb     uint64  `json:"totalMemoryMb"`
	UsedMemoryMb      uint64  `json:"usedMemoryMb"`
	AvailableMemoryMb uint64  `json:"availableMemoryMb"`
	UsedMemoryPercent float64 `json:"usedMemoryPercent"`
	Timestamp         string  `json:"timestamp"`
}

func main() {
	fmt.Println("MemoryGuard Collector iniciado")

	kafkaAddress := getEnv("KAFKA_BOOTSTRAP_SERVERS", "kafka:9092")
	topic := getEnv("KAFKA_TOPIC", "memory-events")
	host := getEnv("HOSTNAME", "unknown-host")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
	})

	defer writer.Close()

	for {
		virtualMemory, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Erro ao coletar memória:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		event := MemoryEvent{
			Host:              host,
			TotalMemoryMb:     virtualMemory.Total / 1024 / 1024,
			UsedMemoryMb:      virtualMemory.Used / 1024 / 1024,
			AvailableMemoryMb: virtualMemory.Available / 1024 / 1024,
			UsedMemoryPercent: virtualMemory.UsedPercent,
			Timestamp:         time.Now().UTC().Format(time.RFC3339),
		}

		payload, err := json.Marshal(event)
		if err != nil {
			fmt.Println("Erro ao converter evento para JSON:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(host),
			Value: payload,
		})

		if err != nil {
			fmt.Println("Erro ao enviar evento para Kafka:", err)
		} else {
			fmt.Println("Evento enviado para Kafka:", string(payload))
		}

		time.Sleep(10 * time.Second)
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
