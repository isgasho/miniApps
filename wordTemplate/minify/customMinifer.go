// Package html minifies HTML5 following the specifications at http://www.w3.org/TR/html5/syntax.html.
package minify

import (
	"io"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/xml"
)

var (
	isBytes    = []byte("=")
	spaceBytes = []byte(" ")
	voidBytes  = []byte("/>")
)

// Minifier is an XML minifier.
type Minifier struct {
	KeepWhitespace bool
}

// Minify minifies XML data, it reads from r and writes to w.
func (o *Minifier) Minify(w io.Writer, r io.Reader) error {
	omitSpace := true // on true the next text token must not start with a space

	attrByteBuffer := make([]byte, 0, 64)

	l := xml.NewLexer(r)
	defer l.Restore()

	tb := NewTokenBuffer(l)
	for {
		t := *tb.Shift()
		if t.TokenType == xml.CDATAToken {
			if len(t.Text) == 0 {
				continue
			}
			if text, useText := xml.EscapeCDATAVal(&attrByteBuffer, t.Text); useText {
				t.TokenType = xml.TextToken
				t.Data = text
			}
		}
		switch t.TokenType {
		case xml.ErrorToken:
			if l.Err() == io.EOF {
				return nil
			}
			return l.Err()
		case xml.DOCTYPEToken:
			if _, err := w.Write(t.Data); err != nil {
				return err
			}
		case xml.CDATAToken:
			if _, err := w.Write(t.Data); err != nil {
				return err
			}
			if len(t.Text) > 0 && parse.IsWhitespace(t.Text[len(t.Text)-1]) {
				omitSpace = true
			}
		case xml.TextToken:
			t.Data = parse.ReplaceMultipleWhitespace(t.Data)

			// whitespace removal; trim left
			if omitSpace && (t.Data[0] == ' ' || t.Data[0] == '\n') {
				t.Data = t.Data[1:]
			}

			// whitespace removal; trim right
			omitSpace = false
			if len(t.Data) == 0 {
				omitSpace = true
			} else if t.Data[len(t.Data)-1] == ' ' || t.Data[len(t.Data)-1] == '\n' {
				omitSpace = true
				i := 0
				for {
					next := tb.Peek(i)
					// trim if EOF, text token with whitespace begin or block token
					if next.TokenType == xml.ErrorToken {
						t.Data = t.Data[:len(t.Data)-1]
						omitSpace = false
						break
					} else if next.TokenType == xml.TextToken {
						// this only happens when a comment, doctype, cdata startpi tag was in between
						// remove if the text token starts with a whitespace
						if len(next.Data) > 0 && parse.IsWhitespace(next.Data[0]) {
							t.Data = t.Data[:len(t.Data)-1]
							omitSpace = false
						}
						break
					} else if next.TokenType == xml.CDATAToken {
						if len(next.Text) > 0 && parse.IsWhitespace(next.Text[0]) {
							t.Data = t.Data[:len(t.Data)-1]
							omitSpace = false
						}
						break
					} else if next.TokenType == xml.StartTagToken || next.TokenType == xml.EndTagToken {
						if !o.KeepWhitespace {
							t.Data = t.Data[:len(t.Data)-1]
							omitSpace = false
						}
						break
					}
					i++
				}
			}

			if _, err := w.Write(t.Data); err != nil {
				return err
			}
		case xml.StartTagToken:
			if o.KeepWhitespace {
				omitSpace = false
			}
			if _, err := w.Write(t.Data); err != nil {
				return err
			}
		case xml.StartTagPIToken:
			if _, err := w.Write(t.Data); err != nil {
				return err
			}
		case xml.AttributeToken:
			if _, err := w.Write(spaceBytes); err != nil {
				return err
			}
			if _, err := w.Write(t.Text); err != nil {
				return err
			}
			if _, err := w.Write(isBytes); err != nil {
				return err
			}

			if len(t.AttrVal) < 2 {
				if _, err := w.Write(t.AttrVal); err != nil {
					return err
				}
			} else {
				// prefer single or double quotes depending on what occurs more often in value
				val := xml.EscapeAttrVal(&attrByteBuffer, t.AttrVal[1:len(t.AttrVal)-1])
				if _, err := w.Write(val); err != nil {
					return err
				}
			}
		case xml.StartTagCloseToken:
			next := tb.Peek(0)
			if next.TokenType == xml.TextToken && parse.IsAllWhitespace(next.Data) {
				next = tb.Peek(1)
			}
			if _, err := w.Write(t.Text); err != nil {
				return err
			}
		case xml.StartTagCloseVoidToken:
			if _, err := w.Write(t.Text); err != nil {
				return err
			}
		case xml.StartTagClosePIToken:
			if _, err := w.Write(t.Text); err != nil {
				return err
			}
		case xml.EndTagToken:
			if o.KeepWhitespace {
				omitSpace = false
			}
			if len(t.Data) > 3+len(t.Text) {
				t.Data[2+len(t.Text)] = '>'
				t.Data = t.Data[:3+len(t.Text)]
			}
			if _, err := w.Write(t.Data); err != nil {
				return err
			}
		}
	}
}
