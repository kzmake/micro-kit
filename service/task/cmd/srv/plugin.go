package main

import (
	//
	// service registry
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/etcd/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	// _ "go.uber.org/automaxprocs"
	//	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/micro/go-plugins/v2/broker/googlepubsub"
)
