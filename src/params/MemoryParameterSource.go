package params

import (
	"sort"

	"github.com/x1n13y84issmd42/oasis/src/contract"
)

// MemoryParameterSource is a parameter source which uses a native map as a source storage.
type MemoryParameterSource struct {
	Data map[string]string
}

// NewMemoryParameterSource creates a new MemoryParameterSource instance.
func NewMemoryParameterSource() *MemoryParameterSource {
	return &MemoryParameterSource{
		Data: map[string]string{},
	}
}

// Add stores a kv pair in the source.
func (ds *MemoryParameterSource) Add(n string, v string) {
	ds.Data[n] = v
}

// Get returns a parameter by it's name.
func (ds *MemoryParameterSource) Get(pn string) string {
	return ds.Data[pn]
}

// Iterate returns an iterable channel to receive parameter tuples.
func (ds *MemoryParameterSource) Iterate() contract.ParameterIterator {
	ch := make(contract.ParameterIterator)

	go func() {
		keys := []string{}
		for pn := range ds.Data {
			keys = append(keys, pn)
		}

		sort.Strings(keys)

		for _, pn := range keys {
			ch <- contract.ParameterTuple{N: pn, V: ds.Data[pn]}
		}

		close(ch)
	}()

	return ch
}
