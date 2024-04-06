package main

import (
	"context"
	"fmt"
	"kv-cache/pkg/kitex_gen/node"
	"kv-cache/pkg/kitex_gen/node/cacheservice"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	kitex_client cacheservice.Client
)

func main() {
	kitex_init()

	setGroupResp, err := kitex_client.SetGroup(context.Background(), &node.SetGroupRequest{Group: "test", MaxBytes: 1024})
	if err != nil {
		panic(err)
	}
	if setGroupResp.Code == 1 {
		panic(fmt.Errorf("invalid set group handle"))
	}

	setKeyResp, err := kitex_client.SetKey(context.Background(), &node.SetKeyRequest{Key: "1", Value: "a", Group: "test"})
	if err != nil {
		panic(err)
	}
	if setKeyResp.Code == 1 {
		panic(fmt.Errorf("invalid set key handle"))
	}

	forever := make(chan bool)
	average_time := make(chan int64, 1000)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			start := time.Now()
			getKeyResp, err := kitex_client.GetKey(context.Background(), &node.GetKeyRequest{Key: "1", Group: "test"})
			if err != nil {
				panic(err)
			}
			if getKeyResp.Code == 1 {
				panic(fmt.Errorf("invalid get key handle"))
			}
			end := time.Now()
			latency := end.Sub(start).Microseconds()
			average_time <- latency
			wg.Done()
		}()
	}
	wg.Wait()

	if len(average_time) != 1000 {
		panic(fmt.Errorf("invalid add"))
	}

	var sum int64 = 0
	for len(average_time) != 0 {
		sum += <-average_time
	}
	fmt.Printf("time cost=%d\n", sum/1000)

	<-forever
}

func kitex_init() {
	var err error
	etcdClient, err = etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	kitex_client, err = cacheservice.NewClient("node1",
		client.WithResolver(etcdClient),
		client.WithRPCTimeout(2*time.Second),
		client.WithConnectTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
}
