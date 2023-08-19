# Encrypto

![](https://img.shields.io/badge/Language-Golang-blue)
![](https://img.shields.io/badge/Tests-Pass-green)
![GitHub release (with filter)](https://img.shields.io/github/v/release/amirhnajafiz/encrypto)

A Golang library for encryption and decryption based on ```AES``` algorithm.
With this library you can encrypt and decrypt your texts by using a 16 bits private key.

## how to use?

Import the library in ```go.mod``` by using ```go get github.com/amirhnajafiz/encrypto```.
After that you can use it like the following example:

```go
package main

import (
	"fmt"

	"github.com/amirhnajafiz/encrypto"
)

func main() {
	key := "0123456789abcdef" // must be of 16 bytes for this example to work
	message := "Lorem ipsum dolor sit amet"

	encrypted, _ := encrypto.Encrypt(key, message)
	decrypted, _ := encrypto.Decrypt(key, encrypted)

	fmt.Println(key, message, encrypted, decrypted)
}
```
