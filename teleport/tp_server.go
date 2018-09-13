package main

import (
	"flag"
	// "net/http"
	// _ "net/http/pprof"
	"runtime"
	"time"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/proto/pbproto"
)

type Hello struct {
	tp.CallCtx
}

func (t *Hello) Say(args *BenchmarkMessage) (*BenchmarkMessage, *tp.Rerror) {
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
	port = flag.Int64("p", 8972, "listened port")
	// cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	delay = flag.Duration("delay", 0, "delay to mock business processing")
	// debugAddr  = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

func main() {
	flag.Parse()

	tp.SetDefaultProtoFunc(pbproto.NewPbProtoFunc)
	tp.SetLoggerLevel("ERROR")

	// go func() {
	// 	tp.Println(http.ListenAndServe(*debugAddr, nil))
	// }()

	server := tp.NewPeer(tp.PeerConfig{
		DefaultBodyCodec: "protobuf",
		ListenPort:       uint16(*port),
	})
	server.RouteCall(new(Hello))
	server.ListenAndServe()
}
