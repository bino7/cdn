package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/bino7/config"
	"cdn/lib"
)

var conf, _ = config.ParseYaml(`
debug: true
port: 5000
maxSize: 1000
showInfo: true
tailOnly: false
root: /home/bino/cdn/
`)

func main() {
	conf.Env().Flag()
	r := martini.NewRouter()
	m := martini.New()
	if conf.UBool("debug") {
		m.Use(martini.Logger())
	}
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	m.Use(func(){

	})
	logger := log.New(os.Stdout, "\x1B[36m[cdn] >>\x1B[39m ", 0)
	m.Map(logger)

	r.Group("", cdn.Cdn(cdn.Config{
		Root:	  conf.UString("root"),
		MaxSize:  conf.UInt("maxSize"),
		ShowInfo: conf.UBool("showInfo"),
		TailOnly: conf.UBool("tailOnly"),

	}))

	logger.Println("Server started at :" + conf.UString("port", "5000"))
	_err := http.ListenAndServe(":"+conf.UString("port", "5000"), m)
	if _err != nil {
		logger.Printf("\x1B[31mServer exit with error: %s\x1B[39m\n", _err)
		os.Exit(1)
	}
}
