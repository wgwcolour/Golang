package gee

import (
	"fmt"
	"net/http"
)

// 通用的接口方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 引擎
type Engine struct {
	router map[string]HandlerFunc
}

// 创建实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	engine.router[method+"-"+pattern] = handler
}

// *实现接口*
func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler,ok := engine.router[key]; ok{
		handler(w,req)
	}else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"404 Not Found:%s\n",req.URL)
	}
}

// GET方法
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	method := "GET"
	engine.addRouter(method, pattern, handler)
}

// 启动方法
func (engine *Engine) Run(addr string)(err error) {
	return http.ListenAndServe(addr,engine)
}
