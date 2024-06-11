package templ

import "bytes"

type DiffBuffer struct {
	bytes.Buffer

	currseg bytes.Buffer

	Delegate bool
	Segs     []string
	Vals     []string
}

func (b *DiffBuffer) Flush() {
	if b.currseg.Len() == 0 {
		return
	}
	b.Segs = append(b.Segs, b.currseg.String())
	b.currseg.Reset()
}

func (b *DiffBuffer) WriteDynamic(s string) (n int, err error) {
	b.Flush()
	b.Vals = append(b.Vals, s)
	if b.Delegate {
		return b.Buffer.WriteString(s)
	}
	return len(s), nil
}

func (b *DiffBuffer) WriteString(s string) (n int, err error) {
	n, err = b.currseg.WriteString(s)
	if err != nil {
		return n, err
	}
	if b.Delegate {
		return b.Buffer.WriteString(s)
	}
	return n, nil
}
