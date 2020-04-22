package gapublic

import (
	"bytes"
	"encoding/json"
	"os"
)

func PrintJsonStringPretty(b string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(b), "", "\t")
	if err != nil {
		panic(err)
	}

	out.WriteTo(os.Stdout)
}

func PringJsonSliceStringPretty(s []string) {
	for _, v := range s {
		PrintJsonStringPretty(v)
	}
}
