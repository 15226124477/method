package main

import (
	"github.com/15226124477/method"
	"github.com/15226124477/start"
)

func main() {

	start.RotateLogs(1, false)
	zipFile := method.UnZipfile{
		ZipFilePath:     "C:\\Users\\suh\\Desktop\\tools\\Hi-TestALL-Server\\dist\\case\\Hi-Survey-Test\\V1.0.20241010_T2.zip",
		ZipType:         ".zip",
		ZipOutputFolder: "C:\\Users\\suh\\Desktop\\tools\\Hi-TestALL-Server\\dist\\case\\Hi-Survey-Test",
	}
	zipFile.Unzip()
}
