package gen

import (
	"reflect"
	"strings"
)

func generateImports(objs ...interface{}) map[string]string {
	rst := map[string]string{}
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
				rst[field.Type.PkgPath()] = canonicalName(field.Type.PkgPath())
			}
		}
	}
	return rst
}

func canonicalName(pkg string) string {
	parts := strings.Split(pkg, "/")
	short := parts[len(parts)-1]
	if strings.Contains(short, "-") {
		parts = strings.Split(short, "-")
		short = parts[len(parts)-1]
	}
	return short
}

func private(field reflect.StructField) bool {
	return strings.ToUpper(field.Name[:1]) != field.Name[:1]
}

func sameModule(a, b reflect.Type) bool {
	return a.PkgPath() == b.PkgPath()
}

func builtin(typ reflect.Type) bool {
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
