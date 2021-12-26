package main

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	dao2 "github.com/gogf/gf/v2/net/gmetrics/example/metrics/dao"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	ServerPort = 3001
)

func init() {
	// 启动prometheus
	g.Server().BindHandler("/metrics", func(r *ghttp.Request) {
		promhttp.Handler().ServeHTTP(r.Response.Writer, r.Request)
	})

	gredis.SetConfig(&gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      0,
	})
	gdb.AddDefaultConfigNode(gdb.ConfigNode{
		Host: "mysql",
		Port: "3306",
		User: "root",
		Pass: "xy123456",
		Name: "mysql",
		Type: "mysql",
	})
}

func main() {
	g.Server().SetPort(ServerPort)
	routes()
	g.Server().Run()
}

func routes() {
	s := g.Server()

	// http监控的中间件
	s.Use(ghttp.Metrics)

	s.BindHandler("/user", func(r *ghttp.Request) {
		// mysql-metrics
		entities, err := dao2.User.Ctx(r.Context()).All()
		if err != nil {
			g.Log().Error(r.Context(), err)
			r.Response.WriteStatusExit(500)
		}
		err = r.Response.WriteJsonExit(entities)
		if err != nil {
			g.Log().Error(r.Context(), err)
			r.Response.WriteStatusExit(500)
		}
	})
	s.BindHandler("/user-rpc", func(r *ghttp.Request) {
		res, err := gclient.New().Use(gclient.Metrics).SetCaller("rpc").Get(r.Context(), fmt.Sprintf("http://127.0.0.1:%d/user", ServerPort))
		if err != nil {
			r.Response.WriteStatusExit(500)
		}
		r.Response.WriteExit(res.ReadAllString())
	})

	s.BindHandler("/user-cache", func(r *ghttp.Request) {
		// redis-metrics
		cacheKey := "user-cache"
		results, err := g.Redis().Do(r.Context(), "GET", cacheKey)
		if err != nil {
			r.Response.WriteStatusExit(500)
		}
		if results.Interface() == nil {
			res, err := gclient.New().Use(gclient.Metrics).SetCaller("cache").Get(r.Context(), fmt.Sprintf("http://127.0.0.1:%d/user", ServerPort))
			if err != nil {
				r.Response.WriteStatusExit(500)
			}
			result := res.ReadAllString()

			_, err = g.Redis().Do(r.Context(), "SET", cacheKey, result)
			if err != nil {
				r.Response.WriteStatusExit(500)
			}
			err = r.Response.WriteJsonExit(result)
			if err != nil {
				r.Response.WriteStatusExit(500)
			}
		}

		err = r.Response.WriteJsonExit(results)
		if err != nil {
			r.Response.WriteStatusExit(500)
		}
	})

	s.BindHandler("/time-zone", func(r *ghttp.Request) {
		r.Response.WriteStatusExit(500)
	})
}
