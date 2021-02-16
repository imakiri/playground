package main

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/pkg/ui"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func launch(s core.Settings) error {
	var server = &http.Server{}
	var redirServer = &http.Server{}

	router, redirRouter, _ := ui.NewWebRouters(s)
	server.Handler = router
	redirServer.Handler = redirRouter

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- redirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- server.ListenAndServeTLS("cert.pem", "privkey.pem")
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func main() {
	var conf core.Config
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

	s := core.Settings{
		Config: conf,
	}

	log.Fatal(launch(s))
}
