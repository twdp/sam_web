// +build zookeeper

package rpc

import (
	"github.com/astaxie/beego"
	"github.com/smallnest/rpcx/client"
)

func init() {
	zkAddrs := beego.AppConfig.DefaultStrings("zk.addr", []string{"localhost:2181"})
	d := client.NewZookeeperDiscovery("sam", "UserFacadeImpl", zkAddrs, nil)
	getUserFacade(d)

	d = client.NewZookeeperDiscovery("sam", "SystemFacadeImpl", zkAddrs, nil)
	getSystemFacade(d)

	d = client.NewZookeeperDiscovery("sam", "SamAgentFacadeImpl", zkAddrs, nil)
	getSamAgentFacade(d)
}
