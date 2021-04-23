# opensearch
Go SDK for Aliyun Services opensearch

## 快速使用
```go
package main

import (
	"fmt"

	"github.com/Excalibur-1/opensearch"
)

const (
	appName         = "<appName>"     // 可以是app名称，也可以是app的版本号
	networkType     = "<networkType>" // 如果没有的话、可以为空
	accessKeyId     = "<AccessKeyId>"
	accessKeySecret = "<AccessKeySecret>"
)

func main() {
	opClient := opensearch.NewClient(networkType, accessKeyId, accessKeySecret, opensearch.Hangzhou)
	args := opensearch.ScrollArgs{
		Args: opensearch.Args{
			AppName:     appName,
			Query:       `config=start:0,hit:10,format:fulljson,rerank_size:200&&query=title:"搜索"`,
		},
		Scroll:     "1m",
		SearchType: "scan",
		ScrollId:   "",
	}
	resp, err := opClient.Scroll(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

	// 第二次查询只需要设置appName、scrollId,且不能设置其他参数，否则会报签名失败
	args.ScrollId = resp.Result.ScrollId
	resp1, err := opClient.Scroll(opensearch.ScrollArgs{
		Args:     opensearch.Args{AppName: appName},
		ScrollId: resp.Result.ScrollId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp1)
}

```