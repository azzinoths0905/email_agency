package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error config file: %s", err))
	}
}

func main() {
	router := newRouter("SendMail", "POST", "/api/v0/email", SendMail)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func newRouter(name string, method string, path string, handlerFunc http.HandlerFunc) (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	router.Name(name).Methods(method).Path(path).HandlerFunc(handlerFunc)
	return
}
