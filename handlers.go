package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	/*receive := html.EscapeString(r.URL.Path)
	  fmt.Printf(receive[1:])
	  fmt.Fprintf(w, "Hello, %q", receive[1:])*/
	fmt.Fprintln(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Todo Index")
	todos := Todos{Todo{Name: "Write presentation", Due: time.Now()},
		Todo{Name: "Host meetup", Due: time.Now()},
	}
	w.Header().Set("Context-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//如果返回json的格式报错，抛出异常
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Println(todoId)
	fmt.Fprintln(w, "Todo show:", todoId)
}
