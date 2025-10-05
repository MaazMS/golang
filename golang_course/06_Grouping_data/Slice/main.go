package main

import "fmt"

func main() {

	//  variable := type {values}

	no := []int{1, 2, 3, 5, 7, 8, 9}
	fmt.Println(no)
	//iterates(no)
	//slicingslice(no)
	//sliceAppend(no)
	//deleteInslice(no)
	//slicemake()
	multiDslice()
}

func iterates(no []int) {
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(no[0])
	fmt.Println(no[1])
	fmt.Println(no[2])
	fmt.Println(no[3])
	fmt.Println(no[4])

	for i, v := range no {
		fmt.Println(i, v)
	}

}

func slicingslice(no []int) {

	fmt.Println(no)
	fmt.Printf("Access index by value %d\n", no[1])
	fmt.Printf("slice a slice from index 1 include to end of slice %d\t\n", no[1:])
	fmt.Printf("Access index by value %d\n", no[4])
	fmt.Printf("slice a slice from index 1 include to exclude last index of  slice %d\t\n", no[1:4])

}

func sliceAppend(no []int) {

	fmt.Println(no)
	no = append(no, 11, 12, 13)
	fmt.Println(no)

	even := []int{2, 4, 6, 8, 10}
	no = append(no, even...)
	fmt.Println(no)

}

func deleteInslice(no []int) {
	fmt.Printf("Original slice\n")
	for i := 0; i < len(no); i++ {
		fmt.Printf("%d\t", no[i])
	}
	fmt.Println()
	fmt.Printf("deleteInslice \n")

	no = append(no[:2], no[:4]...)

	for i := 0; i < len(no); i++ {
		fmt.Printf("%d\t", no[i])
	}
}

func slicemake() {

	no := make([]int, 10, 100)
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(cap(no))
	no[9] = 99

	no = append(no, 15)
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(cap(no))

}

func multiDslice() {

	m1 := []string{"Maaz", "Shaikh"}
	m2 := []string{"golang", "python "}

	fmt.Println(m1)
	fmt.Println(m2)

	m3 := [][]string{m1, m2}
	fmt.Println(m3)

}
