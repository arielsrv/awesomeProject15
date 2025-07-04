package main

import (
	"fmt"
	"iter"
	"slices"

	"awesomeProject15/lq"
)

// Enumerable wraps an iter.Seq to provide fluent operations
type Enumerable[T any] struct {
	seq iter.Seq[T]
}

// From creates a new Enumerable from an iter.Seq
func From[T any](seq iter.Seq[T]) Enumerable[T] {
	return Enumerable[T]{seq: seq}
}

// Filter applies a predicate function and returns a new Enumerable
func (r Enumerable[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return From(Filter(r.seq, predicate))
}

// Map applies a transformation function and returns a new Enumerable
func (r Enumerable[T]) Map(transform func(T) T) Enumerable[T] {
	return From(Map(r.seq, transform))
}

// FlatMap applies a transformation function that returns a sequence and flattens the result
func (r Enumerable[T]) FlatMap(transform func(T) iter.Seq[T]) Enumerable[T] {
	return From(FlatMap(r.seq, transform))
}

// Reduce accumulates values using a reducer function
func (r Enumerable[T]) Reduce(initial T, reducer func(T, T) T) T {
	return Reduce(r.seq, initial, reducer)
}

// ForEach iterates over the sequence and applies a function to each element
func (r Enumerable[T]) ForEach(action func(T)) {
	for value := range r.seq {
		action(value)
	}
}

// ToSlice converts the sequence to a slice
func (r Enumerable[T]) ToSlice() []T {
	var result []T
	for value := range r.seq {
		result = append(result, value)
	}
	return result
}

// Filter takes an iter.Seq and a predicate function, returning a filtered iter.Seq
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Map takes an iter.Seq and a transform function, returning a transformed iter.Seq
func Map[T, U any](sequence iter.Seq[T], transform func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for value := range sequence {
			if !yield(transform(value)) {
				return
			}
		}
	}
}

// FlatMap takes an iter.Seq and a transform function that returns a sequence, flattening the result
func FlatMap[T any](sequence iter.Seq[T], transform func(T) iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range sequence {
			for transformedValue := range transform(value) {
				if !yield(transformedValue) {
					return
				}
			}
		}
	}
}

// Reduce takes an iter.Seq, initial value, and reducer function, returning the accumulated result
func Reduce[T any](sequence iter.Seq[T], initial T, reducer func(T, T) T) T {
	result := initial
	for value := range sequence {
		result = reducer(result, value)
	}
	return result
}

// ReduceTo takes an iter.Seq, initial value of different type, and reducer function
func ReduceTo[T, U any](sequence iter.Seq[T], initial U, reducer func(U, T) U) U {
	result := initial
	for value := range sequence {
		result = reducer(result, value)
	}
	return result
}

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
