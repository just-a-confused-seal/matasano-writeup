package main

import "fmt"
import "encoding/hex"
import "log"
import "encoding/base64"

func main(){
	const str = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decoded_hex, err := hex.DecodeString(str)

	if err != nil{
		log.Fatal(err)
	}

	var base64_encoded = base64.StdEncoding.EncodeToString(decoded_hex)
	fmt.Println(base64_encoded)

}
