// CREDS -> https://github.com/gtank/cryptopasta
// cryptopasta - basic cryptography examples
//
// Written in 2015 by George Tankersley <george.tankersley@gmail.com>
//
// To the extent possible under law, the author(s) have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
//
// You should have received a copy of the CC0 Public Domain Dedication along
// with this software. If not, see // <http://creativecommons.org/publicdomain/zero/1.0/>.

// Provides symmetric authenticated encryption using 256-bit AES-GCM with a random nonce.
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

// NewEncryptionKey generates a random 256-bit key for Encrypt() and
// Decrypt(). It panics if the source of randomness fails.
func NewEncryptionKey() *[32]byte {
	key := [32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return &key
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

func DecryptFile(path string, key *[32]byte) error {
	data, err := readFile(path)
	if err != nil {
		return err
	}

	decrypted, err := Decrypt(data, key)
	if err != nil {
		return err
	}

	err = writeFile(path, decrypted)
	if err != nil {
		return err
	}
	return nil
}

func EncryptFile(path string, key *[32]byte) error {
	data, err := readFile(path)
	if err != nil {
		return err
	}

	encrypted, err := Encrypt(data, key)
	if err != nil {
		return err
	}

	err = writeFile(path, encrypted)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(path string, data []byte) error {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func readFile(path string) (data []byte, e error) {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// get the size
	fileSize := fi.Size()

	// Read file, line by line
	var fileData []byte

	_, err = file.Read(fileData)

	if err != nil {
		return nil, err
	}

	return fileData, nil
}
