package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	jsonYaml = flag.Bool("y", false, "convert json to yaml")
)

func main() {
	flag.Parse()
	switch {
	case *jsonYaml:
		var (
			m   = make(map[string]interface{})
			dec = json.NewDecoder(os.Stdin)
		)
		if err := dec.Decode(&m); err != nil {
			log.Fatal(err)
		}
		b, err := yaml.Marshal(&m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	default:
		log.Fatal("use -h to see conversion options")
	}
}
