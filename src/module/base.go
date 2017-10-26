package module

import (
	"net/http"
)

// Counts 代表用于汇集组件内部计数的类型。
type Counts struct {
	// CalledCount 代表调用计数。
	CalledCount uint64
	// AcceptedCount 代表接受计数。
	AcceptedCount uint64
	// CompletedCount 代表成功完成计数。
	CompletedCount uint64
	// HandlingNumber 代表实时处理数。
	HandlingNumber uint64
}

type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra,omitempty"`
}

type MID string
type Module interface {
	ID() MID
	Addr() string
	Score() uint64
	SetScore(score uint64)
	ScoreCalculator() CalculateScore
	CalledCount() uint64
	AcceptedCount() uint64
	CompletedCount() uint64
	HandlingNumber() uint64
	Counts() Counts
	Summary() SummaryStruct
}

type Downloader interface {
	Module
	Download(req *Request) (*Response, error)
}

type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]Data, []error)
type Analyzer interface {
	Module
	RespParsers() []ParseResponse
	Analyze(resp *Response) ([]Data, []error)
}

type ProcessItem func(item Item) (result Item, err error)
type Pipeline interface {
	Module
	ItemProcessors() []ProcessItem
	Send(item Item) []error
	FailFast() bool
	SetFailFast(failFast bool)
}
