package transport

import (
	"encoding/json"
	"net/http"

	"balance/internal/controller"
)

type transport struct {
	svc *controller.Controller
}

func NewHTTP(svc *controller.Controller) *transport {
	return &transport{
		svc: svc,
	}
}

func (s *transport) GetBalance(w http.ResponseWriter, r *http.Request) {
	var (
		req controller.RequestBalance
		res response
	)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	balance, err := s.svc.GetBalance(r.Context(), req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	res.ResponseData = balance
	body, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *transport) AddBalance(w http.ResponseWriter, r *http.Request) {
	var (
		req controller.RequestBalance
		res response
	)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	balance, err := s.svc.AddBalance(r.Context(), req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	res.ResponseData = balance
	body, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func (s *transport) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		req controller.Order
		res response
	)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	err = s.svc.CreateOrder(r.Context(), req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (s *transport) FinishOrder(w http.ResponseWriter, r *http.Request) {
	var (
		req controller.Order
		res response
	)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	err = s.svc.FinishOrder(r.Context(), req)
	if err != nil {
		res.ErrMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(res)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
}
