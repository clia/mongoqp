# mongoqp
A MongoDB Query Filter expression parser for Go.

The expression, please see MongoDB's document: [https://docs.mongodb.com/manual/tutorial/query-documents](https://docs.mongodb.com/manual/tutorial/query-documents)

## An example expression is like:

```
{ status: "A", age: { $lt: 30 } }
```

## Usage

```Go
package main

import (
    "fmt"
    "github.com/clia/mongoqp"
)

func main() {
    parser := mongoqp.Parser{}

    exp, err := parser.Parse(`{ R_STAT: 10 }`)
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    } else {
        fmt.Printf("%#v\n", exp)
        fmt.Printf("%#v\n", exp.Properties[0])
        fmt.Printf("%#v\n", exp.Properties[0].Value)
    }

    exp2, err := parser.Parse(`{ ERR_S: { $gte: 1 } }`)
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    } else {
        fmt.Printf("%#v\n", exp2)
        fmt.Printf("%#v\n", exp2.Properties[0].Value)
    }
}
```