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
	fmt.Printf("prop=%+v", mysql)
}
