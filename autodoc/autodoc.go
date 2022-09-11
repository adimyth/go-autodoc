package autodoc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"regexp"
	"strings"
)

func format(docstring string) string {
	// replace opening and closing comments section
	docstring = strings.ReplaceAll(docstring, "/*", "")
	docstring = strings.ReplaceAll(docstring, "*/", "")

	// highlight different blocks
	docstring = strings.ReplaceAll(docstring, "Parameters", "**Parameters**")
	docstring = strings.ReplaceAll(docstring, "Returns", "**Returns**")
	docstring = strings.ReplaceAll(docstring, "Raises", "**Raises**")
	docstring = strings.ReplaceAll(docstring, "Examples", "**Examples**")

	// get rid of "----" as some markdown engine process it as headers
	sampleRegexp := regexp.MustCompile(`\--*\-`)
	docstring = sampleRegexp.ReplaceAllString(docstring, "")

	return docstring
}

func DocGenerator(filename string) {
	var final string

	// parsing the source code into an AST tree
	parsed, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// traversing the tree
	ast.Inspect(parsed, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.FuncDecl:
			if n.Doc != nil {
				for _, comment := range n.Doc.List {
					final += fmt.Sprintf("## %s", n.Name.Name)
					final += format(comment.Text)
				}
			}
		}
		return true
	})

	// create target markdown file
	ext := path.Ext(filename)
	filename = strings.Replace(filename, "src", "docs", 1)
	targetFile := filename[0:len(filename)-len(ext)] + ".md"
	f, e := os.Create(targetFile)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	// write to the file
	res, err := f.WriteString(final)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", res)
}
