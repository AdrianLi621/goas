/**
 * @Author adrianli
 * @Description //TODO $
 * @Date 2021/12/6 17:06
 **/
package main

import (
	"flag"
	"github.com/spf13/viper"
	"goas/unit"
	"log"
)

var (
	configFile = flag.String("db", "config/", "the config file")
)

/**
 *  init
 *  @Description:
 **/
func init() {
	flag.Parse()
	viper.SetConfigName("db")
	viper.AddConfigPath(*configFile)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

/**
 *  main
 *  @Description:
 **/
func main() {
	//生成model
	//生成migrate
	//生成curd

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	unit.InitModel(viper.GetViper())

}
