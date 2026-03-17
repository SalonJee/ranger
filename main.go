package main

import (
	//"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	// "path/filepath"
)


type DataBox struct{
	Image string `json:img`
	Video string `json:mov`
	Text string   `json:txt`

}

func checkExtensionExists(DirData string,Extension string)bool {
	isit:=strings.Contains(DirData,Extension)
	return isit
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}


func main(){

	Dir,err :=os.ReadDir(".")
	if(err !=nil){
	fmt.Println("error occurered : ",err)
	}

//for _,file:=range Dir {
	//fmt.Printf("the file is %s ",file.Name())
//}
//files[]:=file.Name()
///if(file.Name()=="image.png"){

//}

var files[] string
for _,file:=range Dir{
	files =append(files,file.Name())
//fmt.Print(files)

}

fmt.Printf("%+v\n",files)

filesStringed :=fmt.Sprint(files)


fmt.Print(checkExtensionExists(filesStringed,".png"))


if checkFileExists("images")!=true{

   if checkExtensionExists(filesStringed,".png")==true{
	print("yes, the folder doenst exists, and the file png also exists")

}

}

}


