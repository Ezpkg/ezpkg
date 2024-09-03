//go:build !genz

// Code generated by genz codez-matchers. DO NOT EDIT.

package codez

import (
	ast "go/ast"
	token "go/token"
)

// FuncDecl
type zFuncDeclMatcher struct {
	_ *ast.FuncDecl

	Doc  CommentGroupMatcher
	Recv FieldListMatcher
	Name IdentMatcher
	Type FuncTypeMatcher
	Body BlockStmtMatcher
}

func (m zFuncDeclMatcher) MatchDecl(cx *_MatchContext, node ast.Decl) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zFuncDeclMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.FuncDecl)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Doc, x.Doc)
	ok, err = match(cx, ok, err, m.Recv, x.Recv)
	ok, err = match(cx, ok, err, m.Name, x.Name)
	ok, err = match(cx, ok, err, m.Type, x.Type)
	ok, err = match(cx, ok, err, m.Body, x.Body)
	return ok, err
}

// GenDecl
type zGenDeclMatcher struct {
	_ *ast.GenDecl

	Doc    CommentGroupMatcher
	TokPos token.Pos
	Tok    token.Token
	Lparen token.Pos
	Specs  SpecListMatcher[ast.Spec]
	Rparen token.Pos
}

func (m zGenDeclMatcher) MatchDecl(cx *_MatchContext, node ast.Decl) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zGenDeclMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.GenDecl)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Doc, x.Doc)
	ok, err = matchList(cx, ok, err, m.Specs, x.Specs)
	return ok, err
}
