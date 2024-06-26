# distributed-cache

### 目录结构
```
kv-cache
|-- concurrent  测试并发时RPC调用缓存的次数，以及响应的速度
|-- head        HTTP服务，基于Gin搭建，用于后面可视化界面
|-- idl
|-- node        各个缓存节点RPC服务，基于Kitex搭建
|-- pkg
	|-- cache_algorithm   提供存储算法接口
	|-- consist           哈希一致算法
	|-- group             隔离缓存
	|-- kitex_gen
	|-- mutex             并发读写安全
	|-- singleflight      防止缓存击穿
```

### 并发测试结果
创建1000条协程同时请求node节点的同一个数据，得到的平均响应为256.85毫秒
![Pasted image 20240406215855.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c63223024a8746d5992905ea6f353c38~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=566&h=59&s=5837&e=png&b=1e1e1e)
<br>
1000条请求中，有242次请求缓存
<br>
![Pasted image 20240406215912.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/cc8d191506794cfe812e5fafc59b9588~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=498&h=65&s=5975&e=png&b=1e1e1e)
