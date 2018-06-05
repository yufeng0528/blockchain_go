package main

import (
	"net/http"

	"github.com/c4pt0r/ini"
)

func main() {
	var conf = ini.NewConf("config.cfg")
	nodeID := conf.String("section_1", "nodeId", "3001")
	conf.Parse()

	router := NewRouter()
	http.ListenAndServe("localhost:"+*nodeID, router)

	cli := CLI{}
	cli.Run()

}
