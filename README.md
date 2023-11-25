README
=====

# 安装需要的依赖包
```golang
网络不好，可能需要安装很长时间。
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.io

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// go install github.com/envoyproxy/protoc-gen-validate@latest
$ go install github.com/bufbuild/protoc-gen-validate@latest
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

$ cd $GOPATH/bin
$ ls -l
total 103424
-rwxr-xr-x  1 yuyongpeng  staff   8.6M 11 24 23:19 protoc-gen-go
-rwxr-xr-x  1 yuyongpeng  staff   8.2M 11 24 23:43 protoc-gen-go-grpc
-rwxr-xr-x  1 yuyongpeng  staff    11M 11 25 00:36 protoc-gen-grpc-gateway
-rwxr-xr-x  1 yuyongpeng  staff    11M 11 24 23:04 protoc-gen-openapiv2
-rwxr-xr-x  1 yuyongpeng  staff    12M 11 25 00:34 protoc-gen-validate
```

## pb项目介绍

按pb目录名称首字母排序:

	├── README.md   # <- @me 
	├── adminpb     # 管理应用pb
	├── basepb      # 基础应用pb
	├── bass        # 上链应用pb
	├── dassetpb    # 藏品应用pb
	├── dscreenpb   # 投屏应用pb
	├── merchantpb  # 商户应用pb
	├── paymentpb   # 支付应用pb
	├── rbacpb      # RBAC应用pb
	└── user        # 用户应用pb

## 开发环境约定

go version = 1.20.1

## buf构建工具

安装

```bash
BIN="/usr/local/bin" && \
VERSION="1.9.0" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"
```

`buf`使用详见[buf官方文档][buf-offical-doc]

buf使用

```bash
# 构建准备
make buf-init

# 构建paymentpb
make proto-payment
```

[buf-offical-doc]: https://docs.buf.build/
