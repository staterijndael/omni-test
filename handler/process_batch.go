package handler

import (
	"context"
	"net/http"
	"omni-test/service"
	"strconv"
)

// ProcessBatch handles the processing of a batch of items.
//
//	@Summary		Process a batch of items
//	@Description	Processes a batch of items based on the provided batch size.
//	@Produce		plain
//	@Param			batch_size	query		int		true	"Batch size for processing"
//	@Success		200			{string}	string	"Batch successfully handled!"
//	@Failure		400			{string}	string	"Bad Request"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/process_batch [get]
func (h *Handler) ProcessBatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	batchSizeStr := r.URL.Query().Get("batch_size")
	if batchSizeStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("batch_size parameter is not provided"))
		if err != nil {
			return
		}
		return
	}

	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("batch_size parameter must be integer"))
		if err != nil {
			return
		}
		return
	}

	batch := make(service.Batch, batchSize)

	err = h.processor.ProcessBatch(ctx, batch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("batch successfully handled!"))
	if err != nil {
		return
	}
}
