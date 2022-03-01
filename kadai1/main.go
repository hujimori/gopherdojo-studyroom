package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

var srcPath = flag.String("src", "", "filepath")
var dstPath = flag.String("dst", "", "filepath")

func main() {
	flag.Parse()

	srcFile, err := os.Open(*srcPath)
	if err != nil {
		fmt.Printf("Not found file : %s\n", *srcPath)
		return
	}
	defer srcFile.Close()

	img, err := jpeg.Decode(srcFile)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Failed to decode file to image : %s\n", *srcPath)
		return
	}

	dstFile, err := os.Create(*dstPath)
	if err != nil {
		fmt.Printf("Failed to create file : %s\n", dstFile)
		return
	}

	defer dstFile.Close()
	err = png.Encode(dstFile, img)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
