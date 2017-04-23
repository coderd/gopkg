package typeconv

import (
	"fmt"
	"reflect"
)

// An invalidPointerError describes an invalid argument passed to Set.
type invalidPointerError struct {
	Type reflect.Type
}

func (i *invalidPointerError) Error() string {
	if i.Type == nil {
		return "typeconv: Set (nil)"
	}

	if i.Type.Kind() != reflect.Ptr {
		return "typeconv: Set(non-pointer " + i.Type.String() + ")"
	}
	return "typeconv: Set(nil " + i.Type.String() + ")"
}

func Set(dst interface{}, src interface{}) error {
	dstReflectV := reflect.ValueOf(dst)
	dstReflectT := reflect.TypeOf(dst)
	if dstReflectV.Kind() != reflect.Ptr || dstReflectV.IsNil() {
		return &invalidPointerError{dstReflectT}
	}

	srcReflectV := reflect.ValueOf(src)
	srcReflectT := reflect.TypeOf(src)

	dstElem := dstReflectV.Elem()
	dstElemT := dstElem.Type()

	if dstElemT != srcReflectT {
		return fmt.Errorf("gconfig: set() dst's element type(%s) and src's type(%s) do not match", dstElemT.String(), srcReflectT.String())
	}

	dstElem.Set(srcReflectV)

	return nil
}

func MapStringBool(v map[string]interface{}) (map[string]bool, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make(map[string]bool, leng)
	for key, value := range v {
		if s, ok := value.(bool); ok {
			r[key] = s
		} else {
			return nil, fmt.Errorf("typeconv: MapStringBool() type `%s` can't convert to `%s`", reflect.TypeOf(value), "bool")
		}
	}

	return r, nil
}

func MapStringInt(v map[string]interface{}) (map[string]int, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make(map[string]int, leng)
	for key, value := range v {
		if s, ok := value.(int); ok {
			r[key] = s
		} else {
			return nil, fmt.Errorf("typeconv: MapStringInt() type `%s` can't convert to `%s`", reflect.TypeOf(value), "int")
		}
	}

	return r, nil
}

func MapStringFloat64(v map[string]interface{}) (map[string]float64, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make(map[string]float64, leng)
	for key, value := range v {
		if s, ok := value.(float64); ok {
			r[key] = s
		} else {
			return nil, fmt.Errorf("typeconv: MapStringFloat64() type `%s` can't convert to `%s`", reflect.TypeOf(value), "float64")
		}
	}

	return r, nil
}

func MapStringString(v map[string]interface{}) (map[string]string, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make(map[string]string, leng)
	for key, value := range v {
		if s, ok := value.(string); ok {
			r[key] = s
		} else {
			return nil, fmt.Errorf("typeconv: MapStringString() type `%s` can't convert to `%s`", reflect.TypeOf(value), "string")
		}
	}

	return r, nil
}

func SliceBool(v []interface{}) ([]bool, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make([]bool, 0, leng)
	for _, value := range v {
		if s, ok := value.(bool); ok {
			r = append(r, s)
		} else {
			return nil, fmt.Errorf("typeconv: SliceBool() type `%s` can't convert to `%s`", reflect.TypeOf(value), "bool")
		}
	}

	return r, nil
}

func SliceFloat64(v []interface{}) ([]float64, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make([]float64, 0, leng)
	for _, value := range v {
		if s, ok := value.(float64); ok {
			r = append(r, s)
		} else {
			return nil, fmt.Errorf("typeconv: SliceFloat64() type `%s` can't convert to `%s`", reflect.TypeOf(value), "float64")
		}
	}

	return r, nil
}

func SliceString(v []interface{}) ([]string, error) {
	if v == nil {
		return nil, nil
	}

	leng := len(v)
	r := make([]string, 0, leng)
	for _, value := range v {
		if s, ok := value.(string); ok {
			r = append(r, s)
		} else {
			return nil, fmt.Errorf("typeconv: SliceString() type `%s` can't convert to `%s`", reflect.TypeOf(value), "string")
		}
	}

	return r, nil
}
