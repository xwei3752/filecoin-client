# filecoin-client

需要自行部署Lotus Node节点：https://lotu.sh/en+getting-started

此库仅添加部分方法，但已经满足钱包充值提现逻辑了，如果需要其他方法，请Fork后自行添加。

充值流程：获取头部高度，从本地高度遍历到头部高度，再根据高度获取区块CID，根据区块CID获取区块的所有消息，判断消息的类型是否0(0为发送Fil)，和接收地址是否是本地的地址。

说明请查询client_test文件。


# 安装

`go get github.com/myxtype/filecoin-client`

# 使用

```go
package main

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
)

func main() {

	client := filecoin.NewClient("http://127.0.0.1:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.cF__3r_0IR9KwZ2nLkqcBW8vuPePruZieJAVvTAoUA4")

	addr, _ := address.NewFromString("t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	b, err := client.WalletBalance(context.Background(), addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}
```

# 离线签名版

我将在后续提供离线签名版（不需要自行搭建Node），因为现在https://infura.io 还未开放Lotus Node的API，所以离线签名版将在后续更新。

具体实现逻辑在官方库中：https://github.com/filecoin-project/go-filecoin

离线签名版我尝试了很多次，但对于交叉编译的情况下不通过，主要是因为这个库`github.com/ipsn/go-secp256k1`

就很郁闷，因为我们的项目太依赖交叉编译了，一般都是直接在开发机编译好，上传到服务器执行。

提示这个错误：`go build github.com/ipsn/go-secp256k1: build constraints exclude all Go files in Path`

有人知道怎么弄吗？

# Lotus文档

https://lotu.sh/en+api-methods
