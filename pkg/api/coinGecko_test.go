package api

import (
	"testing"
)

func TestList(t *testing.T) {
	cg := CoinGecko{}
	_, err := cg.List()

	if err != nil {
		t.Errorf("call to List failed!")
	}
}
