package test

import (
	"bytes"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	//// 定义一个upstream模板
	//upstreamTemplate := `
	//{{- define "upstream" -}}
	//upstream myapp1 {
	//	{{- range .Servers }}
	//	server {{.Address}}:{{.Port}};
	//	{{- end }}
	//}
	//{{- end }}
	//{{ template "upstream" . }}
	//`
	//
	//// 创建一个模板实例
	//tmpl, err := template.New("upstream").Parse(upstreamTemplate)

	// 定义一个upstream模板
	upstreamTemplate := `
	upstream myapp1 {
		{{- range .Servers }}
		server {{.Address}}:{{.Port}};
		{{- end }}
	}
	`
	tmpl, err := template.New("Servers").Parse(upstreamTemplate)
	if err != nil {
		panic(err)
	}

	// 准备要填充到模板中的数据
	servers := []struct {
		Address string
		Port    int
	}{
		{"192.168.1.1", 80},
		{"192.168.1.2", 80},
		{"192.168.1.3", 80},
	}

	// 执行模板填充数据
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]interface{}{
		"Servers": servers,
	})
	if err != nil {
		panic(err)
	}

	// 输出结果
	println(buffer.String())
}
