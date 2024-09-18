package test

import (
	"config-sync/pkg/http/client"
	"fmt"
	"testing"
)

func TestHttpClient(t *testing.T) {
	client := client.NewHttpClient("http://172.27.10.16:8848/nacos/v1/ns/instance/list?NamespaceId=59025239-e2a7-4f5d-856f-e04356f7f043&healthyOnly=true&serviceName=ihawk-gateway-control", client.GET, "")
	code, body, err := client.Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)
	fmt.Println(string(body))
}
