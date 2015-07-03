package seed

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

type seeder struct {
	typ reflect.Type
	n   int
}

func newSeeder(i interface{}, n int) *seeder {
	return &seeder{reflect.TypeOf(i), n}
}

func makeInstances(typ reflect.Type, n int) ([]interface{}, error) {
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("typ must be struct")
	}

	instances := make([]interface{}, n)
	for i := 0; i < n; i++ {
		instances[i] = makeInstance(typ, i)
	}

	return instances, nil
}

func makeInstance(typ reflect.Type, i int) interface{} {
	newVal := reflect.New(typ)

	for j := 0; j < newVal.Elem().NumField(); j++ {
		ftyp := newVal.Elem().Field(j).Type()
		fieldVal := randomValue(ftyp, i)
		newVal.Elem().Field(j).Set(fieldVal)
	}

	return newVal.Interface()
}

func randomValue(typ reflect.Type, i int) reflect.Value {
	numWords := len(words)

	switch typ.Kind() {
	case reflect.String:
		n := rand.Intn(8)

		var parts []string
		for i := 0; i < n; i++ {
			i := rand.Intn(numWords)
			parts = append(parts, words[i])
		}

		val := strings.Title(strings.Join(parts, " "))
		val += fmt.Sprintf("-%d", i) // ensure uniqueness

		return reflect.ValueOf(val)

	default:
		fuzzer := fuzz.New().NilChance(0)
		val := reflect.New(typ)
		fuzzer.Fuzz(val.Interface())
		return val.Elem()
	}
}
