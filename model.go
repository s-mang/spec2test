package spec2test

import (
	"fmt"
	"reflect"
	"strings"
)

type Resource struct {
	key         string
	typ         reflect.Type
	actionsOne  []Action
	actionsMany []Action
}

type actionInfo struct {
	httpMethod HTTPMethod
	hasBody    bool
}

var (
	resources = map[string]Resource{}

	actionHTTPMethods = map[Action]actionInfo{
		Create:  actionInfo{POST, true},
		Read:    actionInfo{GET, false},
		Update:  actionInfo{PUT, true},
		Destroy: actionInfo{DELETE, false},
	}

	keyFor = defaultKeyForFn
)

func AddResource(r interface{}, actionsOne, actionsMany []Action) error {
	rkey := keyFor(r)

	if _, ok := resources[rkey]; ok {
		return ErrorResourceAlreadyAdded
	}

	typ := dereference(reflect.ValueOf(r)).Type()

	if typ.Kind() != reflect.Struct {
		fmt.Println()
		return ErrorStructTypeRequired
	}

	resource := Resource{
		key:         rkey,
		typ:         typ,
		actionsOne:  actionsOne,
		actionsMany: actionsMany,
	}

	resources[rkey] = resource

	return nil
}

func AllResourceTypes() []reflect.Type {
	typs := make([]reflect.Type, 0, len(resources))
	for k := range resources {
		typs = append(typs, resources[k].typ)
	}

	return typs
}

func defaultKeyForFn(resource interface{}) string {
	typ := dereference(resource).Type()
	return strings.ToLower(typ.Name())
}
