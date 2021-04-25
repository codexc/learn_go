//xerrors里Errorf 使用%w来包装error
package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/xerrors"
)

func selectRetDao1() error {
	return xerrors.Errorf("xerrors wrap: %w", sql.ErrNoRows)
}

func main() {
	if err := selectRetDao1(); err != nil {
		fmt.Printf("main errors:%+v\n", err)
	}
	return
}
