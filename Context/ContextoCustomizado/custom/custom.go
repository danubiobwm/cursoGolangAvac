package custom

import (
	"context"
	"time"
)

type Context struct {
	pai     context.Context
	valores map[string]any
}

func NewContext(pai context.Context, valores map[string]any) context.Context {
	return &Context{
		pai:     pai,
		valores: valores,
	}
}

// Deadline retorna um deadline indefinido (ignora o do pai)
func (c *Context) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

// Done ignora o canal Done do contexto pai
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err retorna um erro vazio
func (c *Context) Err() error {
	return nil
}

// Value retorna um valor armazenado no contexto ou procura no pai (comportamento padrão)
func (c *Context) Value(key any) any {
	if v, ok := c.valores[key.(string)]; ok {
		return v
	}

	return c.pai.Value(key)
}
