package user

import (
	"github.com/k0kubun/pp"
)

func (u User) PrettyPrint() {
	pp.Println(u)
}
