//go:build !genz

// Code generated by genz codez-matchers. DO NOT EDIT.

package codez

import (
	ast "go/ast"
	token "go/token"
)

// ImportSpec
type ImportSpecMatcherB struct {
	_ *ast.ImportSpec

	Doc     CommentGroupMatcher
	Name    IdentMatcher
	Path    BasicLitMatcher
	Comment CommentGroupMatcher
	EndPos  token.Pos
}

func (m ImportSpecMatcherB) MatchSpec(cx *MatchContext, node ast.Spec) (ok bool, err error) {
	return m.Match(cx, node)
}

func (m ImportSpecMatcherB) Match(cx *MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.ImportSpec)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Doc, x.Doc)
	ok, err = match(cx, ok, err, m.Name, x.Name)
	ok, err = match(cx, ok, err, m.Path, x.Path)
	ok, err = match(cx, ok, err, m.Comment, x.Comment)
	return ok, err
}

// TypeSpec
type TypeSpecMatcherB struct {
	_ *ast.TypeSpec

	Doc        CommentGroupMatcher
	Name       IdentMatcher
	TypeParams FieldListMatcher
	Assign     token.Pos
	Type       ExprMatcher
	Comment    CommentGroupMatcher
}

func (m TypeSpecMatcherB) MatchSpec(cx *MatchContext, node ast.Spec) (ok bool, err error) {
	return m.Match(cx, node)
}

func (m TypeSpecMatcherB) Match(cx *MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.TypeSpec)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Doc, x.Doc)
	ok, err = match(cx, ok, err, m.Name, x.Name)
	ok, err = match(cx, ok, err, m.TypeParams, x.TypeParams)
	ok, err = match(cx, ok, err, m.Type, x.Type)
	ok, err = match(cx, ok, err, m.Comment, x.Comment)
	return ok, err
}

// ValueSpec
type ValueSpecMatcherB struct {
	_ *ast.ValueSpec

	Doc     CommentGroupMatcher
	Names   ListMatcher[*ast.Ident]
	Type    ExprMatcher
	Values  ListMatcher[ast.Expr]
	Comment CommentGroupMatcher
}

func (m ValueSpecMatcherB) MatchSpec(cx *MatchContext, node ast.Spec) (ok bool, err error) {
	return m.Match(cx, node)
}

func (m ValueSpecMatcherB) Match(cx *MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.ValueSpec)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Doc, x.Doc)
	ok, err = matchList(cx, ok, err, m.Names, x.Names)
	ok, err = match(cx, ok, err, m.Type, x.Type)
	ok, err = matchList(cx, ok, err, m.Values, x.Values)
	ok, err = match(cx, ok, err, m.Comment, x.Comment)
	return ok, err
}
