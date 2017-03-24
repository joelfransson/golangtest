package main

import (
	"testing"
)

func TestWritePrice(t *testing.T) {
	c := computer{processor: "i7", articleInfo: articleInfo{price: 10000}}

	res := writePrice(c)

	if res != "price 10000.000000 inkl.VAT 12500.000000" {
		t.Errorf("FAIL %v", res)
	}

	b := book{title: "A Book", articleInfo: articleInfo{price: 10000}}

	res = writePrice(b)

	if res != "price 10000.000000 inkl.VAT 11000.000000" {
		t.Errorf("FAIL %v", res)
	}
}
