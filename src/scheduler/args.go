package scheduler

import (
	"module"
)

type RequestArgs struct {
	AcceptedDomains []string `json:"accepted_primary_domains"`
	MaxDepth        uint32   `json:"max_depth"`
}

// Same 用于判断两个请求相关的参数容器是否相同。
func (args *RequestArgs) Same(another *RequestArgs) bool {
	if another == nil {
		return false
	}
	if another.MaxDepth != args.MaxDepth {
		return false
	}
	anotherDomains := another.AcceptedDomains
	anotherDomainsLen := len(anotherDomains)
	if anotherDomainsLen != len(args.AcceptedDomains) {
		return false
	}
	if anotherDomainsLen > 0 {
		for i, domain := range anotherDomains {
			if domain != args.AcceptedDomains[i] {
				return false
			}
		}
	}
	return true
}

type DataArgs struct {
	ReqBufferCap         uint32 `json:"req_buffer_cap"`
	ReqMaxBufferNumber   uint32 `json:"req_max_buffer_number"`
	RespBufferCap        uint32 `json:"resp_buffer_cap"`
	RespMaxBufferNumber  uint32 `json:"resp_max_buffer_number"`
	ItemBufferCap        uint32 `json:"item_buffer_cap"`
	ItemMaxBufferNumber  uint32 `json:"item_max_buffer_number"`
	ErrorBufferCap       uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

type ModuleArgs struct {
	Downloaders []module.Downloader
	Analyzers   []module.Analyzer
	Pipelines   []module.Pipeline
}

type Args interface {
	Check() error
}

func (args *RequestArgs) Check() error {
	if args.AcceptedDomains == nil {
		return genError("nil accepted primary domain list")
	}
	return nil
}

func (args *DataArgs) Check() error {
	if args.ReqBufferCap == 0 {
		return genError("zero request buffer capacity")
	}
	if args.ReqMaxBufferNumber == 0 {
		return genError("zero max request buffer number")
	}
	if args.RespBufferCap == 0 {
		return genError("zero response buffer capacity")
	}
	if args.RespMaxBufferNumber == 0 {
		return genError("zero max response buffer number")
	}
	if args.ItemBufferCap == 0 {
		return genError("zero item buffer capacity")
	}
	if args.ItemMaxBufferNumber == 0 {
		return genError("zero max item buffer number")
	}
	if args.ErrorBufferCap == 0 {
		return genError("zero error buffer capacity")
	}
	if args.ErrorMaxBufferNumber == 0 {
		return genError("zero max error buffer number")
	}
	return nil
}

// Check 用于当前参数容器的有效性。
func (args *ModuleArgs) Check() error {
	if len(args.Downloaders) == 0 {
		return genError("empty downloader list")
	}
	if len(args.Analyzers) == 0 {
		return genError("empty analyzer list")
	}
	if len(args.Pipelines) == 0 {
		return genError("empty pipeline list")
	}
	return nil
}

// ModuleArgsSummary 代表组件相关的参数容器的摘要类型。
type ModuleArgsSummary struct {
	DownloaderListSize int `json:"downloader_list_size"`
	AnalyzerListSize   int `json:"analyzer_List_size"`
	PipelineListSize   int `json:"pipeline_list_size"`
}

func (args *ModuleArgs) Summary() ModuleArgsSummary {
	return ModuleArgsSummary{
		DownloaderListSize: len(args.Downloaders),
		AnalyzerListSize:   len(args.Analyzers),
		PipelineListSize:   len(args.Pipelines),
	}
}
