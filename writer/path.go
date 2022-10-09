package writer

import (
	"fmt"
	v1 "main/v1"

	"github.com/xuri/excelize/v2"
)

func WriteAnswerPath(path string, ans *v1.Ans) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	err := f.SetSheetRow("Sheet1", "A1", &[]string{"from", "to", "cap", "dis"})
	if err != nil {
		return err
	}

	for i := 0; i < len(ans.Path); i++ {
		s := ans.Path[i].ToStrArr()
		err := f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+2), &s)
		if err != nil {
			return err
		}
	}

	err = f.SaveAs(path)

	return err
}
