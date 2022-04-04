package main

import (
	"boardgame-helper/middleware/doudizhu"
	"boardgame-helper/router"
	"boardgame-helper/utils/json"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "embed"

	"github.com/rs/cors"
)

//go:embed config.json
var configJson string

// from https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	host, port := json.InitConfig(configJson, GetOutboundIP().String(), "8888")
	r := router.Router()
	fmt.Printf("starting server on %v:%v\n", host, port)
	testItem := doudizhu.Aaa()
	fmt.Printf("%+v\n", testItem)
	log.Fatal(http.ListenAndServe(host+":"+port, cors.Default().Handler(r)))
}
