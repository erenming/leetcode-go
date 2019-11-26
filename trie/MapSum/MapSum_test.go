package MapSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSum(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple", 3)
	//assert.Equal(t, 3, obj.Sum("ap"))
	obj.Insert("app", 2)
	assert.Equal(t, 5, obj.Sum("ap"))
}
