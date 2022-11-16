package main

// 接口嵌套接口

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	UnLock()
}
type File interface {
	ReadWrite
	Lock
	close()
}
