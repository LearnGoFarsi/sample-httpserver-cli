package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/learngofarsi/go-basics-project/api"
)

type TrackRepository interface {
	Upsert(ctx context.Context, track *api.Track) (err error)
	Get(ctx context.Context) (ts []api.Track, err error)
	GetById(ctx context.Context, id string) (t api.Track, err error)
}

type TrackHandler struct {
	repo TrackRepository
}

func NewTrackHandler(repo TrackRepository) TrackHandler {
	return TrackHandler{
		repo: repo,
	}
}

func (h TrackHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.get(w, r)
	} else if r.Method == "POST" {
		h.post(w, r)
	}
}

func (h TrackHandler) post(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	track := api.Track{}
	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.repo.Upsert(r.Context(), &track); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("succeed"))
}

func (h TrackHandler) get(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	if vals.Has("track_id") {
		h.getTrackById(w, r)
	} else {
		h.getTracks(w, r)
	}
}

func (h TrackHandler) getTrackById(w http.ResponseWriter, r *http.Request) {
	trackId := r.URL.Query().Get("track_id")
	var track api.Track
	var err error

	if track, err = h.repo.GetById(r.Context(), trackId); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(track); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func (h TrackHandler) getTracks(w http.ResponseWriter, r *http.Request) {

	var err error
	var tracks []api.Track

	if tracks, err = h.repo.Get(r.Context()); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tracks); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
