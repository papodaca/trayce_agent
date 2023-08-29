package main

import "C"
import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/evanrolfe/dockerdog/api"
	"github.com/evanrolfe/dockerdog/internal"
	"github.com/evanrolfe/dockerdog/internal/sockets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	bpfFilePath       = "bundle/ssl.bpf.o"
	btfFilePath       = "bundle/6.2.0-26-generic.btf"
	sslLibDefault     = "/usr/lib/x86_64-linux-gnu/libssl.so.3"
	grpcServerDefault = "localhost:50051"
)

func extractFile(data []byte, destPath string) {
	f, err := os.Create(destPath)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}

	f.Close()
}

func main() {
	// Parse Command line args
	var pid int
	var libSslPath, grpcServerAddr string
	flag.IntVar(&pid, "pid", 0, "The PID of the docker container to instrument. Or 0 to intsrument this container.")
	flag.StringVar(&libSslPath, "libssl", sslLibDefault, "The path to the libssl shared object.")
	flag.StringVar(&grpcServerAddr, "grpcaddr", grpcServerDefault, "The address of the GRPC server to send observations to.")
	flag.Parse()

	fmt.Println("PID: ", pid)
	fmt.Println("libssl: ", libSslPath)

	// Extract bundled files
	bpfBytes := internal.MustAsset(bpfFilePath)
	btfBytes := internal.MustAsset(btfFilePath)
	btfDestFile := "./5.8.0-23-generic.btf"
	extractFile(btfBytes, btfDestFile)

	// Start the agent
	agent := internal.NewBPFAgent(bpfBytes, btfFilePath, libSslPath)
	defer agent.Close()

	// Create a channel to receive interrupt signals
	interrupt := make(chan os.Signal, 1)
	socketMsgChan := make(chan sockets.SocketMsg)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Agent listing...")
	go agent.ListenForEvents(socketMsgChan)

	// Start a goroutine to handle the interrupt signal
	var wg sync.WaitGroup
	wg.Add(1)

	// API Client
	// Set up a connection to the server.
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewDockerDogAgentClient(conn)

	go func() {
		for {
			// Check if the interrupt signal has been received
			select {
			case <-interrupt:
				wg.Done()
				return
			case socketMsg := <-socketMsgChan:
				fmt.Printf("[MsgEvent] %s - Local: %s, Remote: %s\n", "", socketMsg.LocalAddr, socketMsg.RemoteAddr)
				socketMsg.Debug()

				// Contact the server and print out its response.
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				_, err := grpcClient.SendRequestObserved(ctx, &pb.RequestObserved{Method: "GET", Url: socketMsg.RemoteAddr})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
			}
		}
	}()

	// For testing purposes:
	// cmd := exec.Command("curl", "--parallel", "--parallel-immediate", "--config", "/app/urls.txt", "--http1.1")
	// cmd.Output()
	cmd := exec.Command("ruby", "tmp/request.rb")
	cmd.Output()

	wg.Wait()

	fmt.Println("Done, closing agent.")
	os.Remove(btfDestFile)

	// agent.Close()
}
