package handler

import (
	"log"
	"net/http"

	"data-maturity-assessment/internal/markdown"
	"data-maturity-assessment/internal/middleware"
	"data-maturity-assessment/internal/service"
	"data-maturity-assessment/views/components"
	"data-maturity-assessment/views/pages"
)

// TopHandler は GET / および GET /partials/model-description を処理する。
type TopHandler struct {
	svc *service.MaturityService
}

// NewTopHandler は TopHandler を構築する。
func NewTopHandler(svc *service.MaturityService) *TopHandler {
	return &TopHandler{svc: svc}
}

// Top は GET / のハンドラ。トップページを返す。
func (h *TopHandler) Top(w http.ResponseWriter, r *http.Request) {
	models := h.svc.ListModels()

	if middleware.IsHtmx(r) {
		if err := pages.TopBody(models).Render(r.Context(), w); err != nil {
			log.Printf("error rendering top body: %v", err)
		}
		return
	}
	if err := pages.Top(models).Render(r.Context(), w); err != nil {
		log.Printf("error rendering top page: %v", err)
	}
}

// ModelDescription は GET /partials/model-description のハンドラ。
// 評価モデルの説明文 Partial を返す。HX-Request ヘッダーがない場合は / にリダイレクトする。
func (h *TopHandler) ModelDescription(w http.ResponseWriter, r *http.Request) {
	if !middleware.IsHtmx(r) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	modelKey := r.URL.Query().Get("model")
	m, ok := h.svc.GetModel(modelKey)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := components.ModelDescription(markdown.ToHTML(m.ModelDescription)).Render(r.Context(), w); err != nil {
		log.Printf("error rendering model description: %v", err)
	}
}
