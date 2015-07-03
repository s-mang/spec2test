package spec2test_test

import (
	"fmt"

	"github.com/adams-sarah/spec2test"
)

func ExampleAddResource() {
	fmt.Println("# before:", len(spec2test.AllResourceTypes()))

	type MyType struct {
		Name string
	}

	oneActions := []spec2test.Action{
		spec2test.Create,
		spec2test.Read,
		spec2test.Update,
		spec2test.Destroy,
	}

	manyActions := []spec2test.Action{
		spec2test.Read,
	}

	err := spec2test.AddResource(MyType{"Sarah"}, oneActions, manyActions)
	if err != nil {
		// do something
	}

	fmt.Println("# after:", len(spec2test.AllResourceTypes()))

	// Output: # before: 0
	// # after: 1
}
