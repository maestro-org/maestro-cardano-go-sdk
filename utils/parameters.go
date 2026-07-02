package utils

import (
	"fmt"
	"strings"
)

type Parameters struct {
	params []string
}

func NewParameters() *Parameters {
	return &Parameters{
		params: make([]string, 0),
	}
}
func (p *Parameters) Format() string {
	if len(p.params) == 0 {
		return ""
	}
	return "?" + strings.Join(p.params, "&&")
}

func (p *Parameters) Asset(asset string) {
	p.params = append(p.params, "asset="+asset)
}

func (p *Parameters) Count(amount int) {
	p.params = append(p.params, "count="+fmt.Sprint(amount))
}

func (p *Parameters) Cursor(cursor string) {
	p.params = append(p.params, "cursor="+cursor)
}

func (p *Parameters) Policy(policy string) {
	p.params = append(p.params, "policy="+policy)
}

func (p *Parameters) EpochNo(epochNo int64) {
	p.params = append(p.params, "epoch_no="+fmt.Sprint(epochNo))
}

func (p *Parameters) From(from int64) {
	p.params = append(p.params, "from="+fmt.Sprint(from))
}

func (p *Parameters) To(to int64) {
	p.params = append(p.params, "to="+fmt.Sprint(to))
}

func (p *Parameters) SetAscOrder() {
	p.params = append(p.params, "order=asc")
}

func (p *Parameters) SetDescOrder() {
	p.params = append(p.params, "order=desc")
}

func (p *Parameters) WithCbor() {
	p.params = append(p.params, "with_cbor=true")
}

func (p *Parameters) ResolveDatums() {
	p.params = append(p.params, "resolve_datums=true")
}

func (p *Parameters) FromHeight(height int64) {
	p.params = append(p.params, "from_height="+fmt.Sprint(height))
}
