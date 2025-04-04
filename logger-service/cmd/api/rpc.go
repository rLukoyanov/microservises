package main

import (
	"context"
	"log-service/data"
	"time"
)

type RPCServer struct {
}

type RPCPyaload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPyaload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
