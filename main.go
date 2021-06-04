package main

import (
	"fmt"
	"github.com/sbani/xss-go-project/xss"
)

// Wrapper <img src=a onerror="alert(1)">
func Wrapper(s string) {
	fmt.Println(s)
	xss.Run()
}

// main <img src=a onerror="alert(1)">
func main() {
	Wrapper(`<script>alert(1)</script>`)
}