package lib

// stringを*stringに変換する
func StringPointer(s string) *string {
	return &s
}

// Helper function to convert []string to []*string
func ConvertStringSliceToPointerSlice(slice []string) []*string {
	result := make([]*string, len(slice))
	for i, str := range slice {
		result[i] = &str
	}
	return result
}