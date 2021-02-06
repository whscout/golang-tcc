package tcc

import (
	"tcc_example/library/tcc/runtime"
)

type RuntimeApi struct {
	UrlPattern  string
	RequestInfo *runtime.RequestInfo
	Nodes       []*RuntimeTCC
}

type RuntimeTCC struct {
	Index       int
	Try         *RuntimeNode
	Confirm     *RuntimeNode
	Cancel      *RuntimeNode
	SuccessStep *runtime.SuccessStep
}

type RuntimeNode struct {
	Url     string
	Method  string
	Timeout int
}
