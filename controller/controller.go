package controller

import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/common/hexutil"
import "crypto/rand"
import "math/big"

func GetRandomMessage(length int) (string, error) {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterLength := new(big.Int).SetInt64(int64(len(letters)))

	message := make([]byte, length)
	for i := range message {
		index, err := rand.Int(rand.Reader, letterLength)
		if err != nil {
			return string(message), err
		}
		message[i] = letters[index.Int64()]
	}
	return string(message), nil
}
func Verify(address string, message string, signature string) (bool, error) {
	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return false, nil
	}
	messageHashBytes := crypto.Keccak256Hash([]byte(message)).Bytes()
	publicKey, err := crypto.SigToPub(messageHashBytes, signatureBytes)
	if err != nil {
		return false, nil
	}

	signatureAddress := crypto.PubkeyToAddress(*publicKey).Hex()

	return signatureAddress == address, nil
}
