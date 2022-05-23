package gateway

import "github.com/TCP404/OneTiny-core/gateway/client"

type protocol string

const (
	HTTP    protocol = "http"
	RPC     protocol = "rpc"
	Process protocol = "process"
)

// @summary 获取客户端对象
// @param protocol 协议（http/rpc/process)
// @return *client.Client 客户端对象
// @return error 错误对象
func Client(protocol protocol) client.Client {
	switch protocol {
	case HTTP:
		return client.ProcessClient
	case RPC:
		return client.ProcessClient
	case Process:
		return client.ProcessClient
	default:
		return client.ProcessClient
	}
}
