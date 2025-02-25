package domain

import "github.com/chenmingyong0423/go-mongox/v2"

type Posts struct {
	mongox.Model `bson:",inline"`
	Title        string
	Content      string
	Author       string
	Category     string
	Tags         []string
	Cover        string
	IsTop        bool
}
