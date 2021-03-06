package main
import (
"net/http"
"log"
"fmt"
"time"
"bytes"
)
const avgIterations=10;
const manyIterations=100;

func main(){
	avgGetMany();
	avgPostMany();
}

func avgGetMany(){
	sum:=float64(0);
	for i:=0;i<avgIterations;i++ {
		sum+=getMany();
	}
	log.Printf("average of %d*%d GET requests - %f",avgIterations,manyIterations, sum/avgIterations);
}

func avgPostMany(){
	sum:=float64(0);
	for i:=0;i<avgIterations;i++{
		sum+=postMany();
	}
	log.Printf("average of %d*%d POST requests - %f",avgIterations,manyIterations, sum/avgIterations);
}

func getMany()(float64){
	sum,_:= time.ParseDuration("0s");
	i:=0;
	for ;i<manyIterations;i++ {
		sum+=getOnce(i);
	}
	//log.Println(sum);
	// log.Printf("finished get loop avg %fms", sum.Seconds()*1000/float64(i));
	return sum.Seconds()*1000/float64(i);
}

func postMany()(float64){
	sum,_:=time.ParseDuration("0s");
	i:=0;
	for ;i<manyIterations;i++{
		sum+=postOnce(i);
	}
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
	resp,err:=http.Post("http://localhost:8000/api/test","application/json",bytes.NewBufferString(string(`{"id": 11}`)));
	if(err!=nil){
		log.Println(err);
	}
	resp.Body.Close();
	//log.Println(resp.Status);
	//log.Println(resp.B);
	stop:=time.Now();
	return stop.Sub(start);
}