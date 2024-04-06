namespace go node

struct GetKeyRequest {
    1: string key;
    2: string group;
}

struct GetKeyResponse {
    1: string value;
    2: i16 code;
}

struct SetKeyRequest {
    1: string key;
    2: string value;
    3: string group;
}

struct SetKeyResponse {
    1: i16 code;
}

struct DelKeyRequest {
    1: string key;
    2: string group;
}

struct DelKeyResponse {
    1: i16 code;
}

struct SetGroupRequest {
    1: string group;
    2: i64 max_bytes;
}

struct SetGroupResponse {
    1: i16 code;
}

struct GetAllKeysRequest {
    1: string group;
}

struct Cache {
    1: string key;
    2: string value;
}

struct GetAllKeysResponse {
    1: i16 code;
    3: list<Cache> caches;
}

struct GetAllGroupsRequest {
}

struct GetAllGroupsResponse {
    1: i16 Code;
    2: list<string> groups;
}

service CacheService {
    GetKeyResponse GetKey(1: GetKeyRequest request);

    SetKeyResponse SetKey(1: SetKeyRequest request);

    DelKeyResponse DelKey(1: DelKeyRequest request);

    SetGroupResponse SetGroup(1: SetGroupRequest request);

    GetAllKeysResponse GetAll(1: GetAllKeysRequest request);

    GetAllGroupsResponse GetAllGroup(1: GetAllGroupsRequest request);
}
