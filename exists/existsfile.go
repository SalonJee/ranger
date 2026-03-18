package exists

import (
	"errors"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"ranger.com/dat"
)


func checkExtensionExists(DirData string,Extension string)bool {
	isit:=strings.Contains(DirData,Extension)
	return isit
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}


func checkExistence(DirData string,Extension string)string{
    if checkExtensionExists(DirData,Extension)==true{
	  	if checkFileExists("images")!=true || checkFileExists("videos")!=true{
 
		return "state1"  //state1 means, extension exists, but the folder doesnt
   
}
	} else if checkExtensionExists(DirData,Extension)==true{
	  	if checkFileExists("images")==true || checkFileExists("videos")==true{
 
		return "state2" //both the extension and folder exist
   
}

}
return "no -action"
}


func fullfileNameExtract(DirData string,Extension string) string{
	// to extract the complete file name
//locating the extensions index
extIdx:=strings.Index(DirData,Extension)

//now looking at the text before the position of ext
prefix :=DirData[:extIdx]
if prefix == "[" || prefix== "]"{
	prefix=""
}

//finding space index(position), it searches until it finds the first space
firstboundary :=strings.LastIndexAny(prefix," []()\"'")
// now extracting the files name from the space, to the end of the position of ext
startIdx :=firstboundary +1
endIdx := extIdx + len(Extension)

//now finally extracting the text
return DirData[startIdx:endIdx]
}



//here this function only requires the correct DirecName, and extension required to be checked
func Existence(DirecName string,Extension string){
DirecString :=dat.DirectoryDat(DirecName)

if(checkExistence(DirecString,Extension)=="state1"){
	if Extension == ".png"|| Extension== "png"|| Extension== ".jpeg" ||Extension== "jpeg" {
		fileName :=fullfileNameExtract(DirecString,Extension)
		fmt.Printf("%v\n",fileName)
	   targetpath :=filepath.Join(DirecName,"images")
	   os.Mkdir(targetpath,0755)

	   
	    oldpath:=filepath.Join(DirecName,fileName)
	 fmt.Printf("%v\n",oldpath)
	newpath :=filepath.Join(targetpath,fileName)
	fmt.Printf("%v\n",newpath)
	os.Rename(oldpath,newpath)
	
	}
	if Extension == ".mp4"|| Extension== ".mov" {
		fileName :=fullfileNameExtract(DirecString,Extension)
	   targetpath :=filepath.Join(DirecName,"videos")
	   os.Mkdir(targetpath,0755)

	   oldpath:=filepath.Join(DirecName,fileName)
	  fmt.Printf("%v\n",oldpath)
	newpath :=filepath.Join(targetpath,fileName)
	 fmt.Printf("%v\n",newpath)
	os.Rename(oldpath,newpath)
	
}else if(checkExistence(DirecString,Extension)=="state2"){
}	
}

}