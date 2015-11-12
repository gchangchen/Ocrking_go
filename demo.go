package main

import (
	"./ocrking"
	"fmt"
	"log"
	"os"
)

func main() {
	path, _ := os.Getwd()
	path += "/code.jpg"
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	ocrking := &ocrking.Ocrking{
		Service:  "OcrKingForNumber",
		Language: "eng",
		Type:     "http://xxxxxxxxxxxxxxxxxxxxxxxxxx.com.cn/validateCode",
		Charset:  "7",
		Apikey:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	result, err := ocrking.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
