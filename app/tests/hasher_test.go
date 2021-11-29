package tests

import (
	"URLShortener/app/hasher"
	"testing"
)

func TestEncode(t *testing.T){
	hash, err := hasher.Encode()
	if err != nil{
		t.Errorf("Error appeared")
	}
	if len(hash) != 10{
		t.Errorf("Wrong length of the hash")
	}
}