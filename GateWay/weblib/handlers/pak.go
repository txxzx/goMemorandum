package handlers

import (
	"errors"

	"github.com/txxzx/goMemorandum/GateWay/pkg/loggings"
)

/**
    @date: 2024/7/14
**/

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		loggings.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		loggings.Info(err)
		panic(err)
	}
}
