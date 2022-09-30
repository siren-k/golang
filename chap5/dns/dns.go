package dns

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

// Lookup 구조체는 우리가 관심을 갖는 DNS 정보를 저장한다.
type Lookup struct {
	cname string
	hosts []string
}

// lookup 객체를 출력하는데 이 함수를 사용할 수 있다.
func (d *Lookup) String() string {
	result := ""
	for _, host := range d.hosts {
		result += fmt.Sprintf("%s IN A %s\n", d.cname, host)
	}
	return result
}

// LookupAddress 함수는 입력된 주소에 대한 cname과 host로 구성된
// DNSLookup을 반환한다.
func LookupAddress(address string) (*Lookup, error) {
	cname, err := net.LookupCNAME(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}

	hosts, err := net.LookupHost(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname: cname, hosts: hosts}, nil
}
