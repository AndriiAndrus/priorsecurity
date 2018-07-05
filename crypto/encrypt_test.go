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

package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptDecryptGCM(t *testing.T) {
	randomKey := &[32]byte{}
	_, err := io.ReadFull(rand.Reader, randomKey[:])
	if err != nil {
		t.Fatal(err)
	}

	gcmTests := []struct {
		plaintext []byte
		key       *[32]byte
	}{
		{
			plaintext: []byte("Hello, world!"),
			key:       randomKey,
		},
	}

	for _, tt := range gcmTests {
		ciphertext, err := Encrypt(tt.plaintext, tt.key)
		if err != nil {
			t.Fatal(err)
		}

		plaintext, err := Decrypt(ciphertext, tt.key)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(plaintext, tt.plaintext) {
			t.Errorf("plaintexts don't match")
		}

		ciphertext[0] ^= 0xff
		plaintext, err = Decrypt(ciphertext, tt.key)
		if err == nil {
			t.Errorf("gcmOpen should not have worked, but did")
		}
	}
}

func TestFileEncryption(t *testing.T) {
	key := NewEncryptionKey()

	path := "../test_file.gob"

	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	// get the size
	fileSize := fi.Size()

	err = EncryptFile(path, key)
	if err != nil {
		t.Fatal(err)
	}

	fi2, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Original file size: ", fileSize)
	t.Log("Encrypted file size: ", fi2.Size())

	err = DecryptFile(path, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTextEncryption(t *testing.T) {
	key := NewEncryptionKey()

	text := "Hello, World!"
	data := base64.StdEncoding.EncodeToString([]byte(text))

	encrypted, err := Encrypt([]byte(data), key)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encrypted string: ", string(encrypted))

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatal(err)
	}

	baseText, err := base64.StdEncoding.DecodeString(string(decrypted))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Decrypted string: ", string(baseText))
}

func BenchmarkAESGCM(b *testing.B) {
	randomKey := &[32]byte{}
	_, err := io.ReadFull(rand.Reader, randomKey[:])
	if err != nil {
		b.Fatal(err)
	}

	data, err := ioutil.ReadFile("testdata/big")
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(data)))

	for i := 0; i < b.N; i++ {
		Encrypt(data, randomKey)
	}
}
