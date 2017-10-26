package downloader

import (
	"log"
	"module"
	"module/stub"
	"net/http"
)

type myDownloader struct {
	stub.ModuleInternal
	httpClient http.Client
}

func New(
	mid module.MID,
	client *http.Client,
	scoreCalculator module.CalculateScore) (module.Downloader, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, genParameterError("nil http client")
	}
	return &myDownloader{
		ModuleInternal: moduleBase,
		httpClient:     *client,
	}, nil
}

func (downloader *myDownloader) Download(req *module.Request) (*module.Response, error) {
	downloader.ModuleInternal.IncrHandlingNumber()
	defer downloader.ModuleInternal.DecrHandlingNumber()
	downloader.ModuleInternal.IncrCalledCount()
	if req == nil {
		return nil, genParameterError("nil request")
	}
	httpReq := req.HTTPReq()
	if httpReq == nil {
		return nil, genParameterError("nil HTTP request")
	}
	downloader.ModuleInternal.IncrAcceptedCount()
	log.Printf("Do the request (URL: %s, depth: %d)... \n", httpReq.URL, req.Depth())
	httpResp, err := downloader.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	downloader.ModuleInternal.IncrCompletedCount()
	return module.NewResponse(httpResp, req.Depth()), nil
}
