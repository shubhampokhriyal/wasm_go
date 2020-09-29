package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"syscall/js"
)

func hash(str string) string {
	// Step 1: Convert to byte
	bS := []byte(str)

	// Step 2: Hash
	h := sha1.New()
	h.Write(bS)
	sha := hex.EncodeToString(h.Sum(nil))

	// Step 3: Return
	return sha
}

func hashGo(this js.Value, i []js.Value) interface{} {
	// Step 1: Assert
	str := i[0].String()

	// Step 2: Get the hash
	return hash(str)
}

func verifyGo(this js.Value, i []js.Value) interface{} {
	// Step 1: Assert
	prevHash := i[0].String()
	str := i[1].String()

	// Step 2: Hash
	newHash := hash(str)

	// Step 3: Return
	return prevHash == newHash
}

func registerCallbacks() {
	js.Global().Set("hash", js.FuncOf(hashGo))
	js.Global().Set("verify", js.FuncOf(verifyGo))
}

func main() {
	c := make(chan interface{})
	fmt.Println("Registering callbacks")
	registerCallbacks()
	<-c
}
