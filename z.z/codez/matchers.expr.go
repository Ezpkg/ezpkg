//go:build !genz

// Code generated by genz codez-matchers. DO NOT EDIT.

package codez

import (
	ast "go/ast"
	token "go/token"
)

// ArrayType
type zArrayTypeMatcher struct {
	_ *ast.ArrayType

	Lbrack token.Pos
	Len    ExprMatcher
	Elt    ExprMatcher
}

func (m zArrayTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zArrayTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.ArrayType)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Len, x.Len)
	ok, err = match(cx, ok, err, m.Elt, x.Elt)
	return ok, err
}

// BasicLit
type zBasicLitMatcher struct {
	_ *ast.BasicLit

	ValuePos token.Pos
	Kind     token.Token
	Value    StringMatcher
}

func (m zBasicLitMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zBasicLitMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.BasicLit)
	if !ok {
		return false, nil
	}
	ok, err = matchValue(cx, ok, err, m.Value, x.Value)
	return ok, err
}

// BinaryExpr
type zBinaryExprMatcher struct {
	_ *ast.BinaryExpr

	X     ExprMatcher
	OpPos token.Pos
	Op    token.Token
	Y     ExprMatcher
}

func (m zBinaryExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zBinaryExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.BinaryExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = match(cx, ok, err, m.Y, x.Y)
	return ok, err
}

// CallExpr
type zCallExprMatcher struct {
	_ *ast.CallExpr

	Fun      ExprMatcher
	Lparen   token.Pos
	Args     ExprListMatcher[ast.Expr]
	Ellipsis token.Pos
	Rparen   token.Pos
}

func (m zCallExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zCallExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.CallExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Fun, x.Fun)
	ok, err = matchList(cx, ok, err, m.Args, x.Args)
	return ok, err
}

// ChanType
type zChanTypeMatcher struct {
	_ *ast.ChanType

	Begin token.Pos
	Arrow token.Pos
	Dir   ChanDirMatcher
	Value ExprMatcher
}

func (m zChanTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zChanTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.ChanType)
	if !ok {
		return false, nil
	}
	ok, err = matchValue(cx, ok, err, m.Dir, x.Dir)
	ok, err = match(cx, ok, err, m.Value, x.Value)
	return ok, err
}

// CompositeLit
type zCompositeLitMatcher struct {
	_ *ast.CompositeLit

	Type       ExprMatcher
	Lbrace     token.Pos
	Elts       ExprListMatcher[ast.Expr]
	Rbrace     token.Pos
	Incomplete BoolMatcher
}

func (m zCompositeLitMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zCompositeLitMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.CompositeLit)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Type, x.Type)
	ok, err = matchList(cx, ok, err, m.Elts, x.Elts)
	ok, err = matchValue(cx, ok, err, m.Incomplete, x.Incomplete)
	return ok, err
}

// Ellipsis
type zEllipsisMatcher struct {
	_ *ast.Ellipsis

	Ellipsis token.Pos
	Elt      ExprMatcher
}

func (m zEllipsisMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zEllipsisMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.Ellipsis)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Elt, x.Elt)
	return ok, err
}

// FuncLit
type zFuncLitMatcher struct {
	_ *ast.FuncLit

	Type FuncTypeMatcher
	Body BlockStmtMatcher
}

func (m zFuncLitMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zFuncLitMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.FuncLit)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Type, x.Type)
	ok, err = match(cx, ok, err, m.Body, x.Body)
	return ok, err
}

// FuncType
type zFuncTypeMatcher struct {
	_ *ast.FuncType

	Func       token.Pos
	TypeParams FieldListMatcher
	Params     FieldListMatcher
	Results    FieldListMatcher
}

func (m zFuncTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zFuncTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.FuncType)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.TypeParams, x.TypeParams)
	ok, err = match(cx, ok, err, m.Params, x.Params)
	ok, err = match(cx, ok, err, m.Results, x.Results)
	return ok, err
}

// Ident
type zIdentMatcher struct {
	_ *ast.Ident

	NamePos token.Pos
	Name    StringMatcher
}

func (m zIdentMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zIdentMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.Ident)
	if !ok {
		return false, nil
	}
	ok, err = matchValue(cx, ok, err, m.Name, x.Name)
	return ok, err
}

// IndexExpr
type zIndexExprMatcher struct {
	_ *ast.IndexExpr

	X      ExprMatcher
	Lbrack token.Pos
	Index  ExprMatcher
	Rbrack token.Pos
}

