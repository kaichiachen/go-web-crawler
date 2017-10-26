package buffer

import "errors"

// ErrClosedBufferPool 是表示缓冲池已关闭的错误的变量。
var ErrClosedBufferPool = errors.New("closed buffer pool")
