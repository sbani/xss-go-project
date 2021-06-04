package xss

import (
	"fmt"
	"github.com/pkg/errors"
)

// Run runs an XSS <img src=a onerror="alert(1)">
func Run() {
	err := errors.New(`My fake error <img src=a onerror="alert(1)">`)
	fmt.Println(err)
}