package partner

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	cli := getCli()
	ir, err := Info(cli.ClientToken(), "7211791254106425379")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ir)
}
