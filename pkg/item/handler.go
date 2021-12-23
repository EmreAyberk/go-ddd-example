package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IHandler interface {
	Create()
}
type Handle struct {
	service Service
}

func NewHandler(service Service) Handle {
	return Handle{service: service}
}

func (h *Handle) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		key := r.URL.Query().Get("key")

		if key == "" {
			//GET ALL
			all, err := h.service.GetAll()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			allItemJson, err := json.Marshal(all)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			_, err = w.Write(allItemJson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			w.WriteHeader(http.StatusOK)
		} else {
			// GET ONE
			one, err := h.service.GetOne(key)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			itemJson, err := json.Marshal(one)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			_, err = w.Write(itemJson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
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
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}

		err = json.Unmarshal(reqBody, &item)
		if err != nil {
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
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
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
		} else {
			// DELETE ONE
			err := h.service.DeleteOne(key)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
		}
	}
}
