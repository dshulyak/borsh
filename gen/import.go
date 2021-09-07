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
			pkg := field.Type.PkgPath()
			parts := strings.Split(pkg, "/")
			short := parts[len(parts)-1]
			if strings.Contains(short, "-") {
				parts = strings.Split(short, "-")
				short = parts[len(parts)-1]
			}
			rst[pkg] = short
		}
	}
	return rst
}

func private(field reflect.StructField) bool {
	return strings.ToUpper(field.Name[:1]) != field.Name[:1]
}

func sameModule(a, b reflect.Type) bool {
	return a.PkgPath() == b.PkgPath()
}
