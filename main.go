package main

import (
	"fmt"
	"iter"
	"slices"

	"awesomeProject15/lq"
)

func main() {
	values := slices.Values([]int{2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Fluent API - encadenando operaciones
	fmt.Println("Números pares:")
	lq.From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		ForEach(func(n int) { fmt.Println(n) })

	fmt.Println("\nNúmeros mayores que 5:")
	lq.From(values).
		Filter(func(n int) bool { return n > 5 }).
		ForEach(func(n int) { fmt.Println(n) })

	fmt.Println("\nNúmeros pares multiplicados por 2:")
	lq.From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * 2 }).
		ForEach(func(n int) { fmt.Println(n) })

	fmt.Println("\nNúmeros pares como slice:")
	evenSlice := lq.From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		ToSlice()
	fmt.Println(evenSlice)

	// Ejemplo de FlatMap - expandir cada número a un rango
	fmt.Println("\nFlatMap - expandir cada número a un rango:")
	lq.From(values).
		Filter(func(n int) bool { return n <= 5 }).
		FlatMap(func(n int) iter.Seq[int] {
			return slices.Values([]int{n, n * 2, n * 3})
		}).
		ForEach(func(n int) { fmt.Println(n) })

	// Ejemplo de FlatMap - crear múltiples elementos para cada número
	fmt.Println("\nFlatMap - crear múltiples elementos:")
	lq.From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		FlatMap(func(n int) iter.Seq[int] {
			return slices.Values([]int{n, n + 1, n + 2})
		}).
		ForEach(func(n int) { fmt.Println(n) })

	// Ejemplos de Reduce
	fmt.Println("\nReduce - suma de todos los números:")
	sum := lq.From(values).Reduce(0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("Suma total: %d\n", sum)

	fmt.Println("\nReduce - producto de números pares:")
	product := lq.From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		Reduce(1, func(acc, n int) int {
			return acc * n
		})
	fmt.Printf("Producto de pares: %d\n", product)

	fmt.Println("\nReduce - encontrar el máximo:")
	max := lq.From(values).Reduce(0, func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	fmt.Printf("Máximo: %d\n", max)

	fmt.Println("\nReduce - contar elementos:")
	count := lq.From(values).Reduce(0, func(acc, n int) int {
		return acc + 1
	})
	fmt.Printf("Total de elementos: %d\n", count)

	fmt.Println("\nReduceTo - concatenar números como string:")
	concatenated := lq.ReduceTo(values, "", func(acc string, n int) string {
		if acc == "" {
			return fmt.Sprintf("%d", n)
		}
		return acc + "-" + fmt.Sprintf("%d", n)
	})
	fmt.Printf("Concatenado: %s\n", concatenated)

	// También puedes usar las funciones directamente
	fmt.Println("\nUsando Filter directamente:")
	evenNumbers := lq.Filter(values, func(n int) bool {
		return n%2 == 0
	})
	for value := range evenNumbers {
		fmt.Println(value)
	}
}
