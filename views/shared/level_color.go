package shared

// LevelBgClass は成熟度レベルに応じた Tailwind 背景色クラスを返す。
// 完全なクラス名として定義することで Tailwind CSS スキャンに対応する。
func LevelBgClass(level int) string {
	switch level {
	case 0:
		return "bg-red-50"
	case 1:
		return "bg-orange-50"
	case 2:
		return "bg-yellow-50"
	case 3:
		return "bg-lime-50"
	case 4:
		return "bg-green-50"
	case 5:
		return "bg-blue-50"
	default:
		return "bg-gray-50"
	}
}

// LevelTextClass は成熟度レベルに応じた Tailwind テキスト色クラスを返す。
func LevelTextClass(level int) string {
	switch level {
	case 0:
		return "text-red-700"
	case 1:
		return "text-orange-700"
	case 2:
		return "text-yellow-700"
	case 3:
		return "text-lime-700"
	case 4:
		return "text-green-700"
	case 5:
		return "text-blue-700"
	default:
		return "text-gray-600"
	}
}

// LevelColorRGB は成熟度レベルに応じた CSS rgb() カラー値を返す。
// レーダーチャートの点色など、Tailwind クラス以外で色を指定する場合に使用する。
func LevelColorRGB(level int) string {
	switch level {
	case 0:
		return "rgb(239,68,68)" // red-500
	case 1:
		return "rgb(249,115,22)" // orange-500
	case 2:
		return "rgb(234,179,8)" // yellow-500
	case 3:
		return "rgb(132,204,22)" // lime-500
	case 4:
		return "rgb(34,197,94)" // green-500
	case 5:
		return "rgb(59,130,246)" // blue-500
	default:
		return "rgb(156,163,175)" // gray-400
	}
}

// LevelInteractiveClass は成熟度レベルに応じた hover / has-checked を含む Tailwind クラスを返す。
// 測定画面のラジオボタン選択肢に使用する。
func LevelInteractiveClass(level int) string {
	switch level {
	case 0:
		return "hover:bg-red-50 has-checked:bg-red-50 has-checked:ring-1 has-checked:ring-red-200"
	case 1:
		return "hover:bg-orange-50 has-checked:bg-orange-50 has-checked:ring-1 has-checked:ring-orange-200"
	case 2:
		return "hover:bg-yellow-50 has-checked:bg-yellow-50 has-checked:ring-1 has-checked:ring-yellow-200"
	case 3:
		return "hover:bg-lime-50 has-checked:bg-lime-50 has-checked:ring-1 has-checked:ring-lime-200"
	case 4:
		return "hover:bg-green-50 has-checked:bg-green-50 has-checked:ring-1 has-checked:ring-green-200"
	case 5:
		return "hover:bg-blue-50 has-checked:bg-blue-50 has-checked:ring-1 has-checked:ring-blue-200"
	default:
		return "hover:bg-gray-50 has-checked:bg-gray-100 has-checked:ring-1 has-checked:ring-gray-300"
	}
}
