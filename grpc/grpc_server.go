package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Hello struct{}

func (t *Hello) Say(ctx context.Context, args *BenchmarkMessage) (reply *BenchmarkMessage, err error) {
	s := "OK"
	var i int32 = 100
	args.Field1 = s
	args.Field2 = i
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return args, nil
}

var (
	host      = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	delay     = flag.Duration("delay", 0, "delay to mock business processing")
	debugAddr = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterHelloServer(s, &Hello{})
	s.Serve(lis)
}
