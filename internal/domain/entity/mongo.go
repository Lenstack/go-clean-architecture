package entity

import "context"

type Mongo struct {
	Context        context.Context
	CollectionName string
	Interface      interface{}
	Filter         interface{}
	Options        interface{}
}
