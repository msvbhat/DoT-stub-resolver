package main

import (
	"context"
	"github.com/artyom/dot"
	"log"
)

func resolve(name string) ([]string, error) {
	resolver := dot.Cloudflare()
	ctx := context.Background()
	log.Println(name)
	ips, err := resolver.LookupHost(ctx, name)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
