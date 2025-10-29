package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func main () {
	var data tabInt
	var nData int
	var x1 int
	
	fmt.Scan(&x1)
	fmt.Scan(&nData)
	
	if nData > NMAX {
		nData = NMAX
	}
	
	bacaData(&data, nData)
	cetakData(data, nData)
	
	if sequentialSearch(data, nData, x1) == true {
		fmt.Printf("Hasil pencarian: Bilangan ditemukan. Terdapat %d bilangan %d.", frekuensiBilangan(data, nData, x1), x1)
	} else {
		fmt.Print("Hasil pencarian: Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n int) {
	var i int
	
	if n > NMAX {
		n = NMAX
	}
	
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	var i int
	
	fmt.Print("Data Bilangan: ")
	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	var i, count int
	
	count = 0
	for i = 0; i < n; i++ {
		if A[i] == x {
			count++
		}
	}
	return count
}

func sequentialSearch(A tabInt, n, x int) bool {
	var k int
	var found bool
	
	found = false
	k = 0
	for !found && k < n {
		found = A[k] == x
		k++
	}
	return found
}