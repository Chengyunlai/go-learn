package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "lo"))               // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	//If n < 0, there is no limit on the number of replacements.所以1就是换一个，-1就是换全部
	fmt.Println(strings.Replace(a, "l", "L", 1)) // heLlo
	fmt.Println(strings.Split("a,b,c", ","))     // [a b c]
	fmt.Println(strings.ToLower(a))              // hello
	fmt.Println(strings.ToUpper(a))              // HELLO
	fmt.Println(len(a))                          // 5

}
