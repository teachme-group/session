package closer

type Closer interface {
	AddCloser(closer func())
	AddErrCloser(closer func() error)
}

type closer struct {
	closers    []func()
	errClosers []func() error
}

func New() *closer {
	return &closer{}
}

func (c *closer) AddCloser(closer func()) {
	c.closers = append(c.closers, closer)
}

func (c *closer) AddErrCloser(closer func() error) {
	c.errClosers = append(c.errClosers, closer)
}

func (c *closer) Close() error {
	for _, cl := range c.closers {
		cl()
	}

	for _, cl := range c.errClosers {
		if err := cl(); err != nil {
			return err
		}
	}

	return nil
}
