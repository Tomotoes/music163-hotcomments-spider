package util

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func Check(err error) {
	if err != nil {
		errWrap := errors.Wrap(err,"will exit")
		fmt.Printf("%+v\n", errWrap)
		os.Exit(1)
	}
}
