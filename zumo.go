package zumo

import (
	"github.com/ghodss/yaml"

	"github.com/zumak/zumo/backend"
	"github.com/zumak/zumo/http-server"
	"github.com/zumak/zumo/utils/log"
)

const defaultConfig = `
backend:
  store:
    driver: etcd
    endpoint: localhost:2379
server:
  bind: localhost:5000
`

var VERSION string

type Config struct {
	Backend backend.Config
	Server  server.Config
}

func Run() {
	conf := &Config{}
	err := yaml.Unmarshal([]byte(defaultConfig), conf)
	if err != nil {
		log.Err(err)
		return
	}

	log.Info("load complete, %+v", conf)

	_, err = backend.New(&conf.Backend)
	if err != nil {
		log.Err(err)
		return
	}

	log.Info("Init backend")

	// load core
	if err := server.Run(&conf.Server); err != nil {
		if err != nil {
			log.Err(err)
		}
	}
}
