package writer

import (
	"fmt"
	v1 "main/v1"

	"github.com/xuri/excelize/v2"
)

func WriteAnswer(path string, ans *v1.Ans) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	for i := 0; i < len(ans.Path); i++ {
		s := ans.Path[i].ToStrArr()
		err := f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+1), &s)
		if err != nil {
			return err
		}
	}

	err := f.SaveAs(path)

	return err
}
