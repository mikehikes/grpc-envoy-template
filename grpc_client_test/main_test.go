package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"testing"

	pb "grpc_client_test/template_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPing(t *testing.T) {

	t.Log("Testing PING")

	serverAddr := "localhost:9090"

	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewTemplateServiceClient(conn)

	resp, err := client.Ping(context.Background(), &pb.PingRequest{})

	if err != nil {
		log.Fatalf("fail to ping: %v", err)
	}

	log.Printf("Pong: %s", resp)

	log.Println("PingTest completed")

	defer conn.Close()

}

func TestHelloWorld(t *testing.T) {

	t.Log("Testing HelloWorld")

	testMessage := "Hey there, from Go!"

	serverAddr := "localhost:9090"

	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewTemplateServiceClient(conn)

	t.Log("Sending ", testMessage)

	resp, err := client.HelloWorld(context.Background(), &pb.HelloWorldRequest{
		Name: testMessage,
	})

	if err != nil {
		log.Fatalf("fail to send hello world: %v", err)
	}

	log.Printf("Received : %s", resp)

	t.Log("does it match the expected message? ", testMessage)
	responseText := strings.TrimSpace(resp.Message)
	originalText := strings.TrimSpace(testMessage)
	stringCompare := strings.Compare(responseText, "Hello "+originalText)

	if stringCompare != 0 {
		t.Errorf("Expected %s, got %s", testMessage, resp.Message)
	}

	log.Println("message matches! HelloWorld completed")

	defer conn.Close()

}
