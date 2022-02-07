// This is free and unencumbered software released into the public
// domain.  For more information, see <http://unlicense.org> or the
// accompanying UNLICENSE file.

package mocks_test

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"github.com/a8m/expect"
)

const packagePrefix = "package foo\n\n"

func source(expect func(interface{}) *expect.Expect, pkg string, decls []ast.Decl, scope *ast.Scope) string {
	buf := bytes.Buffer{}
	f := &ast.File{
		Name:  &ast.Ident{Name: pkg},
		Decls: decls,
		Scope: scope,
	}
	err := format.Node(&buf, token.NewFileSet(), f)
	expect(err).To.Be.Nil()
	return buf.String()
}

func parse(expect func(interface{}) *expect.Expect, code string) *ast.File {
	f, err := parser.ParseFile(token.NewFileSet(), "", packagePrefix+code, 0)
	expect(err).To.Be.Nil()
	expect(f).Not.To.Be.Nil()
	return f
}

func typeSpec(expect func(interface{}) *expect.Expect, spec string) *ast.TypeSpec {
	f := parse(expect, spec)
	expect(f.Scope.Objects).To.Have.Len(1)
	for _, obj := range f.Scope.Objects {
		spec, ok := obj.Decl.(*ast.TypeSpec)
		expect(ok).To.Be.Ok()
		return spec
	}
	return nil
}

func method(expect func(interface{}) *expect.Expect, spec *ast.TypeSpec) *ast.FuncType {
	inter, ok := spec.Type.(*ast.InterfaceType)
	expect(ok).To.Be.Ok()
	expect(inter.Methods.List).To.Have.Len(1)
	f, ok := inter.Methods.List[0].Type.(*ast.FuncType)
	expect(ok).To.Be.Ok()
	return f
}
