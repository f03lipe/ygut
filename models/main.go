package models

import (
	"fmt"

	"github.com/f03lipe/ygut/conf"
)

func Setup() {
	fmt.Printf("%+v", conf.C)
}
