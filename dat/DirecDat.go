package dat

import (
	"fmt"
	"os"
)

func DirectoryDat(DirName string)string {
	Dir,err :=os.ReadDir(DirName)
	if(err !=nil){
	fmt.Println("error occurered : ",err)
	}
/*
for _,file:=range Dir {
	fmt.Printf("the file is %s ",file.Name())
}
files[]:=file.Name()
if(file.Name()=="image.png"){

//}
*/

var files[] string
for _,file:=range Dir{
	files =append(files,file.Name())
//fmt.Print(files)

}

fmt.Printf("%+v\n",files)

filesStringed :=fmt.Sprint(files)

return filesStringed
} // here it returns the total folder(directory) data in string form , 
   //can be used in other files