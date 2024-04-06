package process

import "kv-cache/pkg/consist"

var (
	M     *consist.Map
	Group = make([]string, 0, 128)
)

func Init(replicas int) {
	M = consist.New(replicas, nil)
	M.Set("node1", "node2")
}
