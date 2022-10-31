package main

import (
	"fmt"
	web2 "github.com/beego/beego/v2/server/web"
	"github.com/ecnepsnai/web"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.StartServer()
	server := web.New("127.0.0.1:8080")
	fmt.Println(server)

	web2.Run()
}
