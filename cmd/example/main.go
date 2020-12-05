package main

import (
	"fmt"
	"flag"
	"strings"
	"io"
	"os"
)

var inputExpression = flag.String("e","","entering expression")
var inputFile = flag.String("f","","writeble file")
var outputFile = flag.String("o","","outputFilerd")

//ComputeHandler : adding new injection
/*type ComputeHandler struct{

}
*/

func main(){
	flag.Parse()
	var err  error
	var output io.Writer
	if *outputFile != "" {
		outputFile, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("file open error",err)
			os.Exit(1)
		}
	}	else{
		output = os.Stdout
	}
	var input io.Reader
	if *inputFile !=""{
		input,err := os.Open(*inputFile)
		if err!=nil {
			fmt.Println(err)
			return
		}
	}else{
		input = strings.NewReader(*inputExpression)
	}

	handler := &calc.ComputeHandler{
		Input: input,
		Output: output,
	}

}