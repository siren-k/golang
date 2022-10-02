package main

import (
	dataconv2 "golang/internal/03_data_processing/dataconv"
)

func main() {
	dataconv2.ShowConv()
	if err := dataconv2.Strconv(); err != nil {
		panic(err)
	}
	dataconv2.Interfaces()
}
