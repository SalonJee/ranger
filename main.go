package main

import (
	
	"ranger.com/exists"
)

func main(){
	exists.Existence("./trialDir",".png")
	exists.Existence("./trialDir",".md")
	exists.Existence("./trialDir",".mp4")
}