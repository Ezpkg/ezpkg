//go:build !genz

// Code generated by genz codez-matchers. DO NOT EDIT.

package codez

import (
	ast "go/ast"
)

// ArrayType
type ArrayTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// AssignStmt
type AssignStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// BasicLit
type BasicLitMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// BinaryExpr
type BinaryExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// BlockStmt
type BlockStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// BranchStmt
type BranchStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// CallExpr
type CallExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// CaseClause
type CaseClauseMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// ChanType
type ChanTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// CommClause
type CommClauseMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// CompositeLit
type CompositeLitMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// DeclStmt
type DeclStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// DeferStmt
type DeferStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// Ellipsis
type EllipsisMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// EmptyStmt
type EmptyStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// ExprStmt
type ExprStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// ForStmt
type ForStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// FuncDecl
type FuncDeclMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchDecl(decl ast.Decl) (bool, error)
}

// FuncLit
type FuncLitMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// FuncType
type FuncTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// GenDecl
type GenDeclMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchDecl(decl ast.Decl) (bool, error)
}

// GoStmt
type GoStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// Ident
type IdentMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// IfStmt
type IfStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// IncDecStmt
type IncDecStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// IndexExpr
type IndexExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// IndexListExpr
type IndexListExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// InterfaceType
type InterfaceTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// KeyValueExpr
type KeyValueExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// LabeledStmt
type LabeledStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// MapType
type MapTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// ParenExpr
type ParenExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// RangeStmt
type RangeStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// ReturnStmt
type ReturnStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// SelectStmt
type SelectStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// SelectorExpr
type SelectorExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// SendStmt
type SendStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// SliceExpr
type SliceExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// StarExpr
type StarExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// StructType
type StructTypeMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// SwitchStmt
type SwitchStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// TypeAssertExpr
type TypeAssertExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}

// TypeSwitchStmt
type TypeSwitchStmtMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchStmt(stmt ast.Stmt) (bool, error)
}

// UnaryExpr
type UnaryExprMatcher interface {
	Match(node ast.Node) (bool, error)
	MatchExpr(expr ast.Expr) (bool, error)
}