package middleware

import (
	"net"
	"net/http"

	"github.com/webpkg/web"
)

// Chain call something before next callback
// func Chain(next web.Callback, keys ...string) web.Callback {

// 	return func(ctx *web.Context) (web.Data, error) {

// 		before(ctx)

// 		err := bearerAuth(ctx, keys...)

// 		if err != nil {
// 			return nil, err
// 		}

// 		val, err := next(ctx)

// 		after(ctx)

// 		return val, err
// 	}
// }

// Direct call something before next callback
func Direct(next web.Callback, keys ...string) web.Callback {

	return func(ctx *web.Context) (web.Data, error) {

		before(ctx)

		val, err := next(ctx)

		after(ctx)

		return val, err
	}
}

// before call before controller action
func before(ctx *web.Context) {
	ctx.SetContentType("application/json; charset=utf-8")
	//ctx.SetHeader("access-control-allow-origin", "*")
}

// after call after controller action
func after(ctx *web.Context) {
	// fmt.Println("开始记录日志了")

	//测试服务器的IP集合
	// serverIp := proxy.WebConfig().App.ServerIP
	// //获取到的IP
	// ip := RemoteIp(ctx.Request)
	// //循环判断IP
	// ips := strings.Split(serverIp, ",")
	// flag := true
	// for i := 0; i < len(ips); i++ {
	// 	if ips[i] == ip {
	// 		flag = false
	// 		break
	// 	}
	// }

	// //如果不是内部跳转才记录
	// if flag && ip != "127.0.0.1" {
	// 	//记录日志
	// 	// token := ctx.Request.Header.Get("Authorization")
	// 	// fmt.Printf("token的值为：%v \n", token)
	// 	log := model.CreateAccessLog()
	// 	// ip := RemoteIp(ctx.Request)
	// 	// fmt.Printf("ip的值为：%v \n", ip)
	// 	log.ActionIP = ip

	// 	log.RequestMethod = ctx.Request.Method
	// 	log.UserAgent = ctx.Request.Header.Get("User-Agent")

	// 	log.ActionAPI = ctx.Request.Host + ctx.Request.RequestURI
	// 	log.UserID = ctx.UserID

	// 	proxy.PostAccessLogDirect(log)

	// 	//如果token存在且不是匿名用户
	// 	// if token != "" && ctx.UserID != 0 {
	// 	// 	// fmt.Println("token不为空")
	// 	// 	log.UserID = ctx.UserID
	// 	// 	proxy.PostAccessLog(log, token)
	// 	// } else {
	// 	// 	//匿名用户和没有token的接口
	// 	// 	// fmt.Println("token为空")
	// 	// 	log.UserID = 0
	// 	// 	proxy.PostAccessLogDirect(log)

	// 	// }
	// }

}

// bearerAuth bearer authorization
// func bearerAuth(ctx *web.Context, keys ...string) error {

// 	auth := ctx.GetHeader("Authorization")

// 	accessToken, err := rbac.TryParseBearerToken(auth)

// 	if err != nil {
// 		return err
// 	}

// 	cat, err := proxy.GetAuthByAccessToken(accessToken)

// 	if err != nil {
// 		return err
// 	}

// 	ctx.UserID = cat.UserID

// 	if !rbac.Check(cat.Right, keys...) {
// 		return web.ErrForbidden
// 	}

// 	return nil
// }

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-Ip"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	// if remoteAddr == "::1" {
	// 	remoteAddr = "127.0.0.1"
	// }

	return remoteAddr
}
