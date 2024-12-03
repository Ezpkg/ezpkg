package jsonz

import (
	"fmt"
	"strconv"

	"ezpkg.io/fmtz"
)

type Item struct {
	Level int      // level of indentation
	Index int      // index in the parent array or object
	Key   RawToken // optional object "key"
	Token RawToken // [ or { or } or ] or , or value

	path RawPath
}

type RawPath []PathItem

type PathItem struct {
	Index int
	Token RawToken
	Key   RawToken
}

func (x Item) String() string {
	return x.Token.String()
}

func (x Item) Format(f fmt.State, c rune) {
	fz := fmtz.WrapState(f)
	switch {
	case f.Flag('+'):
		fz.Printf("%s → %s", x.path, x.Token)

	default:
		fz.Print(x.Token)
	}
}

func (x Item) IsArrayValue() bool {
	return !x.Key.IsValue() && x.Token.IsValue()
}

func (x Item) IsObjectValue() bool {
	return x.Key.IsValue() && x.Token.IsValue()
}

func (x Item) GetValue() (any, error) {
	return x.Token.GetValue()
}

func (x Item) GetTokenValue() (any, error) {
	return x.Token.GetValue()
}

// GetPath returns the path of the item as a slice of values. The values are the keys of objects (string) and the indexes of arrays (int).
func (x Item) GetPath() []any {
	path := make([]any, len(x.path))
	for i, item := range x.path {
		path[i] = item.Value()
	}
	return path
}

// GetRawPath returns the path of the item as a slice of PathItem.
// IMPORTANT: The result slice should not be modified.
func (x Item) GetRawPath() RawPath {
	return x.path
}

// Path returns the path of the item as a string: [0]."key"[1]...
func (p RawPath) String() string {
	return fmt.Sprint(p)
}

// Format formats the path as a string: [0]."key"[1]...
func (p RawPath) Format(f fmt.State, c rune) {
	fz := fmtz.WrapState(f)
	for _, item := range p {
		fz.Printf("%+s", item)
	}
}

// IsArray returns true if the path item is inside an array.
func (p PathItem) IsArray() bool {
	return p.Token.typ == TokenArrayStart
}

// IsObject returns true if the path item is inside an object.
func (p PathItem) IsObject() bool {
	return p.Token.typ == TokenObjectStart
}

// Value returns the value of the path item. If the item is inside an array, it returns the index. If the item is inside an object, it returns the key.
func (p PathItem) Value() any {
	switch p.Token.typ {
	case TokenArrayStart:
		return p.Index
	case TokenObjectStart:
		v, _ := p.Token.GetString()
		return v
	default:
		return nil
	}
}

// String returns the string representation of the path item. "[0]" for array, ".key" for object.
func (p PathItem) String() string {
	switch {
	case p.IsArray():
		return strconv.Itoa(p.Index)
	case p.IsObject():
		rawKey := p.Key.Raw()
		if canSimplyUnquote(rawKey) {
			rawKey = rawKey[1 : len(rawKey)-1]
		}
		return string(rawKey)
	default:
		return ""
	}
}

// Format formats the path item as a string. "[0]" for array, ".key" for object.
func (p PathItem) Format(f fmt.State, c rune) {
	fz := fmtz.WrapState(f)
	switch {
	case p.IsArray():
		if f.Flag('+') {
			fz.Printf("[%d]", p.Index)
		} else {
			fz.Printf("%d", p.Index)
		}
	case p.IsObject():
		rawKey := p.Key.Raw()
		if canSimplyUnquote(rawKey) {
			rawKey = rawKey[1 : len(rawKey)-1]
		}
		if f.Flag('+') {
			fz.Printf(".%s", rawKey)
		} else {
			fz.Printf("%s", rawKey)
		}
	default:
		// root, empty path
	}
}
