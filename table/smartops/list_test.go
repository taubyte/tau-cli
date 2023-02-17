package smartopsTable_test

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	smartopsTable "github.com/taubyte/tau/table/smartops"
)

func ExampleList() {
	smartops := []*structureSpec.SmartOp{
		{
			Id:   "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
			Name: "someSmartOp1",
			Call: "ping",
		},
		{
			Id:   "QmbUIDhRosp5BaXDASEWSCtpkQCgQCPdRVhnxjiSHfXdC0",
			Name: "someSmartOp2",
			Call: "testlib.ping",
		},
	}

	smartopsTable.List(smartops)

	// Output:
	// ┌─────────────────┬──────────────┬──────────────┐
	// │ ID              │ NAME         │ CALL         │
	// ├─────────────────┼──────────────┼──────────────┤
	// │ QmbAA8...HfXdWH │ someSmartOp1 │ ping         │
	// ├─────────────────┼──────────────┼──────────────┤
	// │ QmbUID...HfXdC0 │ someSmartOp2 │ testlib.ping │
	// └─────────────────┴──────────────┴──────────────┘
}
