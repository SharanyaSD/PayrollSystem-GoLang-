package config

// secret key - sign to jwt
var Secret = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// )

// func GenerateRandomKey(length int) (string, error) {
// 	key := make([]byte, length)
// 	if _, err := rand.Read(key); err != nil {
// 		return "", err
// 	}
// 	return base64.RawURLEncoding.EncodeToString(key), nil
// }

// // The secret key used to sign the JWT
// var Secret string

// // init initializes Secret variable with a random key when the package is initialized
// func init() {
// 	const keyLength = 32 //256 bits
// 	var err error
// 	Secret, err = GenerateRandomKey(keyLength)
// 	if err != nil {
// 		panic(err)
// 	}
//}
