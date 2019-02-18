package rpc

import (
	"github.com/smallnest/rpcx/client"
)

var (
	UserFacade     client.XClient
	SystemFacade   client.XClient
	SamAgentFacade client.XClient
)

func getClients(d client.ServiceDiscovery) {
	getUserFacade(d)
	getSystemFacade(d)
	getSamAgentFacade(d)
}

func getUserFacade(d client.ServiceDiscovery) {
	UserFacade = client.NewXClient("UserFacadeImpl", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	//fmt.Println(UserFacade.Call(context.Background(), "Login", &req.EmailLoginDto{}, &res.LoginDto{}))
}

func getSystemFacade(d client.ServiceDiscovery) {
	SystemFacade = client.NewXClient("SystemFacadeImpl", client.Failtry, client.RandomSelect, d, client.DefaultOption)
}

func getSamAgentFacade(d client.ServiceDiscovery) {
	SamAgentFacade = client.NewXClient("SamAgentFacadeImpl", client.Failtry, client.RandomSelect, d, client.DefaultOption)
}

func CloseClient() {
	if UserFacade != nil {
		UserFacade.Close()
	}
	if SystemFacade != nil {
		SystemFacade.Close()
	}

	if SamAgentFacade != nil {
		SamAgentFacade.Close()
	}
}
