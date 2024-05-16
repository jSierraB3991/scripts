package libs

func GetStringPoint(data string) *string {
	if data == "" {
		return nil
	}
	return &data
}
