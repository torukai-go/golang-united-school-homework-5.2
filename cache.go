package cache

import "time"

type Pair struct {
	Value    string
	Deadline time.Time
}

func NewPair(value string, deadline time.Time) Pair {
	return Pair{Value: value, Deadline: deadline}
}

type Cache struct {
	data map[string]Pair
}

func NewCache() Cache {
	return Cache{data: make(map[string]Pair)}
}

func (C Cache) Get(key string) (string, bool) {

	a, b := C.data[key]

	return a.Value, b
}

func (C Cache) Put(key, value string) {
	C.data[key] = NewPair(value, time.Time{})
}

func (C Cache) Keys() []string {
	mymap := C.data
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}

	return keys
}

func (C Cache) PutTill(key, value string, deadline time.Time) {
}
