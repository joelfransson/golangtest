package main

import "fmt"

type price interface {
	Price() float64
	VatPrice() float64
}

type articleInfo struct {
	price float64
}

type book struct {
	articleInfo
	title string
}
type computer struct {
	articleInfo
	processor string
}

// Interface Impl

func (a articleInfo) Price() float64 {
	return a.price
}

func (a articleInfo) VatPrice() float64 {
	return a.price * 1.25
}

func (a book) VatPrice() float64 {
	return a.price * 1.10
}

// App logic

func writePrice(o price) string {
	return fmt.Sprintf("price %f inkl.VAT %f", o.Price(), o.VatPrice())
}

func main() {
	c := computer{processor: "i7", articleInfo: articleInfo{price: 10000}}
	b := book{title: "A book", articleInfo: articleInfo{price: 50}}

	fmt.Println(writePrice(c))
	fmt.Println(writePrice(b))
}
