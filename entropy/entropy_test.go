package entropy

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

var random = filepath.Join("..", "testdata", "random.bin")

func TestCalculate(t *testing.T) {
	t.Run("Test Calculate", func(t *testing.T) {
		blk, err := fixture(random)

		if err != nil {
			t.Fatalf("Calculate: %v", err)
		}

		if Calculate(blk) != 7.826049715293044 {
			t.Fatal("entropy wrong")
		}
	})
}

func BenchmarkCalculate(b *testing.B) {
	b.Run("Benchmark Calculate", func(b *testing.B) {
		blk, err := fixture(random)

		if err != nil {
			b.Fatalf("Calculate: %v", err)
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = Calculate(blk)
		}
	})
}

func fixture(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	b, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return b, nil
}
