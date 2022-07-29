package main

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//embedTest()
	//embedFS()
	embedTemplate()
	//embedMutex()

}

//go:embed templates cpp
//go:embed cpp
var content embed.FS

func embedMutex() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}

//在模板中的应用
//go:embed templates
var tmpl embed.FS

func embedTemplate() {
	t, _ := template.ParseFS(tmpl, "templates/*.tmpl")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(rw, "index.tmpl", map[string]string{"title": "Golang Embed 测试"})
	})
	http.ListenAndServe(":8080", nil)
}

//go:embed "version.txt"
var f embed.FS

//go:embed templates/*
//go:embed cpp/*
var ff embed.FS

//go:embed templates/*.html
var fHtml embed.FS

//go:embed templates/t[2-3].html
var fHtml1 embed.FS

func embedFS() {
	data, err := f.ReadFile("version.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	fmt.Println()
	fs, err := ff.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range fs {
		fmt.Println(entry.Name())
	}
	fs, err = ff.ReadDir("cpp")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range fs {
		fmt.Println(entry.Name())
	}

	fmt.Println()
	fs, err = fHtml.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range fs {
		fmt.Println(entry.Name())
	}

	fmt.Println()
	fs, err = fHtml1.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range fs {
		fmt.Println(entry.Name())
	}
}

//go:embed version.txt
var version string

//go:embed "version info.txt"
var version1 string

// 软链接&硬链接
// 在当前目录下创建一个指向version.txt的软链接 v
// ln -s version.txt v
////go:embed v
//var v string
//提示pattern v: cannot embed irregular file v
//结论：//go:embed指令不支持文件的软链接

// 在当前目录下创建一个指向version.txt的硬链接 h
// ln  version.txt h
//go:embed h
var h string

//结论：//go:embed指令支持文件的硬链接。因为硬链接本质上是源文件的一个拷贝。

//go:embed version.txt
var versionByte []byte

func embedTest() {

	fmt.Printf("version: %q\n", version)
	fmt.Printf("version1: %q\n", version1)
	//fmt.Printf("v: %q\n", v)
	fmt.Printf("h: %q\n", h)
	fmt.Printf("versionByte: %q\n", versionByte)
}
