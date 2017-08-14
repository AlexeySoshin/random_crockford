package random_crockford_test

import (
	"testing"
	"github.com/alexeysoshin/random_crockford"
)

func TestNextLength(t *testing.T) {

	t.Parallel()

	r := random_crockford.NewRandom(10)

	n := r.Next()

	if len(n) != 10 {
		t.Fatalf("%s is not 10 characters long", n)
	}
}

func TestNextZeroLength(t *testing.T) {
	r := random_crockford.NewRandom(0)

	n := r.Next()

	if n != "" {
		t.Fatalf("Should return empty string")
	}
}

func TestNextNegativeLength(t *testing.T) {
	r := random_crockford.NewRandom(-1)

	n := r.Next()

	if n != "" {
		t.Fatalf("Should return empty string")
	}
}

func TestNextShouldBeRandom(t *testing.T) {

	r := random_crockford.NewRandom(8)

	nexts := map[string]bool{}

	for i := 0; i < 100000; i++ {

		next := r.Next()
		if nexts[next] {
			t.Fatalf("%s repeated", next)
		}

		nexts[next] = true
	}
}

func BenchmarkNext(b *testing.B) {

	r := random_crockford.NewRandom(10)

	for i := 0; i < b.N; i++ {
		_ = r.Next()
	}
}