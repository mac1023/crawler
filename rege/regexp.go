package rege

import (
	"fmt"
	"regexp"
)

func Rege() {

	text := "my email is ccmouse@gmail.com  a@qq.com  b@163.com"
	re := regexp.MustCompilePOSIX(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)`)
	fmt.Println(re)

	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
