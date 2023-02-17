package serviceTable_test

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	serviceTable "github.com/taubyte/tau/table/service"
)

func ExampleList() {
	services := []*structureSpec.Service{
		{
			Id:       "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
			Name:     "someService1",
			Protocol: "/test/v1",
		},
		{
			Id:       "QmbUIDhRosp5BaXDASEWSCtpkQCgQCPdRVhnxjiSHfXdC0",
			Name:     "someService2",
			Protocol: "/test/v2",
		},
	}

	serviceTable.List(services)

	// Output:
	// ┌─────────────────┬──────────────┬──────────┐
	// │ ID              │ NAME         │ PROTOCOL │
	// ├─────────────────┼──────────────┼──────────┤
	// │ QmbAA8...HfXdWH │ someService1 │ /test/v1 │
	// ├─────────────────┼──────────────┼──────────┤
	// │ QmbUID...HfXdC0 │ someService2 │ /test/v2 │
	// └─────────────────┴──────────────┴──────────┘
}
