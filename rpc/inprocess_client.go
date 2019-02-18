// +build !zookeeper

package rpc

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	_ "tianwei.pro/sam-core/init"
	"tianwei.pro/sam-core/rpc"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:anywhere@tcp(127.0.0.1:3306)/sam?charset=utf8&loc=Asia%2FShanghai", 30)

	// create table
	orm.RunSyncdb("default", false, true)

	if beego.BConfig.RunMode != "prod" {
		orm.Debug = true
	}

	d := client.NewInprocessDiscovery()

	s := server.NewServer()
	rpc.AddRegistryPlugin(s)
	rpc.RegisterToRpcx(s)

	go func() {
		s.Serve("tcp", "localhost:1999")
	}()

	getClients(d)
}
