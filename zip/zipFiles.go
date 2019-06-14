package zip

import (
    "archive/zip"
	"os"
	"io"
)

func ZipAllFiles(filename string,files []string) error {
    newZipFile,error := os.Create(filename)
    if error !=nil{
		return error
	}
	defer newZipFile.Close()

	writer := zip.NewWriter(newZipFile)
	defer writer.Close()

	for _,file:= range files{
		error = AddFilesToZip(writer,file)
		if error!=nil{
			return error
		}
	}
	return nil
	} 

	func AddFilesToZip(zipFile *zip.Writer,file string)error{
		fileToZip,error:= os.Create(file)
		if error!=nil{
			return error
		}
		defer fileToZip.Close()

		fileInfo,error2 := fileToZip.Stat()
		if error2!=nil{
			return error2
		}

		header,error3 := zip.FileInfoHeader(fileInfo)
		if error3!=nil{
			return error3
		}
		//header.Name = file
		header.Method = zip.Deflate

		writer,error4 := zipFile.CreateHeader(header)
		if error4!=nil{
			return error4
		}
		_ , error4 = io.Copy(writer,fileToZip)
		return error4
	}