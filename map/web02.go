package main

import (
	"net/http"
	"fmt"
)

// 如下代码所示，我们自己实现了一个简易的路由器

type MyMux struct {

}

func HelloName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello myroute !")
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/"{
		HelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
