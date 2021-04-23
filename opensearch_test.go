package opensearch_test

import (
	"fmt"
	"github.com/Excalibur-1/opensearch"
	"testing"
)

const (
	appName         = "<AppName>"     // 可以是app名称，也可以是app的版本号
	networkType     = "<NetworkType>" // 如果没有的话、可以为空
	accessKeyId     = "<AccessKeyId>"
	accessKeySecret = "<AccessKeySecret>"
)

func TestOpClient_Scroll(t *testing.T) {
	opClient := opensearch.NewClient(networkType, accessKeyId, accessKeySecret, opensearch.Hangzhou)
	args := opensearch.ScrollArgs{
		Args: opensearch.Args{
			AppName: "test_article_search",
			Query:   `query=(title:"搜索" AND title:"搜索") OR (content:"搜索" AND content:"搜索")&&config=start:0,hit:10,format:fulljson`,
		},
		Scroll:     "5s",
		SearchType: "scan",
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

func TestOpClient_Search(t *testing.T) {
	opClient := opensearch.NewClient(networkType, accessKeyId, accessKeySecret, opensearch.Hangzhou)
	args := opensearch.SearchArgs{
		Args: opensearch.Args{
			AppName: appName,
			Query:   `config=start:0,hit:10,format:fulljson&&query=(title:"搜索" AND title:"搜索") OR (content:"搜索" AND content:"搜索")`,
		},
	}
	resp, err := opClient.Search(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
