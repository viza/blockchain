package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

func main() {
	// test vectors
	testVector1 := []byte("ff00000000000000000000000000000000000000000000000000000000000000")
	testVector2 := []byte("aaaa000000000000000000000000000000000000000000000000000000000000")
	testVector3 := []byte("FFFFFFFF")
	testVector4 := []byte("F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	fmt.Println("Please choose what do you want to convert...")
	fmt.Println("Enter 1 - to convert hex string to little endian")
	fmt.Println("Enter 2 - to convert hex string to big endian")
	fmt.Println("Enter 3 - to convert little endian to hex string")
	fmt.Println("Enter 4 - to convert big endian to hex string")

	var item int8
	fmt.Scan(&item)
	switch item {
	case 1: // LITTLE ENDIAN SECTION
		fmt.Println("1 - to convert hex string to little endian")
		byteArr := getHexByteArray() // data from console
		fmt.Println(" - - - - - ")
		fmt.Println("Little endian: ", byteArrayToInt(byteArr))
		fmt.Printf(" - - - - - \n\n")

		fmt.Println(" - - TEST VECTORS - - ")
		// Good online converter: https://www.scadacore.com/tools/programming-calculators/online-hex-converter/

		bin1 := makeHexByteArray(string(testVector1))
		fmt.Printf("Byte data: %v, little endian = %d\n\n", bin1, byteArrayToInt(bin1))

		bin2 := makeHexByteArray(string(testVector2))
		fmt.Printf("Byte data: %v, little endian = %d\n\n", bin2, byteArrayToInt(bin2))

		bin3 := makeHexByteArray(string(testVector3))
		fmt.Printf("Byte data: %v, little endian = %d\n\n", bin3, byteArrayToInt(bin3))

		bin4 := makeHexByteArray(string(testVector4))
		fmt.Printf("Byte data: %v, little endian = %d\n\n", bin4, byteArrayToInt(bin4))

	case 2:
		fmt.Println("2 - to convert hex string to big endian")

		byteArr := getHexByteArray() // data from console
		fmt.Println(" - - - - - ")

		fmt.Println("Big endian: ", byteArrayToInt(reverse(byteArr)))
		fmt.Printf(" - - - - - \n\n")

		bin1 := makeHexByteArray(string(testVector1))
		fmt.Printf("Byte data: %x, big endian = %d\n\n", bin1, bytesToInt64(bin1))

		bin2 := makeHexByteArray(string(testVector2))
		fmt.Printf("Byte data: %x, big endian = %d\n\n", bin2, byteArrayToInt(reverse(bin2)))

		bin3 := makeHexByteArray(string(testVector3))
		fmt.Printf("Byte data: %x, big endian = %d\n\n", bin3, byteArrayToInt(reverse(bin3)))

		bin4 := makeHexByteArray(string(testVector4))
		fmt.Printf("Byte data: %x, big endian = %d\n\n", bin4, byteArrayToInt(reverse(bin4)))

	case 3:
		fmt.Println("3 - to convert little endian to hex string")
		var numStr string
		fmt.Println("Enter number please. Allowed symbols are: 0123456789")
		fmt.Scan(&numStr)
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Panic("Invalid number")
		} else {

			fmt.Printf("Hex string: %x\n", reverse(Int64ToBytes(int64(n))))
		}

	case 4:
		fmt.Println("4 - to convert big endian to hex string")
		var numStr string
		fmt.Println("Enter number please. Allowed symbols are: 0123456789")
		fmt.Scan(&numStr)
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Panic("Invalid number")
		} else {

			fmt.Printf("Hex string: %x\n", Int64ToBytes(int64(n)))
		}

	default:
		fmt.Println("Please enter valid number: 1, 2, 3 or 4...")
	}

}

func checkHexString(err error) {
	if err != nil {
		fmt.Println("You entered invalid string...")
		log.Fatal(err)
	}
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}

func byteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}

func getHexByteArray() []byte {
	var hexStr string
	fmt.Println("Enter hex number as string please(without 0x). Allowed symbols are: 0123456789ABCDEF")
	fmt.Scan(&hexStr)

	// Convert hex representation into binary data (array of bytes)
	binByteStr := make([]byte, hex.DecodedLen(len(hexStr)))
	_, err := hex.Decode(binByteStr, []byte(hexStr))
	checkHexString(err) //check if entered hex string is OK
	fmt.Printf("Hex = 0x%s\nBinary = %b\n", hexStr, binByteStr)

	size := len(hexStr) / 2
	fmt.Printf("Number of needed bytes to store data is %d\n", size)
	return binByteStr
}

func makeHexByteArray(hexStr string) []byte {
	// Convert hex representation into binary data (array of bytes)
	binByteStr := make([]byte, hex.DecodedLen(len(hexStr)))
	_, err := hex.Decode(binByteStr, []byte(hexStr))
	checkHexString(err) //check if entered hex string is OK
	fmt.Printf("Hex = 0x%s\nBinary = %b\n", hexStr, binByteStr)

	size := len(hexStr) / 2
	fmt.Printf("Number of needed bytes to store data is %d\n", size)
	return binByteStr
}

func bytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
