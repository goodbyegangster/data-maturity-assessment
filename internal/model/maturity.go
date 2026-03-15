package model

// MaturityLevel はデータ成熟度の各レベルを表す。
// 英国・豪州・独自モデル共通の構造であり、YAML の maturityLevels[] に対応する。
type MaturityLevel struct {
	Level       int    `yaml:"level"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// ThemeArea はデータ成熟度の評価テーマを表す。
// YAML の themeAreas[] に対応する。英国モデル固有フィールド。
type ThemeArea struct {
	Theme       string `yaml:"theme"`
	Description string `yaml:"description"`
}

// TopicArea は評価カテゴリ（トピックエリア）を表す。
// YAML の topicAreas[] に対応する。
//
// 英国モデル固有フィールド: Summary, Concept, ScopeTheme
// 豪州・独自モデル固有フィールド: Description
type TopicArea struct {
	Topic string `yaml:"topic"`

	// 英国モデル固有
	Summary    string   `yaml:"summary,omitempty"`
	Concept    string   `yaml:"concept,omitempty"`
	ScopeTheme []string `yaml:"scopeTheme,omitempty"`

	// 豪州・独自モデル固有
	Description string `yaml:"description,omitempty"`
}

// Assessment は個々のアセスメント質問を表す。
// YAML の assessments[] に対応する。
//
// 英国モデル固有フィールド: Theme
type Assessment struct {
	ID      int            `yaml:"id"`
	Topic   string         `yaml:"topic"`
	Summary string         `yaml:"summary"`
	Levels  map[int]string `yaml:"levels"`

	// 英国モデル固有
	Theme string `yaml:"theme,omitempty"`

	// サービス初期化時に Markdown → HTML 変換済みの値を格納する（YAML には非対応）
	SummaryHTML string            `yaml:"-"`
	LevelsHTML  map[int]string    `yaml:"-"`
}

// DataMaturityModel はデータ成熟度評価モデルの YAML ファイル全体を表す。
// 英国・豪州・独自いずれの YAML ファイルもこの構造体でパース可能である。
type DataMaturityModel struct {
	Title            string          `yaml:"title"`
	Model            string          `yaml:"model"`
	ModelId          int             `yaml:"modelId"`
	ModelDescription string          `yaml:"modelDescription"`
	MaturityLevels   []MaturityLevel `yaml:"maturityLevels"`
	ThemeAreas       []ThemeArea     `yaml:"themeAreas,omitempty"` // 英国モデル固有フィールド
	TopicAreas       []TopicArea     `yaml:"topicAreas"`
	Assessments      []Assessment    `yaml:"assessments"`
}
