package consist

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	nodes    []int
	hashMap  map[int]string
}

func New(replicas int, hash Hash) *Map {
	if hash == nil {
		hash = crc32.ChecksumIEEE
	}
	return &Map{
		hash:     hash,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
}

func (m *Map) Set(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i <= m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.nodes = append(m.nodes, hash)
			m.hashMap[hash] = node
		}
	}
	sort.Ints(m.nodes)
}

func (m *Map) Get(key string) (node string) {
	if key == "" {
		return node
	}

	hash := int(m.hash([]byte(key)))
	i := sort.Search(len(m.nodes), func(i int) bool {
		return m.nodes[i] >= hash
	})
	return m.hashMap[m.nodes[i%len(m.nodes)]]
}
