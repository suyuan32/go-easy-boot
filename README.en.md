# go-easy-boot

#### Description

This is a management system based on go-zero and support microservice development.

#### Software Architecture

Software architecture description

#### Installation

## Rpc

- Create proto file in the proto directory like below

```proto
syntax = "proto3";

package user;

option go_package = "./;proto";

import "base.proto";

message LoginReq {
  string username = 1;
  string password = 2;
}

message LoginResp {
  int64  id = 1;
  int32  role_id = 2;
  int32  status = 3;
  string username = 4;
  string access_token = 5;
  uint64 access_expire = 6;
}

message RegisterReq {
  string user_name = 1;
  string password = 2;
  string email = 3;
}


message ChangePasswordReq {
  uint64 user_id = 1;
  string old_password = 2;
  string new_password = 3;
}

service User {
  rpc changePassword(ChangePasswordReq) returns (BaseResp);
  rpc createUser(RegisterReq) returns (BaseResp);
}

```

- Use goctl command

```shell
 goctl rpc protoc user.proto --proto_path={proto file absolute path} --go_out=../types --go-grpc_out=../types --zrpc_out=../
```

#### Instructions

1. xxxx
2. xxxx
3. xxxx

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request

