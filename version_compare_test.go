package uversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGLte(t *testing.T) {
	a := assert.New(t)
	a.True(Compare("2.10", ">=", "2.9"))
	a.True(Compare("2.9", ">=", "2.9"))
	a.True(Compare("2", ">", "1"))
	a.True(Compare("2.99", ">", "1.3"))
	a.False(Compare("2.1", ">=", "2.3"))
	a.True(Compare("2.99", ">", "1.3"))
	a.True(Compare("2.991.0", ">", "2.99"))
	a.False(Compare("3.3.8.3", ">=", "3.3.9.0"))
	a.False(Compare("2.99.1", "==", "2.99"))
	a.True("3.3" < "4")
	a.False(Compare("2.df", "<=", "1.3"))
	a.True(Compare("1.2.999", "<", "1.3.0"))
	a.True(Compare("2.1.99", "<=", "2.5.0"))
	a.False(Compare("2.2", "<=", "2.1"))
	a.True(Compare("1.99", "<", "2"))
	a.True(Compare("3.3.4.4", "<", "3.3.6.0"))
	a.False(Compare("3.3.9.6", "<=", "3.3.8.0"))

	a.True(Between("1.3", "<", "2", "<", "3"))
	a.True(Between("1.9", "<", "2.0", "<", "3"))

}
