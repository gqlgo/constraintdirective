package constraintdirective

import (
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
	"slices"
)

func Analyzer(types, excludeFieldNames []string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "constraintdirective",
		Doc:  "constraintdirective finds field and argument without @constraint directive",
		Run:  run(types, excludeFieldNames),
	}
}

func isTargetType(types []string, t *ast.Type) bool {
	if t == nil {
		return false
	}
	if slices.Contains(types, t.NamedType) {
		return true
	}
	return isTargetType(types, t.Elem)
}

func isExcludeFieldName(ignoreFieldNames []string, fieldName string) bool {
	return slices.Contains(ignoreFieldNames, fieldName)
}

func run(types, excludeFieldNames []string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			switch t.Kind {
			case ast.InputObject:
				for _, field := range t.Fields {
					if field == nil || field.Type == nil {
						continue
					}
					if !isTargetType(types, field.Type) || isExcludeFieldName(excludeFieldNames, field.Name) {
						continue
					}
					if field.Directives == nil || field.Directives.ForName("constraint") == nil {
						if field.Position != nil {
							pass.Reportf(field.Position, "%s has no constraint directive", field.Name)
						}
					}
				}
			case ast.Object:
				for _, field := range t.Fields {
					if field == nil {
						continue
					}
					for _, arg := range field.Arguments {
						if arg == nil || arg.Type == nil {
							continue
						}
						if !isTargetType(types, arg.Type) || isExcludeFieldName(excludeFieldNames, arg.Name) {
							continue
						}
						if arg.Directives == nil || arg.Directives.ForName("constraint") == nil {
							if field.Position != nil {
								pass.Reportf(field.Position, "argument %s of %s has no constraint directive", arg.Name, field.Name)
							}
						}
					}
				}
			}
		}

		return nil, nil
	}
}
