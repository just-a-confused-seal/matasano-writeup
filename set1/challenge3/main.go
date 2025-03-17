package main

import "fmt"
import "encoding/hex"
import "log"
import "slices"

func prerequisite(decrypted_text []byte) bool{

	var total_word = len(decrypted_text)

	for x:=0; x < total_word; x++{
		if (32 <= int(decrypted_text[x]) && int(decrypted_text[x]) <= 127) {
			continue
		}else{
			return false
		}	
	}
	return true
}

func chi_squared_calc(decrypted_text []byte) int{
	var tmp_convert = make([]byte, 0)

	charFreq := map[int]float64{
		32: 13.00,
		97: 8.17,  // 'a'
		98: 1.49,  // 'b'
		99: 2.78,  // 'c'
		100: 4.25, // 'd'
		101: 12.70,// 'e'
		102: 2.23, // 'f'
		103: 2.02, // 'g'
		104: 6.09, // 'h'
		105: 6.97, // 'i'
		106: 0.15, // 'j'
		107: 0.77, // 'k'
		108: 4.03, // 'l'
		109: 2.41, // 'm'
		110: 6.75, // 'n'
		111: 7.51, // 'o'
		112: 1.93, // 'p'
		113: 0.10, // 'q'
		114: 5.99, // 'r'
		115: 6.33, // 's'
		116: 9.06, // 't'
		117: 2.76, // 'u'
		118: 0.98, // 'v'
		119: 2.36, // 'w'
		120: 0.15, // 'x'
		121: 1.97, // 'y'
		122: 0.07, // 'z'
	}


	//normalize the uppercase character to lowercase in byte type data
	for x:=0; x < len(decrypted_text); x++{
		if (65 <= int(decrypted_text[x]) && int(decrypted_text[x]) <= 90){
			tmp_convert = append(tmp_convert, byte(int(decrypted_text[x] + 32)))
		}else{
			tmp_convert = append(tmp_convert, decrypted_text[x])
		}
	}

	chi_squared_total := 0.00
	for x:=0; x< len(tmp_convert); x++{
		total_char := 0
		for y:=x; y< len(tmp_convert); y++{
			if tmp_convert[x] == tmp_convert[y]{
				total_char += 1
				slices.Delete(tmp_convert,y,y)
				
			}
		}
		//fmt.Printf("[%s -- %d]",string(tmp_convert[x]),total_char)
		expected_number_char := charFreq[int(tmp_convert[x])] * float64(len(tmp_convert))

		chi_squared := ((expected_number_char - float64(total_char)) * (expected_number_char - float64(total_char))) / float64(total_char)
		chi_squared_total += chi_squared

	}

	return int(chi_squared_total)

}

func main(){
	var encrypted_hex = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded_enc_hex, err := hex.DecodeString(encrypted_hex)
	if err != nil{
		log.Fatal(err)
	}

	var result = make([]byte, 0)
	var tmp_cs = 0
	var chi_squared = 0
	var possible_decrypted_text = ""

	for i:= 0; i < 255; i++{
		for j:= 0; j < len(decoded_enc_hex); j++{
			result = append(result, decoded_enc_hex[j] ^ byte(i))
		}

		if prerequisite(result){
			//fmt.Printf("%s\n", string(result))
			tmp_cs = chi_squared_calc(result)
			if tmp_cs > chi_squared{
				chi_squared = tmp_cs
				possible_decrypted_text = string(result)
			}
		
		}
		
		result = result[:0]
	}

	
	fmt.Printf("%s\n", possible_decrypted_text)
	
}