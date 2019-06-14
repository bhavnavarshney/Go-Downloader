package main

import (
	"fmt"
	"github.com/user/zip"
)

func main() {
	fmt.Println("********* Welcome to GO DOWNLOADER *******")
	//reader := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Println("Enter the number of files to be zipped")
	fmt.Scanln(&n)
	fmt.Println("Enter full path of the files")
	files := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		files[i] = s
	}
	output := "Archived_Done.zip"
	if err := zip.ZipAllFiles(output,files); err!=nil{
		panic(err)
	}
	fmt.Println("Zipped File : ",output)

}
