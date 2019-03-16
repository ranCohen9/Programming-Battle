package main

import(
	"os"
	"fmt"
	"time"
	//"io/ioutil"
)

const FilesDirName="files";
func CreateDirectory(){
	err:=os.RemoveAll(fmt.Sprintf("./%s/",FilesDirName));
	if err!=nil{
		fmt.Println(err);
	} else{
	fmt.Println("no error");
	}
	os.Mkdir(FilesDirName,os.ModePerm);
}

func CreateFile(fileNumber int){
	os.Create(fmt.Sprintf("%s\\file-%d.txt",FilesDirName, fileNumber));
}

func CreateManyFiles(){
	CreateDirectory();
	start :=time.Now();
	for i:=0; i<1000;i++{
		//go CreateFile(i); fault 
		CreateFile(i);
	}
	stop:=time.Now();
	fmt.Println("create file elapsed: ", stop.Sub(start));
}

func WriteToFile(){
	os.Create("contentFile.txt");
	start :=time.Now();
	for i:=0;i<1000;i++{
		f,err:=os.OpenFile("contentFile.txt",os.O_WRONLY|os.O_APPEND,0644);
	if(err!=nil){
		fmt.Println(err);
	}
	f.WriteString(fmt.Sprintf("index is %d\n",i));
	f.Close();
	}
	stop:=time.Now();
	fmt.Println("Write file elapsed: ", stop.Sub(start));
}

func test(index int){
	
}