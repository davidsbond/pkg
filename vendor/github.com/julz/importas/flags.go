package importas

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func flags() flag.FlagSet {
	fs := flag.FlagSet{}
	fs.Var(stringMap(config.RequiredAlias), "alias", "required import alias in form path:alias")
	return fs
}

type stringMap map[string]string

func (v stringMap) Set(val string) error {
	spl := strings.SplitN(val, ":", 2)
	if len(spl) != 2 {
		return errors.New("import flag must be of form path:alias")
	}

	v[spl[0]] = spl[1]
	return nil
}

func (v stringMap) String() string {
	return fmt.Sprintf("%v", (map[string]string)(v))
}
