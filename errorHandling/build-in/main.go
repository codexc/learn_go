//build-in error 1.13 标准包error处理, 虽然可以进行包装, 但无法打印出堆栈信息
package main

import (
	"database/sql"
	"errors"
	"fmt"
)

//sentinel error
func selectRetDao1() error {
	return fmt.Errorf("select Ret Dao1:%w", sql.ErrNoRows)
}

//error type
type businessErr struct {
	err error
	msg string
}

func (e *businessErr) Error() string {
	return e.msg
}

func selectRetDao2() error {
	return fmt.Errorf("select Ret Dao2:%w", &businessErr{sql.ErrNoRows, "business err"})

}

func main() {
	err1 := selectRetDao1()
	if errors.Is(err1, sql.ErrNoRows) {
		fmt.Printf("main error1: %+v\n", err1)
	}

	err2 := selectRetDao2()
	var businessTypeErr *businessErr
	if errors.As(err2, &businessTypeErr) {
		fmt.Printf("main error2: %+v\n", err1)

	}
	return
}
