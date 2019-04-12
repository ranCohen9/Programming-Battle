package main
import(
	"fmt"
	"encoding/json"
    //"log"
    "net/http"
	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("server starting");
	router:=mux.NewRouter();
	router.HandleFunc("/api/test/{id}", doGetMethod).Methods("GET");
	router.HandleFunc("/api/test", doPostMethod).Methods("POST");
	fmt.Println("server listening on port 8000");	
	http.ListenAndServe(":8000",router);
}

func doGetMethod(w http.ResponseWriter, r *http.Request){
	//fmt.Println("go get operation");	
	params:=mux.Vars(r);
	id:=params["id"];
	json.NewEncoder(w).Encode(id);
}

func doPostMethod(w http.ResponseWriter, r *http.Request){
	//fmt.Println("go post operation");
	decoder:=json.NewDecoder(r.Body);
	var obj json.RawMessage;
	decoder.Decode(&obj);
	//fmt.Println(string(obj));
	json.NewEncoder(w).Encode(obj);
}