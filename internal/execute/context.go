package execute

import (
	"context"
)

type Context struct {
	context.Context
	activities []any
}

func (c *Context) Activity(fun any) {

}
