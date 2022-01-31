package form

import (
	"fmt"

	"github.com/desertbit/fillpdf"
	"github.com/shopspring/decimal"
)

const (
	pageIndex     = 1
	BackPageIndex = 2
)

var irsReportingFieldMap = map[IRSReporting]int{
	ReportedAndBasisReported:        0,
	ReportedAndBasisDoesNotReported: 1,
	NotReported:                     2,
}

func ConvertReportToPDFForm(report Report) fillpdf.Form {
	form := make(fillpdf.Form)

	// fill base form for the front page #1
	for key, value := range convertPage(report.Name, report.SocialNumber, pageIndex, report.Front) {
		form[key] = value
	}

	// fill base form for the back page #2
	for key, value := range convertPage(report.Name, report.SocialNumber, BackPageIndex, report.Back) {
		form[key] = value
	}

	return form
}

func convertPage(name, socialNumber string, pageIndex int, page *Page) fillpdf.Form {
	form := make(fillpdf.Form)

	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_1[0]", pageIndex, pageIndex)] = name
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_2[0]", pageIndex, pageIndex)] = socialNumber

	if page == nil {
		return form
	}

	form[fmt.Sprintf("topmostSubform[0].Page%d[0].c%d_1[%d]", pageIndex, pageIndex, irsReportingFieldMap[page.IRSReporting])] = irsReportingFieldMap[page.IRSReporting] + 1

	totalProceeds := decimal.Zero
	totalCostBasis := decimal.Zero
	totalAdjustments := decimal.Zero

	fieldID := 3 // Start unique ID of the field in the PDF template (see details in the PDF structure)
	for i, row := range page.Rows {
		gain := row.CostBasis.Sub(row.Proceeds)

		totalProceeds = totalProceeds.Add(row.Proceeds)
		totalCostBasis = totalCostBasis.Add(row.CostBasis)

		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+0)] = row.Description
		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+1)] = row.DateAcquired.Format("02/01/2006")
		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+2)] = row.DateSold.Format("02/01/2006")
		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+3)] = row.Proceeds.StringFixed(2)
		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+4)] = row.CostBasis.StringFixed(2)
		if row.InstructionCode != "" {
			form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+5)] = row.InstructionCode
		}
		if !row.AdjustmentAmount.IsZero() {
			gain = gain.Add(row.AdjustmentAmount)
			totalAdjustments = totalAdjustments.Add(row.AdjustmentAmount)
			form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+6)] = row.AdjustmentAmount.StringFixed(2)
		}

		form[fmt.Sprintf("topmostSubform[0].Page%d[0].Table_Line1[0].Row%d[0].f%d_%d[0]", pageIndex, i+1, pageIndex, fieldID+7)] = gain.StringFixed(2)

		fieldID += 8 // Next row field IDs
	}

	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_115[0]", pageIndex, pageIndex)] = totalProceeds.StringFixed(2)
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_116[0]", pageIndex, pageIndex)] = totalCostBasis.StringFixed(2)
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_118[0]", pageIndex, pageIndex)] = totalAdjustments.StringFixed(2)
	form[fmt.Sprintf("topmostSubform[0].Page%d[0].f%d_119[0]", pageIndex, pageIndex)] = totalCostBasis.Sub(totalProceeds).Add(totalAdjustments).StringFixed(2)

	return form
}
