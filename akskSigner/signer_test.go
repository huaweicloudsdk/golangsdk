package signer

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestSign(test *testing.T) {
	req, _ := http.NewRequest("GET", "http://www.test.com?a=b", nil)

	signOptins := SignOptions{
		AccessKey: "abc",
		SecretKey: "123",
	}
	Sign(req, signOptins)

	if req.Header.Get("Authorization") == "" {
		test.Error("No header 'Authorization' found")
	}

	if req.Header.Get("X-Sdk-Date") == "" {
		test.Error("No header 'X-Sdk-Date' found")
	}

	if req.Header.Get("Host") != "www.test.com" {
		test.Error("Wrong header(Host) value for AK/SK auth")
	}

	for k, v := range req.Header {
		fmt.Printf("%s:%s\n", k, v[0])
	}
}

func TestStringSort(test *testing.T) {
	keys := []string{"A1", "c2", "f1", "D1", "A2"}
	caseInsensitiveSort(keys)

	if keys[1] != "A2" {
		log.Fatal()
	}

	if keys[4] != "f1" {
		test.Fail()
	}
}

func TestHmacSha256(test *testing.T) {
	secret := []byte("the shared secret key here")
	message := "the message to hash here"

	result := hex.EncodeToString(HmacSha256(message, secret))

	if result != "4643978965ffcec6e6d73b36a39ae43ceb15f7ef8131b8307862ebc560e7f988" {
		test.Fail()
	}
}

func TestHashSha256(test *testing.T) {
	message := "the message to hash here"

	result := hex.EncodeToString(HashSha256([]byte(message)))

	if result != "983564913cd4151d38b1af858da66c653658fcacdc1866134e915b60aded1e78" {
		test.Fail()
	}
}
