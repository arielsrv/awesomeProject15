package lq

import (
	"iter"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Test direct function
	evenNumbers := Filter(values, func(n int) bool {
		return n%2 == 0
	})

	var result []int
	for value := range evenNumbers {
		result = append(result, value)
	}

	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(result, expected) {
		t.Errorf("Filter failed: expected %v, got %v", expected, result)
	}
}

func TestMap(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Test direct function
	doubled := Map(values, func(n int) int {
		return n * 2
	})

	var result []int
	for value := range doubled {
		result = append(result, value)
	}

	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(result, expected) {
		t.Errorf("Map failed: expected %v, got %v", expected, result)
	}
}

func TestFlatMap(t *testing.T) {
	values := slices.Values([]int{1, 2, 3})

	// Test direct function
	expanded := FlatMap(values, func(n int) iter.Seq[int] {
		return slices.Values([]int{n, n * 2, n * 3})
	})

	var result []int
	for value := range expanded {
		result = append(result, value)
	}

	expected := []int{1, 2, 3, 2, 4, 6, 3, 6, 9}
	if !slices.Equal(result, expected) {
		t.Errorf("FlatMap failed: expected %v, got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Test direct function
	sum := Reduce(values, 0, func(acc, n int) int {
		return acc + n
	})

	expected := 15
	if sum != expected {
		t.Errorf("Reduce failed: expected %d, got %d", expected, sum)
	}
}

func TestReduceTo(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Test direct function
	concatenated := ReduceTo(values, "", func(acc string, n int) string {
		if acc == "" {
			return "1"
		}
		return acc + "-" + "1"
	})

	expected := "1-1-1-1-1"
	if concatenated != expected {
		t.Errorf("ReduceTo failed: expected %s, got %s", expected, concatenated)
	}
}

func TestEnumerableFilter(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Test fluent API
	result := From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		ToSlice()

	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(result, expected) {
		t.Errorf("Enumerable.Filter failed: expected %v, got %v", expected, result)
	}
}

func TestEnumerableMap(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Test fluent API
	result := From(values).
		Map(func(n int) int { return n * 2 }).
		ToSlice()

	expected := []int{2, 4, 6, 8, 10}
	if !slices.Equal(result, expected) {
		t.Errorf("Enumerable.Map failed: expected %v, got %v", expected, result)
	}
}

func TestEnumerableFlatMap(t *testing.T) {
	values := slices.Values([]int{1, 2, 3})

	// Test fluent API
	result := From(values).
		FlatMap(func(n int) iter.Seq[int] {
			return slices.Values([]int{n, n + 1, n + 2})
		}).
		ToSlice()

	expected := []int{1, 2, 3, 2, 3, 4, 3, 4, 5}
	if !slices.Equal(result, expected) {
		t.Errorf("Enumerable.FlatMap failed: expected %v, got %v", expected, result)
	}
}

func TestEnumerableReduce(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Test fluent API
	sum := From(values).
		Reduce(0, func(acc, n int) int {
			return acc + n
		})

	expected := 15
	if sum != expected {
		t.Errorf("Enumerable.Reduce failed: expected %d, got %d", expected, sum)
	}
}

func TestEnumerableChaining(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Test chaining multiple operations
	result := From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * 2 }).
		ToSlice()

	expected := []int{4, 8, 12, 16, 20}
	if !slices.Equal(result, expected) {
		t.Errorf("Enumerable chaining failed: expected %v, got %v", expected, result)
	}
}

func TestEnumerableForEach(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	var result []int
	From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		ForEach(func(n int) {
			result = append(result, n)
		})

	expected := []int{2, 4}
	if !slices.Equal(result, expected) {
		t.Errorf("Enumerable.ForEach failed: expected %v, got %v", expected, result)
	}
}

func TestEmptySequence(t *testing.T) {
	values := slices.Values([]int{})

	// Test with empty sequence
	result := From(values).
		Filter(func(n int) bool { return n%2 == 0 }).
		ToSlice()

	if len(result) != 0 {
		t.Errorf("Empty sequence test failed: expected empty slice, got %v", result)
	}
}

func TestFilterEarlyReturn(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Create a filter that stops after the first match
	filtered := Filter(values, func(n int) bool {
		return n%2 == 0
	})

	// Use a yield function that returns false after the first item
	count := 0
	var result []int
	for value := range filtered {
		result = append(result, value)
		count++
		if count == 1 {
			// This should trigger the early return in Filter
			break
		}
	}

	if count != 1 {
		t.Errorf("Filter early return test failed: expected 1 item, got %d", count)
	}

	if len(result) != 1 || result[0] != 2 {
		t.Errorf("Filter early return test failed: expected [2], got %v", result)
	}
}

func TestMapEarlyReturn(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	// Create a map that doubles each value
	mapped := Map(values, func(n int) int {
		return n * 2
	})

	// Use a yield function that returns false after the first item
	count := 0
	var result []int
	for value := range mapped {
		result = append(result, value)
		count++
		if count == 1 {
			// This should trigger the early return in Map
			break
		}
	}

	if count != 1 {
		t.Errorf("Map early return test failed: expected 1 item, got %d", count)
	}

	if len(result) != 1 || result[0] != 2 {
		t.Errorf("Map early return test failed: expected [2], got %v", result)
	}
}

func TestFlatMapEarlyReturn(t *testing.T) {
	values := slices.Values([]int{1, 2, 3})

	// Create a flatmap that expands each value
	flattened := FlatMap(values, func(n int) iter.Seq[int] {
		return slices.Values([]int{n, n * 2, n * 3})
	})

	// Use a yield function that returns false after the first item
	count := 0
	var result []int
	for value := range flattened {
		result = append(result, value)
		count++
		if count == 1 {
			// This should trigger the early return in FlatMap
			break
		}
	}

	if count != 1 {
		t.Errorf("FlatMap early return test failed: expected 1 item, got %d", count)
	}

	if len(result) != 1 || result[0] != 1 {
		t.Errorf("FlatMap early return test failed: expected [1], got %v", result)
	}
}

func TestReduceWithEmptySequence(t *testing.T) {
	values := slices.Values([]int{})

	// Test reduce with empty sequence should return initial value
	result := From(values).Reduce(42, func(acc, n int) int {
		return acc + n
	})

	expected := 42
	if result != expected {
		t.Errorf("Reduce with empty sequence failed: expected %d, got %d", expected, result)
	}
}

func BenchmarkFilter(b *testing.B) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(values, func(n int) bool { return n%2 == 0 })
	}
}

func BenchmarkMap(b *testing.B) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(values, func(n int) int { return n * 2 })
	}
}

func BenchmarkReduce(b *testing.B) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(values, 0, func(acc, n int) int { return acc + n })
	}
}

func BenchmarkEnumerableChaining(b *testing.B) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		From(values).
			Filter(func(n int) bool { return n%2 == 0 }).
			Map(func(n int) int { return n * 2 }).
			ToSlice()
	}
}
