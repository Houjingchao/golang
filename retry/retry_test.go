package retry

import (
	"errors"
	"fmt"
	"testing"
)

var cnt = 0

func TestRetryy(t *testing.T) {
	var err error
	a := func() error {
		if cnt == 3 {
			fmt.Println("cnt ==", cnt)
			return nil
		}

		fmt.Println("cnt ==", cnt)
		cnt++
		return errors.New("test err")
	}
	err = Retry(3, 0, a)
	fmt.Println(err)
}
