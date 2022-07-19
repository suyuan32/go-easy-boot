# go-easy-boot

#### 介绍

基于go-zero开发的分布式后端管理系统

#### 框架

| 用途      | 框架    |
|---------|-------|
| ORM     | GROM  |
| HTTP 请求 | httpx |
|         |       |
|         |       |
|         |       |

#### 安装教程

1. xxxx
2. xxxx
3. xxxx

# Getting Start

## Rpc

- 首先创建proto文件夹并新建自定义proto文件，下面是例子

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

- 使用 goctl 命令生成对应的go文件

```shell
 goctl rpc protoc user.proto --proto_path={proto file absolute path} --go_out=../types --go-grpc_out=../types --zrpc_out=../
```

#### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request

