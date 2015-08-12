package main

import(
	"fmt"
	"log"
	"net/http"
	"../"
	"time"
	"strconv"
)


type HttpHandler struct{
	num int
	es eventsource.EventSource
}

func (hh HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端连接")
	hh.es = eventsource.New(nil, nil)
	hh.es.SetOnClose(func(){
		fmt.Println("hehe~")
	})
	go func(){
		i := 0
		for{
			time.Sleep(10*time.Second)
			hh.es.SendEventMessage("11","init",strconv.Itoa(i))
			i++
			//fmt.Println("CC and Lcc ",hh.es.ConsumersCount() , hh.es.LiveConsumersCount())
		}
	}()
	hh.es.ServeHTTP(w,r)
}

func main(){
	
	
	httpHandler := HttpHandler{
		num : 0,
	}
	
	for prefix, handler := range map[string]http.Handler{
		"/test/": httpHandler,
	} {
		http.Handle(prefix, http.StripPrefix(prefix, handler))
	}

	log.Fatal(http.ListenAndServe(":8081", nil))
}