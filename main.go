package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main() {
	proxy := &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
			//"GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
		Cacher: goproxy.DirCacher("/data/goproxy"),
	}
	if err := http.ListenAndServe("0.0.0.0:8765", proxy); err != nil {
		panic(err)

	}
}
