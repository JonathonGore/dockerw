package dockerw

import (
	"fmt"
	"strings"
)

type Copy struct {
	value string
}

func (c *Copy) verify() error {
	return nil
}

func (c *Copy) toString() string {
	return c.value
}

// Copy consumes two directories to form a CP directive in the Dockerfile
func (w *Writer) Copy(d1, d2 string) *Writer {
	d1 = strings.TrimSpace(d1)
	d2 = strings.TrimSpace(d2)

	d := fmt.Sprintf("COPY %v %v", d1, d2)

	w.addCommand(&From{d})

	return w
}
