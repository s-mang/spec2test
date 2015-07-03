package spec2test

type ResourceType int

const (
	One ResourceType = iota
	Many
)

type Action int

const (
	Create Action = iota
	Read
	Update
	Destroy
)

type HTTPMethod int

const (
	GET HTTPMethod = iota
	HEAD
	POST
	PUT
	DELETE
	TRACE
	OPTIONS
	CONNECT
	PATCH
)
