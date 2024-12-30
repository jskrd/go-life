package models

import (
	"syscall/js"
)

type Canvas struct {
	Context js.Value
	Element js.Value
}
