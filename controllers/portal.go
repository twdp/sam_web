package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam-core/dto/req"
	"tianwei.pro/sam-core/dto/res"
	"tianwei.pro/sam-web/rpc"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/api/v1/captcha", store)
}

type PortalController struct {
	controller.RestfulController
}

// @router /email [post]
func (p *PortalController) LoginByEmail() {
	email := p.GetString("email")
	password := p.GetString("password")

	if email == "" || password == "" {
		p.SetSession("needCpt", true)
		p.E500(business.H{
			"msg":     "账号或密码不能为空",
			"needCpt": true,
		})
		return
	}

	reply := &res.LoginDto{}
	if err := rpc.UserFacade.Call(context.Background(), "Login", &req.EmailLoginDto{
		Email:    email,
		Password: password,
	}, reply); err != nil {
		logs.Error("call login failed. %v", err)
		p.E500("系统错误,稍后重试")
	}
	if reply.Err != nil {
		count := p.GetSession("passwordWrongCount")
		c := 1
		if count == nil {
			p.SetSession("passwordWrongCount", c)
		} else {
			c = count.(int)
			p.SetSession("passwordWrongCount", c+1)
		}
		result := make(map[string]interface{})
		result["msg"] = reply.Err.Error()
		if c > 2 {
			result["needCpt"] = true
		}
		p.E500(result)
		return
	}

	param := sam_agent.VerifyTokenParam{
		SystemInfoParam: sam_agent.SystemInfoParam{
			AppKey: beego.AppConfig.String("appkey"),
			Secret: beego.AppConfig.String("secret"),
		},
		Token: reply.Token,
	}

	userInfo := &sam_agent.UserInfo{}
	if err := rpc.SamAgentFacade.Call(context.Background(), "VerifyToken", param, userInfo); err != nil {
		p.E500("系统错误,稍后重试")
		return
	}
	if reply.Err != nil {
		p.E500("系统错误,稍后重试")
		return
	}

	p.SetSession(sam_agent.SamUserInfoSessionKey, userInfo)

	// 跨 域Secure必须为false
	p.Ctx.SetCookie(sam_agent.SamTokenCookieName, reply.Token, 30*24*60*60, "/", beego.AppConfig.DefaultString("topdomain", "/"), false, true)

	p.ReturnJson(business.H{
		"token": reply.Token,
	})
}
