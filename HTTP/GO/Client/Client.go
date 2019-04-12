package main
import (
"net/http"
"log"
"fmt"
"time"
)


func main(){
	getMany();
	//avgGetMany();
}

func avgGetMany(){
	sum:=float64(0);
	for i:=0;i<100;i++ {
		sum+=getMany();
	}
	log.Printf("total avg %f", sum/100);
}

func getMany()(float64){
	sum,_:= time.ParseDuration("0s");
	i:=0;
	for ;i<100;i++ {
		sum+=getOnce(i);
	}
	//log.Println(sum);
	log.Printf("finished get loop avg %fms", sum.Seconds()*1000/float64(i));
	return sum.Seconds()*1000/float64(i);
}

func getOnce(id int)(time.Duration){
	start:=time.Now();
	_, err:=http.Get(fmt.Sprintf("http://localhost:8000/api/test/%d",id));
	if(err!=nil){
		log.Println(err);
	}
	//log.Println(resp);
	stop:=time.Now();
	//log.Println(stop.Sub(start));
	return stop.Sub(start);
}

func postOnce(id int)(time.Duration){
	start:=time.Now();
	_,err:=http.Post("http://localhost:8000/api/test","text/json",nil);
	if(err!=nil){
		log.Println(err);
	}
	stop:=time.Now();
	return stop.Sub(start);
}