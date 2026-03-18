package main

import (
	"os"
	"ranger.com/exists"
)

func main(){

	extension :=[...]string{"png","/.png",".jpeg",".jpg","jpg","jpeg",".mp4","mp4"}
	args :=os.Args
	if len(args) > 1 { //since args[0] is the main program name, we start at args[1]
    targetDir := args[1]
	for _, ext:=range extension{
  exists.Existence(targetDir,ext)
	}
	
	
}

}