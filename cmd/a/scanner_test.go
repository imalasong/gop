package main

import (
	"fmt"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/scanner"
	"github.com/goplus/gop/token"
	"testing"
)

func TestScanner(t *testing.T) {

	filename := "test.go"
	var src []byte = []byte(`var b   string  aa \n (aa + aaa)`)

	fset := token.NewFileSet()
	f := fset.AddFile(filename, -1, len(src))

	var s scanner.Scanner
	pmode := parser.ParseComments | parser.SaveAbsFile

	var m scanner.Mode
	if pmode&parser.ParseComments != 0 {
		m = scanner.ScanComments
	}
	s.Init(f, src, nil, m)
	for pos, tok, lit := s.Scan(); tok != token.EOF; pos, tok, lit = s.Scan() {
		fmt.Printf("1pos = %v,tok = %v,lit=%v\n", pos, tok, lit)
	}

}
