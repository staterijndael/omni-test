package handler

import "omni-test/service/processor"

type Handler struct {
	processor *processor.Processor
}

func NewHandler(processor *processor.Processor) *Handler {
	return &Handler{
		processor: processor,
	}
}
