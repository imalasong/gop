package main

import (
	"fmt"
	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/printer"
	"github.com/goplus/gop/token"
	"os"
)

func main() {

	//ast1()
	ast2()

	//goto1()
}

func ast2() {
	// 创建一个token集合，用于词法分析和语法分析
	fset := token.NewFileSet()

	// 解析源代码并返回AST
	src := `package a

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`
	node, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 修改AST：例如，添加一个新的函数声明
	node.Decls = append(node.Decls, &ast.FuncDecl{
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Slash: token.NoPos,
					Text:  "//11111",
				},
			},
		},
		Name: ast.NewIdent("foo"),
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("fmt"),
							Sel: ast.NewIdent("Println"),
						},
						Args: []ast.Expr{&ast.BasicLit{
							Kind:  token.STRING,
							Value: `"Hello from foo!"`,
						}},
					},
				},
			},
		},
	})

	// 打开一个新文件用于写入生成的源码
	outFile, err := os.Create("./cmd/a/output.go")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer outFile.Close()

	// 使用go/printer包将AST输出为Go源代码并写入文件
	err = printer.Fprint(outFile, fset, node)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("生成的源代码文件已写入output.go")
}
