package cdn

import (
	"github.com/go-martini/martini"
	"strings"
)

type Config struct {
	Root 	 string
	MaxSize  int
	TailOnly bool
	ShowInfo bool
}

var conf Config

func Cdn(c Config) func(r martini.Router) {
	conf = c
	if conf.MaxSize == 0 {
		conf.MaxSize = 1000
	}
	if strings.LastIndex(conf.Root,"/")!=len(conf.Root)-1 {
		conf.Root=conf.Root+"/"
	}

	return func(r martini.Router) {
		if conf.ShowInfo {
			//r.Get("/:name", getIndex)
			//r.Get("/:name/_stats", getStat)
		}
		r.Post("/:name",auth, post)
		r.Get("/:name/:file",auth, get)
	}
}


