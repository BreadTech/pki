package types

import (
	"encoding/asn1"
)

type ASN1 struct {
	Values []*asn1.RawValue
}

func (a *ASN1) push(v *asn1.RawValue) {
	a.Values = append(a.Values, v)
}

func (a *ASN1) pop() (v *asn1.RawValue) {
	n := a.length()
	v = a.Values[n-1]
	a.Values = a.Values[:n-1]
	return
}

func (a *ASN1) length() int {
	return len(a.Values) - 1
}

func (a *ASN1) isEmpty(v *asn1.RawValue) bool {
	return a.length() == 0
}

func ParseASN1(dat []byte) (*ASN1, []byte, error) {
	out := new(ASN1)
	stack := new(ASN1)

	root := new(asn1.RawValue)
	asnRest, err := asn1.Unmarshal(dat, root)
	if err != nil {
		return nil, asnRest, err
	}

	for {
		cur := stack.pop()

		next := new(asn1.RawValue)
		asnRest, err := asn1.Unmarshal(cur.Bytes, next)
		if err != nil {
			return nil, asnRest, err
		}
		if next.Tag == asn1.TagSequence {
			stack.push(next)
		}

		for len(asnRest) > 0 {
			more := new(asn1.RawValue)
		}

	}
	return out, asnRest, nil
}
