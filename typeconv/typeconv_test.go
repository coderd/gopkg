package typeconv

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	var src interface{} = map[string]int{"k": 1}
	var dst map[string]int

	// Test correct cases
	err := Set(&dst, src)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(dst, map[string]int{"k": 1}))

	// Test failed cases
	// type mismatch
	var dst2 map[string]bool
	err = Set(&dst2, src)
	assert.NotNil(t, err)

	// dst is not pointer
	err = Set(dst2, src)
	assert.NotNil(t, err)

	// dst is nil pointer
	var dst3 = &dst2
	dst3 = nil
	err = Set(dst3, src)
	assert.NotNil(t, err)
}

func TestMapStringBool(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": true,
		"k2": false,
	}
	_, err = MapStringBool(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringBool(v1)
	assert.NotNil(t, err)
}

func TestMapStringInt(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringInt(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2.0,
	}
	_, err = MapStringInt(v1)
	assert.NotNil(t, err)
}

func TestMapStringFloat64(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": 1.0,
		"k2": 2.1,
	}
	_, err = MapStringFloat64(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringFloat64(v1)
	assert.NotNil(t, err)
}

func TestMapStringString(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": "v1",
		"k2": "v2",
	}
	_, err = MapStringString(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringString(v1)
	assert.NotNil(t, err)
}

func TestSliceBool(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		true,
		false,
	}
	_, err = SliceBool(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = SliceBool(v1)
	assert.NotNil(t, err)
}

func TestSliceFloat64(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		1.0,
		2.1,
	}
	_, err = SliceFloat64(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = SliceFloat64(v1)
	assert.NotNil(t, err)
}

func TestSliceString(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		"v1",
		"v2",
	}
	_, err = SliceString(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = SliceString(v1)
	assert.NotNil(t, err)
}
