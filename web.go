package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
		0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x28,
		0x01, 0x03, 0x00, 0x00, 0x00, 0xb6, 0x30, 0x2a, 0x2e, 0x00, 0x00, 0x00,
		0x03, 0x50, 0x4c, 0x54, 0x45, 0x5a, 0xc3, 0x5a, 0xad, 0x38, 0xaa, 0xdb,
		0x00, 0x00, 0x00, 0x0b, 0x49, 0x44, 0x41, 0x54, 0x78, 0x01, 0x63, 0x18,
		0x61, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x01, 0xe2, 0xb8, 0x75, 0x22, 0x00,
		0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		// fmt.Println("username:", r.Form["username"])
		submint := r.Form["username"]
		ss := fmt.Sprintf(strings.Join(submint, ","))
		if ss == "picture" {
			w.Write([]byte{
				0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
				0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x28,
				0x01, 0x03, 0x00, 0x00, 0x00, 0xb6, 0x30, 0x2a, 0x2e, 0x00, 0x00, 0x00,
				0x03, 0x50, 0x4c, 0x54, 0x45, 0x5a, 0xc3, 0x5a, 0xad, 0x38, 0xaa, 0xdb,
				0x00, 0x00, 0x00, 0x0b, 0x49, 0x44, 0x41, 0x54, 0x78, 0x01, 0x63, 0x18,
				0x61, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x01, 0xe2, 0xb8, 0x75, 0x22, 0x00,
				0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
			})
		}
		if ss == "方程求解" {
			fmt.Fprintln(w, "2110310056于见昊")
			fmt.Fprintln(w, "方程求解")
			fmt.Fprintln(w, "————————————————————————————")
			fmt.Fprintln(w, "ID-----a------b------c")
			fmt.Fprintln(w, "1-----1.0-----4.0----3.0")
			fmt.Fprintln(w, "\n方程具有两个实根解\n")
			fmt.Fprintln(w, "r1=-1.0,r2=-3.0")
			fmt.Fprintln(w, "————————————————————————————")
			fmt.Fprintln(w, "ID-----a------b------c")
			fmt.Fprintln(w, "2-----0.0----2.0----2.0")
			fmt.Fprintln(w, "\n方程不是正规解\n")
			fmt.Fprintln(w, "————————————————————————————")
			fmt.Fprintln(w, "ID-----a------b------c")
			fmt.Fprintln(w, "3-----4.0-----2.0----1.0")
			fmt.Fprintln(w, "\n方程无实根\n")
			fmt.Fprintln(w, "————————————————————————————")
			fmt.Fprintln(w, "ID-----a------b------c")
			fmt.Fprintln(w, "4-----2.0-----4.0----2.0")
			fmt.Fprintln(w, "\n方程具有相同实根解\n")
			fmt.Fprintln(w, "r1=r2=-1.0")
			fmt.Fprintln(w, "————————————————————————————")
			fmt.Fprintln(w, "ID-----a------b------c")
			fmt.Fprintln(w, "5-----0.0-----0.0----0.0")
			fmt.Fprintln(w, "\n方程不是正规解\n")
		}
		if ss == "智能体" {
			fmt.Fprintln(w, "2110310056于见昊")
			fmt.Fprintln(w, "这是智能体设计、实现、写作的实验测试")
			fmt.Fprintln(w, "可信度量:,5000,3,0.1")
		}
		if ss == "小八哥吃大花生" {
			fmt.Fprintln(w, "2110310056于见昊")
			fmt.Fprintln(w, "小/八哥/吃/大/花生")
			fmt.Fprintln(w, "词法元Token识别:小\n词法元Token识别:八哥\n词法元Token识别:吃\n词法元Token识别:大\n词法元Token识别:：花生")
		}

	}
}

func main() {

	http.HandleFunc("/picture", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
