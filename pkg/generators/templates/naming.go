package templates

import (
	"go/ast"
	"go/parser"
	"strings"
)

// TypeDisplayName returns the value with special characters replaced with
// there semantic meaning and then converted to pascal case. For example,
//
//	*string
//
// becomes
//
//	NullableString
func TypeDisplayName(value string) string {
	expr, error := parser.ParseExpr(value)

	if error != nil {
		panic(error)
	}

	var parts []string

	var f func(n ast.Node) bool

	f = func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.InterfaceType:
			parts = append(parts, "interface")
		case *ast.FieldList:
			// Not sure if I need to do something here
		case *ast.StarExpr:
			parts = append(parts, "nullable")
		case *ast.Ident:
			parts = append(parts, x.Name)
		case *ast.ArrayType:
			ast.Inspect(x.Elt, f)
			parts = append(parts, "slice")
			return false
		case *ast.StructType:
			parts = append(parts, "struct")
		case *ast.Field:
			for _, name := range x.Names {
				parts = append(parts, name.Name)
			}
			return false
		case nil:
			// Nothing to be done
		default:
			panic(x)
		}
		return true
	}

	ast.Inspect(expr, f)

	return ToPascalCase(strings.Join(parts, "_"))
}
