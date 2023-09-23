package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"bytes"
)

func encryptAES(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 패딩된 블록 크기로 plaintext를 확장합니다.
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	// IV 생성
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 암호화
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func decryptAES(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// IV 추출
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 복호화
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 패딩 제거
	ciphertext = PKCS7Unpadding(ciphertext)

	return ciphertext, nil
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS7Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func main() {
	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	plaintext := []byte("Hello, AES encryption!")

	ciphertext, err := encryptAES(key, plaintext)
	if err != nil {
		fmt.Println("암호화 중 오류 발생:", err)
		return
	}

	decryptedText, err := decryptAES(key, ciphertext)
	if err != nil {
		fmt.Println("복호화 중 오류 발생:", err)
		return
	}

	fmt.Println("암호문(hex):", hex.EncodeToString(ciphertext))
	fmt.Println("복호화된 텍스트:", string(decryptedText))
}
