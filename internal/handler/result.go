package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"data-maturity-assessment/internal/service"
	"data-maturity-assessment/views/pages"
	"data-maturity-assessment/views/shared"
)

// ResultHandler は POST /result を処理する。
type ResultHandler struct {
	svc *service.MaturityService
}

// NewResultHandler は ResultHandler を構築する。
func NewResultHandler(svc *service.MaturityService) *ResultHandler {
	return &ResultHandler{svc: svc}
}

// Result は POST /result のハンドラ。
// フォームデータを受け取り、スコアを計算して結果表示画面を返す。
// バリデーションエラーの場合は /assessment にリダイレクトする。
func (h *ResultHandler) Result(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	modelKey := r.FormValue("model")
	m, ok := h.svc.GetModel(modelKey)
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	selectedTopics := r.Form["topicAreas"]
	if len(selectedTopics) == 0 {
		http.Redirect(w, r, "/assessment?model="+modelKey, http.StatusFound)
		return
	}

	// answers[{ID}] = levelKey 形式のフォーム値をパース
	answers := make(map[int]int)
	for key, values := range r.Form {
		if !strings.HasPrefix(key, "answers[") || !strings.HasSuffix(key, "]") {
			continue
		}
		idStr := key[len("answers[") : len(key)-1]
		id, err := strconv.Atoi(idStr)
		if err != nil || len(values) == 0 {
			continue
		}
		level, err := strconv.Atoi(values[0])
		if err != nil {
			continue
		}
		answers[id] = level
	}

	// 成熟度レベルの最小・最大を取得
	minLevel := m.MaturityLevels[0].Level
	maxLevel := m.MaturityLevels[len(m.MaturityLevels)-1].Level

	filteredAssessments := h.svc.FilterAssessments(m, selectedTopics)

	// 未回答のアセスメントは最小レベルで補完
	for _, a := range filteredAssessments {
		if _, ok := answers[a.ID]; !ok {
			answers[a.ID] = minLevel
		}
	}

	scores := h.svc.CalculateScores(m, selectedTopics, answers)

	// Chart.js 用 JSON を生成
	labels := make([]string, len(scores))
	scoreVals := make([]int, len(scores))
	selectedFlags := make([]bool, len(scores))
	for i, s := range scores {
		labels[i] = s.Topic
		if s.Selected {
			scoreVals[i] = s.Score
		} else {
			scoreVals[i] = minLevel // 未選択トピックはチャート軸最小値で表示
		}
		selectedFlags[i] = s.Selected
	}

	// 点色: 選択済みトピックはスコアレベルの配色、未選択は灰色
	const grayColor = "rgb(156,163,175)"
	pointColors := make([]string, len(scores))
	for i, s := range scores {
		if s.Selected {
			pointColors[i] = shared.LevelColorRGB(s.Score)
		} else {
			pointColors[i] = grayColor
		}
	}

	labelsJSON, err := json.Marshal(labels)
	if err != nil {
		log.Printf("error marshaling labels: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	scoresJSON, err := json.Marshal(scoreVals)
	if err != nil {
		log.Printf("error marshaling scores: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	selectedJSON, err := json.Marshal(selectedFlags)
	if err != nil {
		log.Printf("error marshaling selected flags: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	pointColorsJSON, err := json.Marshal(pointColors)
	if err != nil {
		log.Printf("error marshaling point colors: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// レベル名ルックアップ
	levelNames := h.svc.LevelNames(m)

	// アセスメント回答をスコアのトピック順でグループ化
	// answers マップはアセスメント ID → レベルキー
	answersByTopic := make(map[string][]pages.AssessmentResult)
	for _, a := range filteredAssessments {
		levelKey, answered := answers[a.ID]
		if !answered {
			continue
		}
		answersByTopic[a.Topic] = append(answersByTopic[a.Topic], pages.AssessmentResult{
			Summary:       a.SummaryHTML,
			SelectedLevel: levelKey,
			SelectedName:  levelNames[levelKey],
			SelectedText:  a.LevelsHTML[levelKey],
		})
	}

	topicGroups := make([]pages.ResultTopicGroup, 0, len(scores))
	for _, s := range scores {
		topicGroups = append(topicGroups, pages.ResultTopicGroup{
			Topic:    s.Topic,
			Score:    s.Score,
			Selected: s.Selected,
			Results:  answersByTopic[s.Topic],
		})
	}

	data := pages.ResultData{
		ModelTitle:      m.Title,
		LabelsJSON:      string(labelsJSON),
		ScoresJSON:      string(scoresJSON),
		SelectedJSON:    string(selectedJSON),
		PointColorsJSON: string(pointColorsJSON),
		MinLevel:        minLevel,
		MaxLevel:        maxLevel,
		TopicGroups:     topicGroups,
	}

	if err := pages.Result(data).Render(r.Context(), w); err != nil {
		log.Printf("error rendering result page: %v", err)
	}
}

