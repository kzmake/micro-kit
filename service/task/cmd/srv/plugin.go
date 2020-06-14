package main

import (
	// service registry
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
)
