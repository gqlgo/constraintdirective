package main

import (
	"flag"
	"github.com/gqlgo/constraintdirective"
	"github.com/gqlgo/gqlanalysis/multichecker"
	"strings"
)

func main() {
	var types, excludes string
	flag.StringVar(&types, "types", "Int,Float,String", "target types. it can specify multiple values separated by `,`")
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL field name. it can specify multiple values separated by `,`")
	flag.Parse()
	multichecker.Main(
		constraintdirective.Analyzer(strings.Split(types, ","), strings.Split(excludes, ",")),
	)
}
