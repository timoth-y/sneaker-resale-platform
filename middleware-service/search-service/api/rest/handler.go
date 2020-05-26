package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"search-service/core/meta"
	"search-service/core/model"
	"search-service/core/service"
	"search-service/env"
	"search-service/usecase/business"
	"search-service/usecase/serializer/json"
	"search-service/usecase/serializer/msg"
)

type RestfulHandler interface {
	// Endpoint handlers:
	Get(http.ResponseWriter, *http.Request)
	GetBy(http.ResponseWriter, *http.Request)
	GetSKU(http.ResponseWriter, *http.Request)
	GetBrand(http.ResponseWriter, *http.Request)
	GetModel(http.ResponseWriter, *http.Request)
	PostOne(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	PostAll(http.ResponseWriter, *http.Request)
	PostQuery(http.ResponseWriter, *http.Request)
	// Middleware:
	Authenticator(next http.Handler) http.Handler
}

type handler struct {
	search      service.ReferenceSearchService
	sync        service.ReferenceSyncService
	auth        service.AuthService
	contentType string
}

func NewHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.CommonConfig) RestfulHandler {
	return &handler{
		search,
		sync,
		auth,
		config.ContentType,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"][0]
	params := NewRequestParams(r)

	ref, err := h.search.Search(query, params)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, ref, http.StatusOK)
}

func (h *handler) GetBy(w http.ResponseWriter, r *http.Request) {
	field := chi.URLParam(r,"field")
	query := r.URL.Query()["query"][0]
	params := NewRequestParams(r)

	refs, err := h.search.SearchBy(field, query, params)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetSKU(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	params := NewRequestParams(r)

	refs, err := h.search.SearchSKU(sku, params)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	params := NewRequestParams(r)

	refs, err := h.search.SearchBrand(brand, params)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) GetModel(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")
	params := NewRequestParams(r)

	refs, err := h.search.SearchModel(model, params)
	if err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, refs, http.StatusOK)
}

func (h *handler) PostOne(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "referenceId")
	if err := h.sync.SyncOne(code);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["referenceId"]
	params := NewRequestParams(r)
	if err := h.sync.Sync(codes, params);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) PostAll(w http.ResponseWriter, r *http.Request) {
	params := NewRequestParams(r)
	if err := h.sync.SyncAll(params);  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) PostQuery(w http.ResponseWriter, r *http.Request) {
	query, err := h.getRequestQuery(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := NewRequestParams(r)
	if err := h.sync.SyncQuery(query, params);  err != nil {
		if errors.Cause(err) == business.ErrReferenceNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.contentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.contentType).Encode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(raw); err != nil {
			log.Println(err)
		}
	}
}

func (h *handler) getRequestBody(r *http.Request) (*model.SneakerReference, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeReference(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *handler) getRequestQuery(r *http.Request) (meta.RequestQuery, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).DecodeMap(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *handler) serializer(contentType string) service.SneakerSearchSerializer {
	if contentType == "application/x-msgpack" {
		return msg.NewSerializer()
	}
	return json.NewSerializer()
}