func (m zIndexExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zIndexExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.IndexExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = match(cx, ok, err, m.Index, x.Index)
	return ok, err
}

// IndexListExpr
type zIndexListExprMatcher struct {
	_ *ast.IndexListExpr

	X       ExprMatcher
	Lbrack  token.Pos
	Indices ExprListMatcher[ast.Expr]
	Rbrack  token.Pos
}

func (m zIndexListExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zIndexListExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.IndexListExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = matchList(cx, ok, err, m.Indices, x.Indices)
	return ok, err
}

// InterfaceType
type zInterfaceTypeMatcher struct {
	_ *ast.InterfaceType

	Interface  token.Pos
	Methods    FieldListMatcher
	Incomplete BoolMatcher
}

func (m zInterfaceTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zInterfaceTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.InterfaceType)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Methods, x.Methods)
	ok, err = matchValue(cx, ok, err, m.Incomplete, x.Incomplete)
	return ok, err
}

// KeyValueExpr
type zKeyValueExprMatcher struct {
	_ *ast.KeyValueExpr

	Key   ExprMatcher
	Colon token.Pos
	Value ExprMatcher
}

func (m zKeyValueExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zKeyValueExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.KeyValueExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Key, x.Key)
	ok, err = match(cx, ok, err, m.Value, x.Value)
	return ok, err
}

// MapType
type zMapTypeMatcher struct {
	_ *ast.MapType

	Map   token.Pos
	Key   ExprMatcher
	Value ExprMatcher
}

func (m zMapTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zMapTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.MapType)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Key, x.Key)
	ok, err = match(cx, ok, err, m.Value, x.Value)
	return ok, err
}

// ParenExpr
type zParenExprMatcher struct {
	_ *ast.ParenExpr

	Lparen token.Pos
	X      ExprMatcher
	Rparen token.Pos
}

func (m zParenExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zParenExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.ParenExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	return ok, err
}

// SelectorExpr
type zSelectorExprMatcher struct {
	_ *ast.SelectorExpr

	X   ExprMatcher
	Sel IdentMatcher
}

func (m zSelectorExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zSelectorExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.SelectorExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = match(cx, ok, err, m.Sel, x.Sel)
	return ok, err
}

// SliceExpr
type zSliceExprMatcher struct {
	_ *ast.SliceExpr

	X      ExprMatcher
	Lbrack token.Pos
	Low    ExprMatcher
	High   ExprMatcher
	Max    ExprMatcher
	Slice3 BoolMatcher
	Rbrack token.Pos
}

func (m zSliceExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zSliceExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.SliceExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = match(cx, ok, err, m.Low, x.Low)
	ok, err = match(cx, ok, err, m.High, x.High)
	ok, err = match(cx, ok, err, m.Max, x.Max)
	ok, err = matchValue(cx, ok, err, m.Slice3, x.Slice3)
	return ok, err
}

// StarExpr
type zStarExprMatcher struct {
	_ *ast.StarExpr

	Star token.Pos
	X    ExprMatcher
}

func (m zStarExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zStarExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.StarExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	return ok, err
}

// StructType
type zStructTypeMatcher struct {
	_ *ast.StructType

	Struct     token.Pos
	Fields     FieldListMatcher
	Incomplete BoolMatcher
}

func (m zStructTypeMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zStructTypeMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.StructType)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.Fields, x.Fields)
	ok, err = matchValue(cx, ok, err, m.Incomplete, x.Incomplete)
	return ok, err
}

// TypeAssertExpr
type zTypeAssertExprMatcher struct {
	_ *ast.TypeAssertExpr

	X      ExprMatcher
	Lparen token.Pos
	Type   ExprMatcher
	Rparen token.Pos
}

func (m zTypeAssertExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zTypeAssertExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.TypeAssertExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	ok, err = match(cx, ok, err, m.Type, x.Type)
	return ok, err
}

// UnaryExpr
type zUnaryExprMatcher struct {
	_ *ast.UnaryExpr

	OpPos token.Pos
	Op    token.Token
	X     ExprMatcher
}

func (m zUnaryExprMatcher) MatchExpr(cx *_MatchContext, node ast.Expr) (ok bool, err error) {
	return m.Match(cx, node)
}
func (m zUnaryExprMatcher) Match(cx *_MatchContext, node ast.Node) (ok bool, err error) {
	x, ok := node.(*ast.UnaryExpr)
	if !ok {
		return false, nil
	}
	ok, err = match(cx, ok, err, m.X, x.X)
	return ok, err
}
