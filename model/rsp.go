package model

type BaseRsp interface {
	GetCode() int
	GetMsg() string
}
