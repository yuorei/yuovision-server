package lib

// stringを*stringに変換する
func StringPointer(s string) *string {
	return &s
}
