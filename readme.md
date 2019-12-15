# JavaScript Client for [Cayley](https://github.com/cayleygraph/cayley)

### [Documentation](https://cayleygraph.github.io/javascript-client/)

### Installation

```bash
go get -u github.com/cayleygraph/go-client
```

### Usage

#### Log 10 nodes from the graph

```go
package main

import (
    "context"
    "fmt"
    "github.com/cayleygraph/go-client"
)

func main() {
	c := client.NewAPIClient(client.NewConfiguration())
	result, _, _ := c.QueriesApi.Query(context.TODO(), "gizmo", "g.V().getLimit(10)")
	fmt.Printf("%v", result.Result)
}
```
