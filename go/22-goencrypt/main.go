package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func getInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("Empty message is not allowed")
	}
	return text, nil
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Work(secret string, isDecrypt bool, text string) {
	if !isDecrypt {
		encText, err := Encrypt(text, secret)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
		}
		fmt.Println(encText)
	} else {
		decText, err := Decrypt(text, secret)
		if err != nil {
			fmt.Println("error decrypting your encrypted text: ", err)
		}
		fmt.Println(decText)
	}
}

func main() {

	fScret := flag.String("secret", "abc&1*~#^2^#s0^=)^^7%b34", "secret for encrypt and descrypt password")
	fIsDecrypt := flag.Bool("isDecrypt", false, "Choosee descrypt or encrypt")
	flag.Parse()

	text, err := getInput(os.Stdin, flag.Args()...)

	if err != nil {
		fmt.Println("faile read input: ", err)
	}
	Work(*fScret, *fIsDecrypt, text)

}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string, secret string) (string, error) {

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
