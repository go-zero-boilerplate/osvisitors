package osvisitors

import (
	"fmt"
	"runtime"
	"strings"
)

func GetRuntimeOsType() (OsType, error) {
	for _, os := range AllList {
		visitor := &GoOSNameVisitor{}
		os.Accept(visitor)
		if strings.EqualFold(runtime.GOOS, visitor.name) {
			return os, nil
		}
	}
	return nil, fmt.Errorf("github.com/go-zero-boilerplate/osvisitors does not currently support OS name '%s'", runtime.GOOS)
}
