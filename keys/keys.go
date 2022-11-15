package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	//8-, 16-, 32-, 64-, 128-, 256-, 512- , 1024-, 2048-, 4096-bits sequence
	pow := []int16{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

	// possible number of different keys for 8,16,32...bits
	for i := 0; i < len(pow); i++ {
		total := math.Pow(2, float64(pow[i]))
		fmt.Printf("Total possible key values for %d bits are %v\n", pow[i], total)
	}

	// SECTION FOR 8 bits
	// generate hex key
	// 4 bits for 1 hex symbol (numbers from 1 to 16 in binary 0000-1111, in hex - 0x00-0xFF)
	key8Bit := generateRandomHexKey(pow[0] / 4)
	fmt.Println("Generated 8-bits key = ", key8Bit)
	keyList8 := saveAndPrintAll8BitsKeys()
	find8BitKey(key8Bit, keyList8)

	// SECTION FOR 16 bits
	key16Bit := generateRandomHexKey(pow[1] / 4)
	fmt.Println("Generated 16-bits key = ", key16Bit)

	keyList16 := saveAll16BitsKeys()
	find16BitKey(key16Bit, keyList16)

	// SECTION FOR 32 bits
	//key32Bit := "00013E1B"
	key32Bit := generateRandomHexKey(pow[2] / 4)
	fmt.Printf("Generated 32-bits key = %v\n", key32Bit)
	//time consuming - several hours, uncomment if needed
	//find32BitKey(key32Bit)

}

func generateRandomHexKey(length int16) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%X", b)[:length]
}

func saveAndPrintAll8BitsKeys() [255]string {
	var UINT8_MIN uint8 = uint8(0)
	const UINT8_MAX = ^uint8(0)
	var keys [UINT8_MAX]string
	i := UINT8_MIN

	for {
		keys[i] = fmt.Sprintf("%02X", i)
		fmt.Printf("Key %d = %s\n", i, keys[i])
		i++
		if i == UINT8_MAX {
			keys[i-1] = fmt.Sprintf("%02X", i)
			fmt.Printf("Key %d = %s\n", i, keys[i-1])
			break
		}
	}
	return keys
}

func saveAll16BitsKeys() [65535]string {
	var UINT16_MIN uint16 = uint16(0)
	const UINT16_MAX = ^uint16(0)
	var keys [UINT16_MAX]string
	i := UINT16_MIN

	for {
		keys[i] = fmt.Sprintf("%04X", i)
		//fmt.Printf("Key %d = %s\n", i, keys[i])
		i++
		if i == UINT16_MAX {
			keys[i-1] = fmt.Sprintf("%04X", i)
			//fmt.Printf("Key %d = %s\n", i, keys[i-1])
			break
		}
	}
	return keys
}

func find8BitKey(key string, array [255]string) {
	defer timeTracker(time.Now(), "Find the key...")
	for i := 0; i < len(array)-1; i++ {
		if key == array[i] {
			fmt.Printf("Generated key in array is found %s = %s\n", key, array[i])
		}
	}
}

func find16BitKey(key string, array [65535]string) {
	defer timeTracker(time.Now(), "Find the key...")
	for i := 0; i < len(array)-1; i++ {
		if key == array[i] {
			fmt.Printf("Generated key in array is found %s = %s\n", key, array[i])
		}
	}
}

func find32BitKey(key string) {
	defer timeTracker(time.Now(), "Find the key...")

	for ii := 0; ii < 4294967295; ii++ {
		if key == strings.ToUpper(strconv.FormatInt(int64(ii), 16)) {
			fmt.Printf("Key is found %s = %v\n", key, ii)
		} //else {
		//fmt.Printf("Key %s, index %v\n", key, strings.ToUpper(strconv.FormatInt(int64(ii), 16)))
		//}
	}
}

func timeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// There is no sence try to brute force 64-, 128-, 256-, 512- , 1024-, 2048-, 4096-bits sequence
