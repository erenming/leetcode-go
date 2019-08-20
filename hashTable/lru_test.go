package hashTable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	l := Constructor(3)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	assert.Equal(t, -1, l.Get(4))
	assert.Equal(t, 2, l.Get(2))
	assert.Equal(t, 2, l.Head.Next.Value)
	assert.Equal(t, 1, l.Tail.Prev.Value)

	l.Put(4, 4)
	assert.Equal(t, -1, l.Get(1))
	assert.Equal(t, 4, l.Head.Next.Value)
	assert.Equal(t, 3, l.Tail.Prev.Value)


}
