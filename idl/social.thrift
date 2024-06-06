namespace go social

include "user.thrift"

struct BaseResp {
    1: i64 code
    2: string msg
}

struct FollowRequest {
    1: string user_id
    2: string to_user_id
    3: string action_type
}

struct FollowResponse {
    1: BaseResp baseResp
}

struct FollowListRequest {
    1: string user_id
    2: i64 page_num
    3: i64 page_size
}

struct FollowListResponse {
    1: BaseResp baseResp
    2: list<user.User> users
}

struct FansListRequest {
    1: string uid
    2: i64 page_num
    3: i64 page_size
}

struct FansListResponse {
    1: BaseResp baseResp
    2: list<user.User> users
}

struct FriendsListRequest {
    1: i64 page_num
    2: i64 page_size
    3: string uid
}

struct FriendsListResponse {
    1: BaseResp baseResp
    2: list<user.User> users
}

service SocialService {
    FollowResponse Follow(1: FollowRequest req)
    FollowListResponse FollowList(1: FollowListRequest req)
    FansListResponse FansList(1: FansListRequest req)
    FriendsListResponse FriendsList(1: FriendsListRequest req)
}