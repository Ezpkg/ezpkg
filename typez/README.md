<p align="center">
<a href="https://ezpkg.io">
<img alt="gopherz" src="https://ezpkg.io/_/gopherz.png" style="width:420px">
</a>
</p>

# ezpkg.io/typez

[![PkgGoDev](https://pkg.go.dev/badge/github.com/ezpkg/typez)](https://pkg.go.dev/github.com/ezpkg/typez/v2)
[![GitHub License](https://img.shields.io/github/license/ezpkg/typez)](https://github.com/ezpkg/typez/tree/main/LICENSE)
[![version](https://img.shields.io/github/v/tag/ezpkg/typez?label=version)](https://github.com/ezpkg/typez/tags)

Package typez provides generic functions for working with types.

## Installation

```sh
go get -u ezpkg.io/typez@v0.0.1
```

## Examples

```go
typez.In(1, 1, 2, 3)    // true
typez.In("A", "B", "C") // false

type A struct{X int}
typez.Coalesce(0, 1, 2, 3) // 1
typez.Coalesce(nil, &A{10}, &A{20}) // &A{10}
```

## About ezpkg.io

As I work on various Go projects, I often find myself creating utility functions, extending existing packages, or developing packages to solve specific problems. Moving from one project to another, I usually have to copy or rewrite these solutions. So I created this repository to have all these utilities and packages in one place. Hopefully, you'll find them useful as well.

For more information, see the [main repository](https://github.com/ezpkg/ezpkg).

## Author

<a href="https://olivernguyen.io"><img alt="olivernguyen.io" src="https://olivernguyen.io/_/badge.png" height="28px"></a>&nbsp;&nbsp;[![github](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/iOliverNguyen)