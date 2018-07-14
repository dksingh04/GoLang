/*
  The code is to demonstrate the use of base64 stream encoder/decoder and encoding/decoding of given data.
*/

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Simple String data"
	var encodedData []byte

	// Use of Encoder and Decoder Stream API of base64
	byteBuffer := bytes.NewBuffer(encodedData)
	encoder := base64.NewEncoder(base64.StdEncoding, byteBuffer)
	encoder.Write([]byte(data))

	encodedStr := string(byteBuffer.Bytes())
	fmt.Println(encodedStr)
	encoder.Close()

	var decodedData = make([]byte, len(encodedStr))
	decoder := base64.NewDecoder(base64.StdEncoding, byteBuffer)
	decoder.Read(decodedData)

	fmt.Println(string(decodedData))

	// Encoding and decoding the given data
	encodedStr = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedStr)

	decodedByte, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		fmt.Println("Error in decoding the data!!")
	}
	fmt.Println(string(decodedByte))

}
