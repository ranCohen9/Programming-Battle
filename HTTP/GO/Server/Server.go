package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	router:=mux.NewRouter();

	http.ListenAndServe(":4000",router);
}

func test(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode()
}