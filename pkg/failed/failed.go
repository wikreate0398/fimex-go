package failed

import (
	"fmt"
	"wikreate/fimex/pkg/logger"
)

func PanicOnError(err error, msg string) {
	if err != nil {
		msg = fmt.Sprintf("%v: %v", msg, err.Error())
		logger.Panic(logger.LogInput{Msg: msg})
		panic(msg)
	}
}
