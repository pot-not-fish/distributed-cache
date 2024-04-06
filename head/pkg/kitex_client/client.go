package kitex_client

import (
	"kv-cache/pkg/kitex_gen/node/cacheservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdClient discovery.Resolver

	Node1Client cacheservice.Client

	Node2Client cacheservice.Client

	once sync.Once
)

func Init() {
	var err error

	once.Do(func() {
		etcdClient, err = etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
		if err != nil {
			panic(err)
		}

		Node1Client, err = cacheservice.NewClient("node1", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}

		Node2Client, err = cacheservice.NewClient("node2", client.WithResolver(etcdClient))
		if err != nil {
			panic(err)
		}
	})
}
