package security

import (
	"crypto/rand"
	"fmt"
)

//var privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)

func GenerateToken() (string, error) {
	randomBytes := make([]byte, 16)
	rand.Reader.Read(randomBytes)
	randomNumberString := fmt.Sprintf("%x", randomBytes)
	// TODO sign token
	//hashAlgorithm := crypto.SHA1.New()
	//hashAlgorithm.Write([]byte(randomNumberString))
	//hash := hashAlgorithm.Sum(nil)
	//signed, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	//if err != nil {
	//	return "", nil
	//}

	//randomNumberString += randomNumberString + "_" + string(signed)

	return randomNumberString, nil
}

func IsTokenValid(token string) (bool, error) {
	return true, nil
}
