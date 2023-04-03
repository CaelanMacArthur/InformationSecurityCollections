package main 

import (

	"os"
	"fmt"
	"log"
)


var ( 

	fileInfo os.FileInfo
	err error 
) 

func getData(file string) {

	
	fileInfo, err  = os.Stat(file)

	if err != nil {

		log.Fatalln(err)
	}

	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System infor: %+v\n\n", + fileInfo.Size())

}




