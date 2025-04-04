# ita

Package ita provides adaptors for iter.Seq and iter.Seq2.

[![Go Report Card](https://goreportcard.com/badge/github.com/shu-go/ita)](https://goreportcard.com/report/github.com/shu-go/ita)
[![Go Reference](https://pkg.go.dev/badge/github.com/shu-go/ita.svg)](https://pkg.go.dev/github.com/shu-go/ita)
![MIT License](https://img.shields.io/badge/License-MIT-blue)

# Install

```
go get github.com/shu-go/ita
```

# Examples

## NumAt adaptor

```go
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
```

## KeysValues adaptor

```go
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
```

## Chain

```go
ita.Chain[string](slices.Values([]string{"rei", "ichi", "ni", "san", "shi", "go"})).
	Pipe(strings.ToUpper).
	Filter(func(s string) bool { return len(s) <= 2 }).
	Exec(func(s string) { fmt.Println(s) })

// Output:
// NI
// GO
```

<!--  vim: set et ft=markdown sts=4 sw=4 ts=4 tw=0 :  -->
