package partner

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	sr := getCli()
	token := sr.ClientToken()
	fmt.Println(token)
	cr, err := List(token, "", 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cr)
}
