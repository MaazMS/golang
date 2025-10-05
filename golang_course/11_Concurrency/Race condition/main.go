package main

import "runtime"

func main() {

	go say("World")
	say("hello")
}

func say(s string) {

	for i := 0; i <= 3; i++ {
		runtime.Gosched()
		println(i, s)
	}
}

/* output without Gosched
0 hello
1 hello
2 hello
3 hello
*/

/* output with Gosched
0 hello
0 World
1 hello
1 World
2 hello
2 World
3 hello

*/
