package scheduler

import (
	"errs"
	"log"
	"module"
	"toolkit/buffer"
)

// genError 用于生成爬虫错误值。
func genError(errMsg string) error {
	return errs.NewCrawlerError(errs.ERROR_TYPE_SCHEDULER,
		errMsg)
}

// genErrorByError 用于基于给定的错误值生成爬虫错误值。
func genErrorByError(err error) error {
	return errs.NewCrawlerError(errs.ERROR_TYPE_SCHEDULER,
		err.Error())
}

// genParameterError 用于生成爬虫参数错误值。
func genParameterError(errMsg string) error {
	return errs.NewCrawlerErrorBy(errs.ERROR_TYPE_SCHEDULER,
		errs.NewIllegalParameterError(errMsg))
}

// sendError 用于向错误缓冲池发送错误值。
func sendError(err error, mid module.MID, errorBufferPool buffer.Pool) bool {
	if err == nil || errorBufferPool == nil || errorBufferPool.Closed() {
		return false
	}
	var crawlerError errs.CrawlerError
	var ok bool
	crawlerError, ok = err.(errs.CrawlerError)
	if !ok {
		var moduleType module.Type
		var errorType errs.ErrorType
		ok, moduleType = module.GetType(mid)
		if !ok {
			errorType = errs.ERROR_TYPE_SCHEDULER
		} else {
			switch moduleType {
			case module.TYPE_DOWNLOADER:
				errorType = errs.ERROR_TYPE_DOWNLOADER
			case module.TYPE_ANALYZER:
				errorType = errs.ERROR_TYPE_ANALYZER
			case module.TYPE_PIPELINE:
				errorType = errs.ERROR_TYPE_PIPELINE
			}
		}
		crawlerError = errs.NewCrawlerError(errorType, err.Error())
	}
	if errorBufferPool.Closed() {
		return false
	}
	go func(crawlerError errs.CrawlerError) {
		if err := errorBufferPool.Put(crawlerError); err != nil {
			log.Println("The error buffer pool was closed. Ignore error sending.")
		}
	}(crawlerError)
	return true
}
