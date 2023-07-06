package config

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	config := GetConfig()
	config = SetConfig("config", "yaml", "", []string{"."})
	//config := NewConfig("config", "yaml", "", []string{"."})
	var mysql = &MySQL{}
	err := config.Load("mysql", mysql)
	//config.Config.ReadInConfig()
	//err := config.Config.UnmarshalKey("mysql", mysql)
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}
	fmt.Printf("mysql prop=%+v \n", mysql)

	var redis = &Redis{}
	err = config.Load("redis", redis)
	fmt.Printf("redis prop=%+v \n", redis)

	var hertz = &Hertz{}
	err = config.Load("hertz", hertz)
	fmt.Printf("hertz prop=%+v \n", hertz)

	var mongo = &Mongo{}
	err = config.Load("mongo", mongo)
	fmt.Printf("mongo prop=%+v \n", mongo)
}
