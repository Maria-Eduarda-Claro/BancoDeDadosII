package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InicializeApi() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHelloWorld)
	r.HandleFunc("/ aluno", handleAluno)
	http.ListenAndServe(":8080", r)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprint(w, "Olá Mundo!")
}

func handleAluno(w http.ResponseWriter, r *http.Request) {
	switch (r.Method=="POST"){
		fmt.Printf(w,"Em obras!!")

	}else{
		http.Error(w, "Método não implemetado", http.StatusMethodNotAllowed)
	}

}
