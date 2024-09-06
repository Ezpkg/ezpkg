//go:build !genz

// Code generated by genz codez-matchers. DO NOT EDIT.

package codez

import (
	"go/ast"
)

// ArrayType
type ArrayTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// AssignStmt
type AssignStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// BasicLit
type BasicLitMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// BinaryExpr
type BinaryExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// BlockStmt
type BlockStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// BranchStmt
type BranchStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// CallExpr
type CallExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// CaseClause
type CaseClauseMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// ChanType
type ChanTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// CommClause
type CommClauseMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// Comment
type CommentMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// CommentGroup
type CommentGroupMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// CompositeLit
type CompositeLitMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// DeclStmt
type DeclStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// DeferStmt
type DeferStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// Ellipsis
type EllipsisMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// EmptyStmt
type EmptyStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// ExprStmt
type ExprStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// Field
type FieldMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// FieldList
type FieldListMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// File
type FileMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// ForStmt
type ForStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// FuncDecl
type FuncDeclMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchDecl(cx *MatchContext, decl ast.Decl) (bool, error)
}

// FuncLit
type FuncLitMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// FuncType
type FuncTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// GenDecl
type GenDeclMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchDecl(cx *MatchContext, decl ast.Decl) (bool, error)
}

// GoStmt
type GoStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// Ident
type IdentMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// IfStmt
type IfStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// ImportSpec
type ImportSpecMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// IncDecStmt
type IncDecStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// IndexExpr
type IndexExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// IndexListExpr
type IndexListExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// InterfaceType
type InterfaceTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// KeyValueExpr
type KeyValueExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// LabeledStmt
type LabeledStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// MapType
type MapTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// ParenExpr
type ParenExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// RangeStmt
type RangeStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// ReturnStmt
type ReturnStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// SelectStmt
type SelectStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// SelectorExpr
type SelectorExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// SendStmt
type SendStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// SliceExpr
type SliceExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// StarExpr
type StarExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// StructType
type StructTypeMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// SwitchStmt
type SwitchStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// TypeAssertExpr
type TypeAssertExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// TypeSpec
type TypeSpecMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}

// TypeSwitchStmt
type TypeSwitchStmtMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchStmt(cx *MatchContext, stmt ast.Stmt) (bool, error)
}

// UnaryExpr
type UnaryExprMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
	MatchExpr(cx *MatchContext, expr ast.Expr) (bool, error)
}

// ValueSpec
type ValueSpecMatcher interface {
	Match(cx *MatchContext, node ast.Node) (bool, error)
}
