package handler

import (
	"log"
	"net/http"
	"sort"

	"data-maturity-assessment/internal/middleware"
	"data-maturity-assessment/internal/model"
	"data-maturity-assessment/internal/service"
	"data-maturity-assessment/views/components"
	"data-maturity-assessment/views/pages"
)

// AssessmentHandler は GET /assessment および GET /partials/questions を処理する。
type AssessmentHandler struct {
	svc *service.MaturityService
}

// NewAssessmentHandler は AssessmentHandler を構築する。
func NewAssessmentHandler(svc *service.MaturityService) *AssessmentHandler {
	return &AssessmentHandler{svc: svc}
}

// Assessment は GET /assessment のハンドラ。
// model クエリパラメータが不正な場合は / にリダイレクトする。
// HX-Request ヘッダーがある場合は Partial のみ返す。
func (h *AssessmentHandler) Assessment(w http.ResponseWriter, r *http.Request) {
	modelKey := r.URL.Query().Get("model")
	m, ok := h.svc.GetModel(modelKey)
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	allTopics := topicNames(m)
	groups := buildTopicGroups(m, h.svc.LevelNames(m), h.svc.FilterAssessments(m, allTopics))

	if middleware.IsHtmx(r) {
		if err := pages.AssessmentBody(m, groups).Render(r.Context(), w); err != nil {
			log.Printf("error rendering assessment body: %v", err)
		}
		return
	}
	if err := pages.Assessment(m, groups).Render(r.Context(), w); err != nil {
		log.Printf("error rendering assessment page: %v", err)
	}
}

// Questions は GET /partials/questions のハンドラ。
// 選択されたトピックエリアに絞った質問一覧 Partial を返す。
// HX-Request ヘッダーがない場合は /assessment にリダイレクトする。
func (h *AssessmentHandler) Questions(w http.ResponseWriter, r *http.Request) {
	modelKey := r.URL.Query().Get("model")
	m, ok := h.svc.GetModel(modelKey)
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if !middleware.IsHtmx(r) {
		http.Redirect(w, r, "/assessment?model="+modelKey, http.StatusFound)
		return
	}

	selected := r.URL.Query()["topicAreas"]
	if len(selected) == 0 {
		w.Header().Set("HX-Trigger", `{"validationError":"少なくとも1つのトピックエリアを選択してください。"}`)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	assessments := h.svc.FilterAssessments(m, selected)
	groups := buildTopicGroups(m, h.svc.LevelNames(m), assessments)

	if err := components.QuestionAccordion(groups).Render(r.Context(), w); err != nil {
		log.Printf("error rendering questions: %v", err)
	}
}

// topicNames はモデルの全トピックエリア名を定義順で返す。
func topicNames(m *model.DataMaturityModel) []string {
	names := make([]string, len(m.TopicAreas))
	for i, ta := range m.TopicAreas {
		names[i] = ta.Topic
	}
	return names
}

// buildTopicGroups はアセスメントリストをトピックエリアの定義順でグループ化する。
// レベル選択肢のソートおよびレベル名の解決も行う。
func buildTopicGroups(m *model.DataMaturityModel, levelNames map[int]string, assessments []model.Assessment) []components.TopicGroup {

	// トピック → []AssessmentOption のマップを作成
	groupMap := make(map[string][]components.AssessmentOption)
	for _, a := range assessments {
		keys := make([]int, 0, len(a.Levels))
		for k := range a.Levels {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		levels := make([]components.LevelOption, 0, len(keys))
		for _, k := range keys {
			levels = append(levels, components.LevelOption{
				Key:  k,
				Name: levelNames[k],
				Text: a.LevelsHTML[k],
			})
		}

		groupMap[a.Topic] = append(groupMap[a.Topic], components.AssessmentOption{
			ID:      a.ID,
			Summary: a.SummaryHTML,
			Levels:  levels,
		})
	}

	// topicAreas の定義順を維持してグループを構築
	seen := make(map[string]bool)
	groups := make([]components.TopicGroup, 0, len(m.TopicAreas))
	for _, ta := range m.TopicAreas {
		if seen[ta.Topic] {
			continue
		}
		seen[ta.Topic] = true
		if opts, ok := groupMap[ta.Topic]; ok {
			groups = append(groups, components.TopicGroup{
				Topic:   ta.Topic,
				Options: opts,
			})
		}
	}
	return groups
}
