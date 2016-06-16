package pkg

import "reflect"

func TypeName(val interface{}) string {
	typ := reflect.TypeOf(val)
	if typ.Kind() == reflect.Ptr {
		return typ.Elem().Name()
	}
	return typ.Name()
}
