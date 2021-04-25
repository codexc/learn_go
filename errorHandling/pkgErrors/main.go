// pkg/errors包使用, Wrap包装堆栈及额外信息
package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func selectRetDao1() error {
	return errors.Wrap(sql.ErrNoRows, "pkg/errors wrap")
}

func main() {
	if err := selectRetDao1(); err != nil {
		//fmt.Printf("main errors:%v", err)
		fmt.Printf("main errors:%+v", err)
	}
	return
}
