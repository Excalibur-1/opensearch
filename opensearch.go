package opensearch

import (
	"encoding/json"
	"strings"

	opensearch "github.com/Excalibur-1/opensearch/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	ConnectTimeout = 5000
	ReadTimeout    = 10000
	AutoRetry      = true
	IgnoreSSL      = false
	MaxIdleConn    = 50
)

// appName 可以是 app 的版本信息. 也可以是 app 名称.
// requestParams 信息 , 查询相关请求构造可参数可参考  https://help.aliyun.com/document_detail/204327.html
func search(endpoint, accessKeyId, accessKeySecret, appName string, requestParams map[string]interface{}, res interface{}) (err error) {
	// 创建请求用客户端实例
	// Endpoint 为 要访问服务的区域域名.
	// AccessKeyId 及AccessKeySecret 用于构造鉴权信息.
	config := &opensearch.Config{
		Endpoint:        tea.String(endpoint),
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	// New  一个client, 用以发送请求.
	client, err := opensearch.NewClient(config)
	// 如果 NewClient 过程中出现异常. 则 返回 _clientErr 且输出 错误信息.
	if err != nil {
		return
	}

	// 请求发送的配置参数. 用以请求配置及连接池配置.
	runtime := &util.RuntimeOptions{
		ConnectTimeout: tea.Int(ConnectTimeout),
		ReadTimeout:    tea.Int(ReadTimeout),
		Autoretry:      tea.Bool(AutoRetry),
		IgnoreSSL:      tea.Bool(IgnoreSSL),
		MaxIdleConns:   tea.Int(MaxIdleConn),
	}

	// 发送请求的方法调用.
	response, err := client.Request(
		tea.String("GET"),
		tea.String("/v3/openapi/apps/"+appName+"/search"),
		requestParams,
		nil,
		nil,
		runtime)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(response), &res); err != nil {
		return
	}

	return
}

type OpClient struct {
	endPoint, accessKeyId, accessKeySecret string
}

func NewClient(networkType, accessKeyId, accessKeySecret string, region Region) *OpClient {
	var sb strings.Builder
	sb.WriteString(networkType)
	sb.WriteString("opensearch-")
	sb.WriteString(string(region))
	sb.WriteString(".aliyuncs.com")
	return &OpClient{
		endPoint:        sb.String(),
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
	}
}

// -----------应用操作类API----------

// Search 提供查询数据的功能，最大返回5000
func (c *OpClient) Search(args SearchArgs) (resp SearchResponse, err error) {
	params := convertSearchArgsToMap(args)
	err = search(c.endPoint, c.accessKeyId, c.accessKeySecret, args.AppName, params, &resp)
	return
}

// Scroll 提供导出全部数据的功能
func (c *OpClient) Scroll(args ScrollArgs) (resp SearchResponse, err error) {
	params := convertScrollArgsToMap(args)
	err = search(c.endPoint, c.accessKeyId, c.accessKeySecret, args.AppName, params, &resp)
	return
}
