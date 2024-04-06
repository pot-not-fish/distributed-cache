package handler

import (
	"context"
	"kv-cache/pkg/group"
	"kv-cache/pkg/kitex_gen/node"
)

type CacheServiceImpl struct{}

func (c *CacheServiceImpl) GetKey(ctx context.Context, request *node.GetKeyRequest) (*node.GetKeyResponse, error) {
	group, err := group.GetGroup(request.Group)
	if err != nil {
		return &node.GetKeyResponse{Code: 1}, nil
	}

	value, ok := group.Get(request.Key)
	if !ok {
		return &node.GetKeyResponse{Code: 1}, nil
	}

	return &node.GetKeyResponse{Code: 0, Value: string(value)}, nil
}

func (c *CacheServiceImpl) SetKey(ctx context.Context, request *node.SetKeyRequest) (*node.SetKeyResponse, error) {
	group, err := group.GetGroup(request.Group)
	if err != nil {
		return &node.SetKeyResponse{Code: 1}, nil
	}

	group.Set(request.Key, []byte(request.Value))

	return &node.SetKeyResponse{Code: 0}, nil
}

func (c *CacheServiceImpl) DelKey(ctx context.Context, request *node.DelKeyRequest) (*node.DelKeyResponse, error) {
	group, err := group.GetGroup(request.Group)
	if err != nil {
		return &node.DelKeyResponse{Code: 1}, nil
	}

	if err = group.Del(request.Group); err != nil {
		return &node.DelKeyResponse{Code: 1}, nil
	}

	return &node.DelKeyResponse{Code: 0}, nil
}

func (c *CacheServiceImpl) SetGroup(ctx context.Context, request *node.SetGroupRequest) (*node.SetGroupResponse, error) {
	_, err := group.NewGroup(request.Group, request.MaxBytes, nil, nil)
	if err != nil {
		return &node.SetGroupResponse{Code: 1}, nil
	}

	return &node.SetGroupResponse{Code: 0}, nil
}

func (c *CacheServiceImpl) GetAll(ctx context.Context, request *node.GetAllKeysRequest) (*node.GetAllKeysResponse, error) {
	group, err := group.GetGroup(request.Group)
	if err != nil {
		return &node.GetAllKeysResponse{Code: 1}, nil
	}

	keys_string, values_byte := group.GetAll()
	caches := make([]*node.Cache, 0, len(values_byte))
	for i := range values_byte {
		caches = append(caches, &node.Cache{Key: keys_string[i], Value: string(values_byte[i])})
	}
	return &node.GetAllKeysResponse{Code: 0, Caches: caches}, nil
}

func (c *CacheServiceImpl) GetAllGroup(ctx context.Context, request *node.GetAllGroupsRequest) (*node.GetAllGroupsResponse, error) {
	return &node.GetAllGroupsResponse{Code: 0, Groups: group.GetAllGroups()}, nil
}
