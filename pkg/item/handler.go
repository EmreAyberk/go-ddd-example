package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EmreAyberk/go-ddd-example/pkg/cache"
)

type IHandler interface {
	Create()
}
type Handle struct {
	service  Service
	memCache cache.Cache
}

func NewHandler(service Service, memoryCache cache.Cache) Handle {
	return Handle{
		service:  service,
		memCache: memoryCache,
	}
}

func (h *Handle) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		key := r.URL.Query().Get("key")

		if key == "" {
			//GET ALL
			cacheKey := h.memCache.CacheKey(r)
			cachedValue, _ := h.memCache.Get(cacheKey)

			if cachedValue != nil {
				_, err := w.Write(cachedValue)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, writeErr := w.Write([]byte(err.Error()))
					if writeErr != nil {
						return
					}
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}

			all, err := h.service.GetAll()

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}

			allItemJson, err := json.Marshal(all)
			h.memCache.Set(cacheKey, allItemJson)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
			_, err = w.Write(allItemJson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
			w.WriteHeader(http.StatusOK)
		} else {
			// GET ONE
			cacheKey := h.memCache.CacheKey(r)
			cachedValue, _ := h.memCache.Get(cacheKey)

			if cachedValue != nil {
				_, err := w.Write(cachedValue)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, writeErr := w.Write([]byte(err.Error()))
					if writeErr != nil {
						return
					}
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}

			one, err := h.service.GetOne(key)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
			itemJson, err := json.Marshal(one)
			h.memCache.Set(cacheKey, itemJson)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
			_, err = w.Write(itemJson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	case "POST":
		var item Item
		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Print(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, writeErr := w.Write([]byte(err.Error()))
			if writeErr != nil {
				return
			}
			return
		}

		cacheKey := h.memCache.CacheKey(r)
		h.memCache.Set(cacheKey, reqBody)

		err = json.Unmarshal(reqBody, &item)
		if err != nil {
			fmt.Print(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, writeErr := w.Write([]byte(err.Error()))
			if writeErr != nil {
				return
			}
			return
		}

		err = h.service.Create(item)

		if err != nil {
			return
		}
		fmt.Printf("done")
		w.WriteHeader(http.StatusCreated)
	case "DELETE":
		key := r.URL.Query().Get("key")

		if key == "" {
			//DELETE ALL
			err := h.service.DeleteAll()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
		} else {
			// DELETE ONE
			err := h.service.DeleteOne(key)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, writeErr := w.Write([]byte(err.Error()))
				if writeErr != nil {
					return
				}
				return
			}
		}
	}
}
