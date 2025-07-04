package lq

import (
	"iter"
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
