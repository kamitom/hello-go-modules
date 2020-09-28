package main

import (
	f "fmt"

	morestr "github.com/kamitom/hello/morestrings"

	cmp "github.com/google/go-cmp/cmp"

	quote "rsc.io/quote"

	fk "syreclabs.com/go/faker"
)

func main() {
	f.Println("hello, go modules2")
	f.Println("hello, 東京走著瞧!")
	f.Println("hello, https://golang.org/doc/code.html !")
	f.Println("hola, https://youtu.be/3ydlmgwZBbg !")

	string_must_reverse := "!!123 !motimaK rehpog ,aloH"
	// string_must_reverse2 := "台灣"

	f.Println(morestr.ReverseRunes(string_must_reverse))
	f.Println(cmp.Diff("hello world", "hello go"))

	f.Println(quote.Opt())
	f.Println(quote.Glass())
	f.Println(quote.Go())

	f.Println("Hola,", fk.Name().FirstName())

}
