package ita_test

import (
	"bytes"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/shu-go/ita"
)

func Example_numAt() {
	na := myList{
		s: []string{"rei", "ichi", "ni", "san", "shi", "go"},
	}
	// myList has methods Num() and Item(int), and has no methods that return []string or iter.Seq[string].

	for e := range ita.NumAt(na.Num, na.Item) {
		fmt.Println(e)
	}

	// Output:
	// rei
	// ichi
	// ni
	// san
	// shi
	// go
}

func Example_keysValues() {
	kv := myMap{
		m: map[string]int{"rei": 0, "ichi": 1, "ni": 2, "san": 3, "shi": 4, "go": 5},
	}
	// myMap has methods KeysSorted() and ValueOf(string), and has no methods that return iter.Seq2[string, int].

	for k, v := range ita.KeysValues(kv.KeysSorted, kv.ValueOf) {
		fmt.Println(k, v)
	}

	// Output:
	// go 5
	// ichi 1
	// ni 2
	// rei 0
	// san 3
	// shi 4
}

func Example_pipe() {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go"}

	seq := slices.Values(s)
	seq = ita.Pipe(seq, strings.ToUpper)
	seq = ita.Pipe(seq, func(str string) string { return str[:2] })
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// RE
	// IC
	// NI
	// SA
	// SH
	// GO
}

func Example_filter() {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go"}

	seq := slices.Values(s)
	seq = ita.Filter(seq, func(str string) bool { return len(str) >= 3 })
	seq = ita.Filter(seq, func(str string) bool { return strings.HasPrefix(str, "s") })
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// san
	// shi
}

func Example_exec() {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go"}

	seq := slices.Values(s)
	seq = ita.Filter(seq, func(str string) bool { return len(str) >= 3 })
	seq = ita.Filter(seq, func(str string) bool { return strings.HasPrefix(str, "s") })
	ita.Exec(seq, func(str string) {
		fmt.Println(str)
	})

	// Output:
	// san
	// shi
}

func Example_takeFirst() {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go"}
	for first := range ita.TakeFirst(slices.All(s)) {
		fmt.Println(first)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
}

func Example_takeSecond() {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go"}
	for second := range ita.TakeSecond(slices.All(s)) {
		fmt.Println(second)
	}

	// Output:
	// rei
	// ichi
	// ni
	// san
	// shi
	// go
}

func Example_chain() {
	ita.Chain[string](slices.Values([]string{"rei", "ichi", "ni", "san", "shi", "go"})).
		Pipe(strings.ToUpper).
		Filter(func(s string) bool { return len(s) <= 2 }).
		Exec(func(s string) { fmt.Println(s) })

	// Output:
	// NI
	// GO
}

func BenchmarkChain(b *testing.B) {
	s := []string{"rei", "ichi", "ni", "san", "shi", "go", "roku", "nana", "hachi", "queue"}

	b.Run("manual", func(b *testing.B) {
		buf := bytes.Buffer{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf.Reset()
			for _, e := range s {
				buf.WriteString(strings.ToUpper(e)[:2])
			}
			//gotwant.Test(b, buf.String(), "REICNISASHGORONAHAQU")
		}
	})

	b.Run("ita funcs", func(b *testing.B) {
		buf := bytes.Buffer{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf.Reset()
			seq := slices.Values(s)
			seq = ita.Pipe(seq, strings.ToUpper)
			seq = ita.Pipe(seq, func(s string) string { return s[:2] })
			ita.Exec(seq, func(s string) { buf.WriteString(s) })
			//gotwant.Test(b, buf.String(), "REICNISASHGORONAHAQU")
		}
	})

	b.Run("ita.Chain", func(b *testing.B) {
		buf := bytes.Buffer{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf.Reset()
			ita.Chain[string](slices.Values(s)).
				Pipe(strings.ToUpper).
				Pipe(func(s string) string { return s[:2] }).
				Exec(func(s string) { buf.WriteString(s) })
			//gotwant.Test(b, buf.String(), "REICNISASHGORONAHAQU")
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type myList struct {
	s []string
}

func (na myList) Num() int {
	return len(na.s)
}

func (na myList) Item(i int) string {
	return na.s[i]
}

////////////////////////////////////////////////////////////////////////////////

type myMap struct {
	m map[string]int
}

func (m myMap) KeysSeq() iter.Seq[string] {
	return slices.Values(slices.Sorted(maps.Keys(m.m)))
}

func (m myMap) KeysSorted() []string {
	return slices.Sorted(maps.Keys(m.m))
}

func (m myMap) ValueOf(key string) int {
	return m.m[key]
}
