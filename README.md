# constraintdirective

[![pkg.go.dev][gopkg-badge]][gopkg]

`constraintdirective` finds fields and arguments without @constraint directive in your GraphQL schema files.

```graphql
input MutationInput {
    name: String! @constraint(minLength: 1)
    field: String! # want "field has no constraint directive"
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
package main

import (
	"flag"
	"github.com/gqlgo/constraintdirective"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	multichecker.Main(
		constraintdirective.Analyzer(),
	)
}
```

`constraintdirective` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/constraintdirective/cmd/constraintdirective@latest
```

The `constraintdirective` command has a flag, `schema` which will be parsed and analyzed by constraintdirective's Analyzer.

```sh
$ constraintdirective -schema="server/graphql/schema/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql".

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/constraintdirective
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/constraintdirective?status.svg
