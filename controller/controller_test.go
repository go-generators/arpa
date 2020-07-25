package controller

import "testing"
import "crypto/ecdsa"
import "github.com/stretchr/testify/assert"
import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/common/hexutil"

func TestGetRandomMessage(t *testing.T) {
	assert := assert.New(t)
	message, err := GetRandomMessage(10)
	assert.Nil(err)
	assert.Equal(len(message), 10)
}

func TestVerify(t *testing.T) {
	assert := assert.New(t)
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	assert.Nil(err)
	message := "randommessagetosign"
	messageHashBytes := crypto.Keccak256Hash([]byte(message)).Bytes()

	signatureBytes, err := crypto.Sign(messageHashBytes, privateKey)

	assert.Nil(err)
	signature := hexutil.Encode(signatureBytes)

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	assert.True(ok)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	verified, err := Verify("wrongaddress", "wrongmessage", "wrogsignature")
	assert.False(verified)
	assert.Nil(err)

	verified, err = Verify("wrongaddress", "wrongmessage", signature)
	assert.False(verified)
	assert.Nil(err)

	verified, err = Verify("wrongaddress", message, signature)
	assert.False(verified)
	assert.Nil(err)

	verified, err = Verify(address, message, signature)
	assert.True(verified)
	assert.Nil(err)
}
