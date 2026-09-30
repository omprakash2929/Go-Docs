package main

import (
	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("./example.txt")

	// if err != nil {
	// 	//? log the error
	// 	panic(err)
	// }

	// defer f.Close()

	//* File Info

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//? log the error
	// 	panic(err)
	// }
	// fmt.Println("File name:", fileInfo.Name())
	// fmt.Println("File size:", fileInfo.Size())
	// fmt.Println("File time:", fileInfo.ModTime())
	// fmt.Println("File sys:", fileInfo.Sys())

	//* Read Files

	// buf := make([]byte, 11)
	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	//* Second way to read files

	// data, err := os.ReadFile("./example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//* Folder read

	// dir, err := os.Open(".")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir(), fi.Type())
	// }

	//* Create File

	// f, err := os.Create("test.txt") //? create files
	// if err != nil {
	// 	panic(err)

	// }
	// defer f.Close()

	// f.WriteString("Hi Go")     //? Write Files
	// f.WriteString("Hi Python") //? Write Files

	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	//! Read and write onther file (streaming fashion)

	// sourceFile, err := os.Open("./example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Written to new File sucessfully")

	//? Delete File

	err := os.Remove("./example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted")

}
