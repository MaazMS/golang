package main

import "fmt"

func main() {

	m := map[string]int{
		"Maaz":   24,
		"Shaikh": 23,
	}
	fmt.Println(m["Maaz"])
	fmt.Println(m["Shaikh"])

	//  key and value which is not available in map using v , ok idiom
	fmt.Println("key and value which is not available in map using v , ok idiom")
	v, ok := m["Noksy"]
	fmt.Println(v)
	fmt.Println(ok)

	// 	key and value which is not available in map using v , ok idiom and if condition
	fmt.Println("key and value which is not available in map using v , ok idiom and if condition")
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	addKeyValue()
	deleteKey()

}

func addKeyValue() {
	m := map[string]int{
		"Maaz":   24,
		"Shaikh": 23,
	}

	m["user"] = 500

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func deleteKey() {

	m := map[string]int{
		"Maaz":   24,
		"Shaikh": 23,
	}
	fmt.Println("Before delete map key ")
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "Shaikh")
	fmt.Println("\nAfter delete map key ")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
