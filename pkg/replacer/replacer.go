package replacer

import (
	"fmt"
	"math"
	"strconv"

	"github.com/slavamuravey/wbppc/pkg/util"

	"github.com/xuri/excelize/v2"
)

func Replace() {
	f, err := excelize.OpenFile("b.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet := "Общий отчет"
	data, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}

	index, err := createSheetIndex("a.xlsx", "Отчёт по скидкам для акции", "f")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, row := range data {
		vendorCodeCol, err := excelize.ColumnNameToNumber("e")
		if err != nil {
			fmt.Println(err)
			return
		}

		if vendorCodeCol > len(row) || vendorCodeCol <= 0 {
			fmt.Println(err)
			return
		}

		vendorCode := row[vendorCodeCol-1]

		priceWithoutDiscountCol, err := excelize.ColumnNameToNumber("l")
		if err != nil {
			fmt.Println(err)
			return
		}

		if priceWithoutDiscountCol > len(row) || priceWithoutDiscountCol <= 0 {
			fmt.Println(err)
			return
		}

		priceWithoutDiscountStr := row[priceWithoutDiscountCol-1]
		priceWithoutDiscount, err := util.NumberFromString(priceWithoutDiscountStr)
		if err != nil {
			continue
		}

		v, ok := index[vendorCode]
		if !ok {
			continue
		}

		priceWithDiscountCol, err := excelize.ColumnNameToNumber("d")
		if err != nil {
			fmt.Println(err)
			return
		}

		if priceWithDiscountCol > len(v) || priceWithDiscountCol <= 0 {
			continue
		}

		priceWithDiscountStr := v[priceWithDiscountCol-1]

		priceWithDiscount, err := util.NumberFromString(priceWithDiscountStr)
		if err != nil {
			continue
		}

		discount := CalcDiscountPercent(priceWithoutDiscount, priceWithDiscount)
		if err != nil {
			fmt.Println(err)
			return
		}

		f.SetCellValue(sheet, "p"+strconv.Itoa(i+1), discount)
	}
	f.Save()
}

func CalcDiscountPercent(fromPrice float64, toPrice float64) int {
	return int(math.Ceil((1 - toPrice/fromPrice) * 100))
}

func createSheetIndex(filename string, sheet string, indexColumnName string) (map[string][]string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	indexColumnNumber, err := excelize.ColumnNameToNumber(indexColumnName)
	if err != nil {
		return nil, err
	}

	return util.CreateIndex(data, indexColumnNumber-1), nil
}
