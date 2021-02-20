package main

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/ei/web"
	"github.com/imakiri/playground/gate"
	"github.com/spf13/viper"
	"log"
)

func launch(c cfg.Config) error {
	var err error
	var gs gate.GeneralService
	var ws *web.Service

	gs, err = gate.NewService(c)
	if err != nil {
		return err
	}

	ws, err = web.NewService(c.EI, gs)
	if err != nil {
		return err
	}

	rsc := make(chan error)
	sc := make(chan error)

	if c.EI.Web.LaunchRedir {
		go func(rsc chan error) {
			rsc <- ws.RedirServer.ListenAndServe()
		}(rsc)
	} else {
		println("Redirect server launch has been cancelled")
	}

	go func(sc chan error) {
		sc <- ws.Server.ListenAndServeTLS("cert.pem", "privkey.pem")
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func main() {
	var conf cfg.Config
	var err error
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(launch(conf))
}
