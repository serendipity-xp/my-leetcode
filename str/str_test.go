package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_str(t *testing.T) {
	t.Run("FindNeedleStr", func(t *testing.T) {
		index := FindNeedleStr("abcdefg", "cd")
		assert.Equal(t,index,2)
		index = FindNeedleStr("abcdefg", "")
		assert.Equal(t,index,0)
		index = FindNeedleStr("abcdefg", "dc")
		assert.Equal(t,index,-1)
	})
}

func TestStrRune(t *testing.T) {
	var a  = rune("a"[0])
	fmt.Println(a)
}
