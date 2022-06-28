package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Enter size of n*n matrix: ")
	var size int
	fmt.Scanln(&size)
	check := math.Log2(float64(size))
	check = math.Ceil(check)
	padding := int(math.Pow(2, check)) - size // needed to make the size of matrix a power of 2 so that it can be divided by 2 on and on
	var m1 = make([][]int, size+padding)
	for i := range m1 {
		m1[i] = make([]int, size+padding)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("Enter %d,%d element of first matrix: ", i, j)
			fmt.Scanf("%d", &m1[i][j])
		}
	}
	// fill padding with zero vals
	fmt.Println("Your first matrix is: ", m1)

	var m2 = make([][]int, size+padding)
	for i := range m2 {
		m2[i] = make([]int, size+padding)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("Enter %d,%d element of second matrix: ", i, j)
			fmt.Scanf("%d", &m2[i][j])
		}
	}
	fmt.Println("Your second matrix is: ", m2)

	m3 := multiply(m1, m2)
	var result = make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[i][j] = m3[i][j]
		}
	}

	fmt.Println("Resultant matrix:", result)
}

// Strassen algo
// X = [[A B] [C D]] and Y = [[E F] [G H]]
// P1 = A(F-H)
// P2 = (A+B)H
// P3 = (C+D)E
// P4 = D(G-E)
// P5 = (A+D)(E+H)
// P6 = (B-D)(G+H)
// P7 = (A-C)(E+H)
// X*Y = [[P5+P4-P2+P6 P1+P2][P3+P4 P1+P5-P3-P7]]
func multiply(m1, m2 [][]int) [][]int {
	n := len(m1)
	if n == 0 {
		return [][]int{{0}}
	}
	if n == 1 {
		return [][]int{{m1[0][0] * m2[0][0]}}
	}
	var m3 = make([][]int, n)
	for i := range m3 {
		m3[i] = make([]int, n)
	}
	A := leftHalf(m1[:n/2])
	B := rightHalf(n, m1[:n/2])
	C := leftHalf(m1[n/2:])
	D := rightHalf(n, m1[n/2:])
	E := leftHalf(m2[:n/2])
	F := rightHalf(n, m2[:n/2])
	G := leftHalf(m2[n/2:])
	H := rightHalf(n, m2[n/2:])

	P1 := multiply(A, subtract(F, H))
	P2 := multiply(add(A, B), H)
	P3 := multiply(add(C, D), E)
	fmt.Println("P3:", P3)
	P4 := multiply(D, subtract(G, E))
	fmt.Println("P4:", P4)
	P5 := multiply(add(A, D), add(E, H))
	P6 := multiply(subtract(B, D), add(G, H))
	P7 := multiply(subtract(A, C), add(E, H))

	R1 := add(subtract(add(P5, P4), P2), P6)
	fmt.Println("R1:", R1)
	R2 := add(P1, P2)
	fmt.Println("R2:", R2)
	R3 := add(P3, P4)
	fmt.Println("R3:", R3)
	R4 := subtract(subtract(add(P1, P5), P3), P7)
	fmt.Println("R4:", R4)

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			m3[i][j] = R1[i][j]
			m3[i][j+(n/2)] = R2[i][j]
			m3[i+(n/2)][j] = R3[i][j]
			m3[i+(n/2)][j+(n/2)] = R4[i][j]
		}
	}
	return m3
}

func add(m1, m2 [][]int) [][]int {
	var m3 = make([][]int, len(m1[0]))
	for i := range m3 {
		m3[i] = make([]int, len(m1[0]))
	}
	for i := range m3 {
		for j := range m3 {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}

	return m3
}

func subtract(m1, m2 [][]int) [][]int {
	var m3 = make([][]int, len(m1[0]))
	for i := range m3 {
		m3[i] = make([]int, len(m1[0]))
	}
	for i := range m3 {
		for j := range m3 {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}

	return m3
}

func leftHalf(m1 [][]int) [][]int {
	var m3 = make([][]int, len(m1))
	for i := range m3 {
		m3[i] = make([]int, len(m1))
	}
	for i := range m3 {
		for j := range m3 {
			m3[i][j] = m1[i][j]
		}
	}
	return m3
}

func rightHalf(size int, m1 [][]int) [][]int {
	var m3 = make([][]int, len(m1))
	for i := range m3 {
		m3[i] = make([]int, len(m1))
	}
	for i := range m3 {
		for j := range m3 {
			m3[i][j] = m1[i][j+len(m3)]
		}
	}
	return m3
}
