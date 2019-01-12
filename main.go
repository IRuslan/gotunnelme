package main

import (
	"github.com/iruslan/gotunnelme/src/gotunnelme"
	"os"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "gotunnelme <local port> or gotunnelme <local port> <subdomain>")
		os.Exit(1)
	}
	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	d := ""
	if len(os.Args) > 2 {
		d = os.Args[2]
	}

	t := gotunnelme.NewTunnel()
	url, err := t.GetUrl(d)
	if err != nil {
		panic(err)
	}
	print(url)
	err = t.CreateTunnel(p)
	if err != nil {
		panic(err)
	}
	t.StopTunnel()
}
