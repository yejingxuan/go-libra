package main

import (
	"context"
	"fmt"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/store/etcd"
	"time"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.Run(testETCD)
}

func testETCD() error {
	cli, _ := etcd.StdConfig().Build()
	defer cli.Close()
	// put
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	cli.Put(ctx, "/testdir/testkey2", "hello-world")

	// get
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	resp, _ := cli.Get(ctx, "/testdir/testkey2")

	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return nil
}
