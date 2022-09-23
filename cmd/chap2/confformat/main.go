package main

import (
	confformat2 "golang/chap2/confformat"
)

func main() {
	if err := confformat2.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat2.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat2.OtherJSONExample(); err != nil {
		panic(err)
	}
}
