package main

import (
	"fmt"
	"sync"
)

// Operation adalah interface yang mendefinisikan operasi matematika
type Operation interface {
	Calculate() int
}

// Addition adalah struct yang mengimplementasikan interface Operation untuk operasi penjumlahan
type Addition struct {
	a, b int
}

// Calculate mengembalikan hasil penjumlahan dari a dan b
func (add Addition) Calculate() int {
	return add.a + add.b
}

// Multiplication adalah struct yang mengimplementasikan interface Operation untuk operasi perkalian
type Multiplication struct {
	a, b int
}

// Calculate mengembalikan hasil perkalian dari a dan b
func (mul Multiplication) Calculate() int {
	return mul.a * mul.b
}

func main() {
	var wg sync.WaitGroup

	add := Addition{a: 3, b: 4}
	mul := Multiplication{a: 3, b: 4}

	// Simpan fungsi performOperation dalam variabel
	performOperation := func(op Operation, description string, wg *sync.WaitGroup) {
		defer wg.Done()
		result := op.Calculate()
		fmt.Println(description, result)
	}

	wg.Add(2)
	go performOperation(add, "Hasil penjumlahan:", &wg)
	go performOperation(mul, "Hasil perkalian:", &wg)

	wg.Wait()
}