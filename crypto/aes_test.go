package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const KEY = "214111111113ptLJq3JrZMCMMNNdL93e"

func TestEncryptDecrypt(t *testing.T) {
	encryptTxt, err := Encrypt(KEY, "hoank")
	assert.Nil(t, err)
	// Cipher key must be 32 chars long because block size is 16 bytes
	val, err := Decrypt(KEY, encryptTxt)
	assert.Nil(t, err)
	assert.Equal(t, "hoank", val)
}
