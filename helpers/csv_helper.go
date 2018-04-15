package helpers

import (
	"encoding/csv"
	"os"
)

func ToCsv(path string, records []interface{}) error {

	// O_WRONLY:書き込みモード開く, O_CREATE:無かったらファイルを作成
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0) // ファイルを空っぽにする(2回目以降用)
	if err != nil {
		return err
	}

	ary := make([][]string, len(records))
	writer := csv.NewWriter(file)
	writer.WriteAll(ary)
	writer.Flush()

	return nil
}
