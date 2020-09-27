package main

import (
	f "fmt"

	morestr "github.com/kamitom/hello/morestrings"
)

func main() {
	f.Println("hello, go modules")
	f.Println("hello, 東京走著瞧!")
	f.Println("hello, https://golang.org/doc/code.html !")
	f.Println("hola, https://youtu.be/3ydlmgwZBbg !")

	string_must_reverse := "!!123 !motimaK rehpog ,aloH"

	f.Println(morestr.ReverseRunes(string_must_reverse))

}
