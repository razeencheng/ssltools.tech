package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kataras/iris"
)

var buildTime string
var version string

func main() {

	fmt.Println("BuildTime: ", buildTime)
	fmt.Println("Version: ", version)

	var v bool
	flag.BoolVar(&v, "v", false, "version")
	flag.Parse()

	if v {
		os.Exit(0)
	}

	app := iris.Default()

	tmpl := iris.HTML("./public", ".html")
	tmpl.Reload(os.Getenv("RUN_MODE") == "dev")

	app.StaticWeb("/public", "./public")
	app.RegisterView(tmpl)
	app.Get("/", HandleHome)
	app.Get("demo/{id}", HandleDemoView)

	app.Run(iris.Addr(":8080"))
}

// Site doc
type Site struct {
	Title     string
	Author    string
	Demos     *DemoTable
	CopyRight string
	Beian     string
}

// DemoTable doc
type DemoTable struct {
	THead  []string
	Fields []*DemoField
}

// DemoField doc
type DemoField struct {
	ID     int
	Fields []string
	URL    string
}

var siteInfo = Site{
	Title:  "SSL Demo 演示网站",
	Author: "By Razeen",
	Demos: &DemoTable{
		THead: []string{"名称", "描述"},
		Fields: []*DemoField{
			&DemoField{1, []string{"TLSv1.3 常规演示页面", "TLSv1.3 开启基本特性"}, "https://tls13.ssltools.tech"},
			&DemoField{2, []string{"TLSv1.3 常规演示页面", "TLSv1.3 开启基本特性"}, "https://tls13.ssltools.tech"},
		},
	},
	CopyRight: "2019 SSLTOOLS.TECH",
	Beian:     "沪ICP备19009734号-1",
}

// Demo doc
type Demo struct {
	Title string
	Desc  string
}

// HandleHome 首页
func HandleHome(c iris.Context) {
	c.ViewData("", siteInfo)
	c.View("index.html")
}

// HandleDemoView 处理不同Demo页展示
func HandleDemoView(c iris.Context) {

	demo := &Demo{
		Title: "Demo 不存在",
	}

	id, err := strconv.Atoi(c.Params().Get("id"))
	if err == nil {
		for _, v := range siteInfo.Demos.Fields {
			if v.ID == id {
				demo.Title = v.Fields[0]
				demo.Desc = v.Fields[1]
			}
		}
	}

	c.ViewData("", demo)
	c.View("demo.html")
	return
}
