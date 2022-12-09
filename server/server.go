package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"wf_api/server/internal/config"
	"wf_api/server/internal/handler"
	"wf_api/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var frontendPath = filepath.Join(os.Getenv("WF_DIR"), "dist")
var indexUrl = filepath.Join(frontendPath, "index.html")

var configFile = flag.String("f", "etc/server.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors(), rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		path = filepath.Join(frontendPath, path)
		_, err := os.Stat(path)
		if err != nil {
			http.ServeFile(w, r, indexUrl)
			return
		}
		http.ServeFile(w, r, path)
	})))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
