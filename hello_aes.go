package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

func main() {
    originalText := "encrypt this golang"
    fmt.Println(originalText)

    key := []byte("example key 1234")

    // encrypt value to base64
    cryptoText := encrypt(key, originalText)
    fmt.Println(cryptoText)

    // encrypt base64 crypto to original value
    text := decrypt(key, cryptoText)
    fmt.Printf(text)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
    // key := []byte(keyText)
    plaintext := []byte(text)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    // The IV needs to be unique, but not secure. Therefore it's common to
    // include it at the beginning of the ciphertext.
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    // convert to base64
    return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
    ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    // The IV needs to be unique, but not secure. Therefore it's common to
    // include it at the beginning of the ciphertext.
    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)

    // XORKeyStream can work in-place if the two arguments are the same.
    stream.XORKeyStream(ciphertext, ciphertext)

    return fmt.Sprintf("%s", ciphertext)
}

func main2() {
    //The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
    key := "opensesame123456" // 16 bytes!

    block, err := aes.NewCipher([]byte(key))

    if err != nil {
        panic(err)
    }

    fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)

    str := []byte("Hello World!")

    // 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
    ciphertext := []byte("abcdef1234567890")
    iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

    // encrypt

    encrypter := cipher.NewCFBEncrypter(block, iv)

    encrypted := make([]byte, len(str))
    encrypter.XORKeyStream(encrypted, str)

    fmt.Printf("%s encrypted to %v\n", str, encrypted)

    // decrypt

    decrypter := cipher.NewCFBDecrypter(block, iv) // simple!

    decrypted := make([]byte, len(str))
    decrypter.XORKeyStream(decrypted, encrypted)

    fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)

}
