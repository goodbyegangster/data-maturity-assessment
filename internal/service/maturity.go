package service

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"data-maturity-assessment/internal/markdown"
	"data-maturity-assessment/internal/model"

	"gopkg.in/yaml.v3"
)

// MaturityService はデータ成熟度評価モデルのビジネスロジックを提供する。
// データは起動時にオンメモリにロードされ、以降は読み取り専用で使用する。
type MaturityService struct {
	// models は model.Model をキーとしたデータ成熟度評価モデルのマップ
	models map[string]*model.DataMaturityModel
	// orderedModels は ListModels で返す順序を保持するスライス
	orderedModels []*model.DataMaturityModel
}

// TopicScore はトピックエリアごとのスコアを表す。
type TopicScore struct {
	Topic    string
	Score    int  // 選択されていない場合は 0
	Selected bool // トピックエリアが選択されているかどうか
}

// NewMaturityService は dataDir 内の YAML ファイルをすべて読み込み、MaturityService を構築する。
func NewMaturityService(dataDir string) (*MaturityService, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read data directory %q: %w", dataDir, err)
	}

	svc := &MaturityService{
		models: make(map[string]*model.DataMaturityModel),
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		path := filepath.Join(dataDir, entry.Name())
		m, err := loadModel(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load model from %q: %w", path, err)
		}

		svc.models[m.Model] = m
		svc.orderedModels = append(svc.orderedModels, m)
	}

	sort.Slice(svc.orderedModels, func(i, j int) bool {
		return svc.orderedModels[i].ModelId < svc.orderedModels[j].ModelId
	})

	return svc, nil
}

// loadModel は指定パスの YAML ファイルを読み込み、DataMaturityModel を返す。
// ロード後にアセスメントの Markdown を HTML に事前変換する。
func loadModel(path string) (*model.DataMaturityModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var m model.DataMaturityModel
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	for i := range m.Assessments {
		a := &m.Assessments[i]
		a.SummaryHTML = markdown.ToHTML(a.Summary)
		a.LevelsHTML = make(map[int]string, len(a.Levels))
		for k, v := range a.Levels {
			a.LevelsHTML[k] = markdown.ToHTML(v)
		}
	}

	return &m, nil
}

// GetModel は指定されたモデルキーに対応するモデルを返す。
// 該当するモデルが存在しない場合は false を返す。
func (s *MaturityService) GetModel(modelKey string) (*model.DataMaturityModel, bool) {
	m, ok := s.models[modelKey]
	return m, ok
}

// ListModels は利用可能なモデルの一覧をファイルロード順で返す。
func (s *MaturityService) ListModels() []*model.DataMaturityModel {
	return s.orderedModels
}

// LevelNames はモデルの成熟度レベル名のルックアップテーブルを返す。
func (s *MaturityService) LevelNames(m *model.DataMaturityModel) map[int]string {
	names := make(map[int]string, len(m.MaturityLevels))
	for _, ml := range m.MaturityLevels {
		names[ml.Level] = ml.Name
	}
	return names
}

// toTopicSet は文字列スライスを存在チェック用のセットに変換する。
func toTopicSet(topics []string) map[string]struct{} {
	set := make(map[string]struct{}, len(topics))
	for _, t := range topics {
		set[t] = struct{}{}
	}
	return set
}

// FilterAssessments は指定されたトピック名のリストに一致するアセスメントを、
// assessments の定義順を維持したまま返す。
func (s *MaturityService) FilterAssessments(m *model.DataMaturityModel, selectedTopics []string) []model.Assessment {
	topicSet := toTopicSet(selectedTopics)

	var result []model.Assessment
	for _, a := range m.Assessments {
		if _, ok := topicSet[a.Topic]; ok {
			result = append(result, a)
		}
	}
	return result
}

// CalculateScores は選択されたトピックエリアごとのスコアを topicAreas の定義順で返す。
//
// スコアの計算:
//   - 対象トピックに属する回答済みアセスメントのレベル値（整数キー）を合算し、回答数で除算（小数点以下切り捨て）
//
// 未選択のトピックエリアも Score=0, Selected=false として結果に含む。
func (s *MaturityService) CalculateScores(m *model.DataMaturityModel, selectedTopics []string, answers map[int]int) []TopicScore {
	selectedSet := toTopicSet(selectedTopics)

	// トピックごとの回答レベル値を集計
	topicLevels := make(map[string][]int)
	for _, a := range m.Assessments {
		if _, ok := selectedSet[a.Topic]; !ok {
			continue
		}
		if level, answered := answers[a.ID]; answered {
			topicLevels[a.Topic] = append(topicLevels[a.Topic], level)
		}
	}

	// topicAreas の定義順でスコアを構築
	scores := make([]TopicScore, 0, len(m.TopicAreas))
	for _, ta := range m.TopicAreas {
		_, selected := selectedSet[ta.Topic]
		ts := TopicScore{
			Topic:    ta.Topic,
			Selected: selected,
		}
		if selected {
			levels := topicLevels[ta.Topic]
			if len(levels) > 0 {
				sum := 0
				for _, l := range levels {
					sum += l
				}
				ts.Score = sum / len(levels) // 小数点以下切り捨て
			}
		}
		scores = append(scores, ts)
	}

	return scores
}
