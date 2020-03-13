package api

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"product-service/core/service"
	"product-service/scenario/business"
	"product-service/scenario/serializer/json"
	"product-service/scenario/serializer/msg"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	sneakerProductService business.SneakerProductService
}

func NewHandler(sneakerProductService business.SneakerProductService) RedirectHandler {
	return &handler{sneakerProductService: sneakerProductService}
}


func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) service.SneakerProductSerializer {
	if contentType == "application/x-msgpack" {
		return &msg.SneakerProduct{}
	}
	return &json.SneakerProduct{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	sneakerProduct, err := h.sneakerProductService.Retrieve(code)
	if err != nil {
		if errors.Cause(err) == business.ErrProductNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	contentType := "application/json"
	responseBody, err := h.serializer(contentType).Encode(sneakerProduct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	sneakerProduct, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.sneakerProductService.Store(sneakerProduct)
	if err != nil {
		if errors.Cause(err) == business.ErrProductInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(sneakerProduct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusOK)
}
