package maglev

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// TestLookup tests the lookup field in table
func TestLookup(t *testing.T) {
	table := getTestingMaglevTable()

	if !reflect.DeepEqual(table.lookup, []int{
		1, 2, 0, 2, 0, 1, 0,
	}) {
		t.Errorf("table lookup field not the same")
	}
}

func TestDistribution(t *testing.T) {
	const size = 30000

	var names []string
	for i := 0; i < size; i++ {
		names = append(names, fmt.Sprintf("backend-%d", i))
	}

	table := New(names, SmallM)

	r := make([]int, size)
	rand.Seed(0)
	for i := 0; i < 1e8; i++ {
		idx := table.Lookup(uint64(rand.Int63()))
		r[idx]++
	}

	var total int
	var max = 0
	for _, v := range r {
		total += v
		if v > max {
			max = v
		}
	}

	mean := float64(total) / size
	t.Logf("max=%v, mean=%v, peak-to-mean=%v", max, mean, float64(max)/mean)
}

func getTestingMaglevTable() *Table {
	return New([]string{
		"backend-0",
		"backend-1",
		"backend-2",
	}, 7)
}
