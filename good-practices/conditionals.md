
## Conditionals

```go
package main

import (
	"fmt"
	"time"
)

func con3(i int) bool {
	time.Sleep(1 * time.Second)
	return i > 1
}

func main() {
	con1 := true
	con2 := true

	// recommended flow
	if con1 && con2 {
		fmt.Println("valid 1")
	}
	// flow with unnecessary if
	if con1 {
		if con2 {
			fmt.Println("valid 2")
		}
	}

	// flow with if required
	if con1 {
		if con2 {
			fmt.Println("valid 3")
		}
		if con3(2) {
			fmt.Println("valid 4")
		}
	}

	// flow not recommended
	if con3(2) && con1 && con2 {
		fmt.Println("valid 5")
	}

	// flow recommended, first validate the simple logic codes and then the complex codes

	if con1 && con2 {
		if con3(2) {
			fmt.Println("valid 6")
		}
	}
}

```

**Code** https://go.dev/play/p/IhNwozykNOo