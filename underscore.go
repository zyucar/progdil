package main

import (
	"fmt"
	"strings"
)

func ReplaceUnderscoreWithSpace(str string) string {
        list := strings.Split(str, "")

        for i := 1 ; i < len(list)-1 ; i++ {
		if list[i] == "_" {
			list[i] = " "
                }
        }
        return strings.Join(list, "")
}
    
func main() {
	var word string
	fmt.Println("Bir kelime giriniz.")
	fmt.Scanf("%s", &word)
	fmt.Println(ReplaceUnderscoreWithSpace(word))
}

