package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/informeai/gogql/controllers"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (ro *Router) GraphQL() error {
	var e error
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				e = err
				w.Write([]byte(fmt.Sprintf("{status: error, msg: %v}", err.Error())))
				return
			}
			defer r.Body.Close()
			result, err := controllers.NewGraphQlController().Exec(string(body))
			if err != nil {
				e = err
				w.Write([]byte(fmt.Sprintf("{status: error, msg: %v}", err.Error())))
				return
			}
			w.Write([]byte(fmt.Sprintf("{status: success, %v}", result)))
			return
		}
		w.Write([]byte("not method allowed"))
	})
	return e
}

func (ro *Router) Start() error {
	if err := ro.GraphQL(); err != nil {
		return err
	}
	fmt.Printf("running in port: %v\n", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil); err != nil {
		return err
	}
	return nil
}
