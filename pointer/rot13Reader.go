package pointer

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func ROT13(b byte)  byte    {
    if b >= 65 && b <= 90   {
        return 65 + ((b - 65 + 13) % 26)
    }else if b >= 97 && b <= 122    {
        return 97 + ((b - 97 + 13) % 26)
    }
    return b
}

func (rrd rot13Reader) Read(b []byte) (int, error)	{
    _len, err := rrd.r.Read(b)
    if err != nil   {
        return 0, err
    }
    for i := range b[:_len] {
        b[i] = ROT13(b[i])
    }
    return _len, nil
}

func Rot13()    {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
