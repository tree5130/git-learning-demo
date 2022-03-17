package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func f(writer http.ResponseWriter, request *http.Request) {
	//定义模板 function.tmpl
	//解析模板
	t := template.New("function.tmpl") // 创建一个模板对象, 名字与模板的名字对应
	_, err := t.ParseFiles("./function.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err%v", err)
		return
	}
	name := "渲染"
	//渲染模板
	t.Execute(writer, name)

}

func main() {
	http.HandleFunc("/", f)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("http server start failed err %v", err)
		return
	}
}

