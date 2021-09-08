package gen

import (
	"reflect"
	"strings"
)

func generateImports(objs ...interface{}) map[string]struct{} {
	rst := map[string]struct{}{}
	for _, obj := range objs {
		typ := reflect.TypeOf(obj)
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if sameModule(field.Type, typ) {
				continue
			}
			if private(field) {
				continue
			}
			if builtin(field.Type) {
				continue
			}
			if needsImport(field.Type) {
				rst[canonicalPath(field.Type)] = struct{}{}
			}
		}
	}
	return rst
}

func canonicalPath(typ reflect.Type) string {
	if typ.Kind() == reflect.Ptr {
		return typ.Elem().PkgPath()
	}
	return typ.PkgPath()
}

func private(field reflect.StructField) bool {
	return strings.ToUpper(field.Name[:1]) != field.Name[:1]
}

func sameModule(a, b reflect.Type) bool {
	if b.Kind() == reflect.Ptr {
		b = b.Elem()
	}
	if a.Kind() == reflect.Ptr {
		a = a.Elem()
	}
	return a.PkgPath() == b.PkgPath()
}

func builtin(typ reflect.Type) bool {
	if typ.Kind() == reflect.Ptr {
		return typ.Elem().PkgPath() == ""
	}
	return typ.PkgPath() == ""
}

func needsImport(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Ptr:
		return true
	case reflect.Slice:
		return true
	}
	return false
}
