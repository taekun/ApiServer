package main

import (
	"log"
	"flag"
	"github.com/taekun/apiserver/internal/app/apiserver"
)

var(
	configpath string
)

func init(){
	flag.StringVar(&configpath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()
{
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configpath, config)
	if err !=nil{
		log.Fatal(err)
	}

	s := apiserver.New()
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}