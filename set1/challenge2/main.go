package main

import "encoding/hex"
import "fmt"
import "log"

func main(){
	var hex_encode_1 = "1c0111001f010100061a024b53535009181c"
	var hex_encode_2 = "686974207468652062756c6c277320657965"

	decoded_hex_1, err := hex.DecodeString(hex_encode_1)
	if err != nil{
		log.Fatal(err)
	}

	decoded_hex_2, err := hex.DecodeString(hex_encode_2)
	if err != nil{
		log.Fatal(err)
	}

	var result = make([]byte,0)

	for x:= 0; x<len(decoded_hex_1); x++ {
		result = append(result, decoded_hex_1[x] ^ decoded_hex_2[x])		
	}

	var encode_hex = hex.EncodeToString(result)
	fmt.Printf("%s\n", encode_hex)
}
