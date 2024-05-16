package mapper

import "log"

func getDataFfm(rows []string) interface{} {
	err := validateVoidData(rows, []int{5, 8, 16, 17, 18, 19, 21})
	if err != nil {
		log.Panic(err)
	}
	// 15 string
	log.Println(rows[10])
	return nil
}
