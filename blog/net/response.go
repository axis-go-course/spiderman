package net

type Response interface {
	SetHeader(key, val string)
	Write([]byte) (int, error)
	WriteHeader(int)
}
