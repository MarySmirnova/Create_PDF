package form

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	ReportedAndBasisReported        IRSReporting = "A"
	ReportedAndBasisDoesNotReported IRSReporting = "B"
	NotReported                     IRSReporting = "C"
)

type Report struct {
	Name         string `json:"name"`
	SocialNumber string `json:"social_number"`
	Front        *Page  `json:"front,omitempty"`
	Back         *Page  `json:"back,omitempty"`
}

type IRSReporting string

type Page struct {
	IRSReporting IRSReporting `json:"irs_reporting"`
	Rows         []Row
}

type Row struct {
	Description      string          `json:"description"`
	DateAcquired     time.Time       `json:"date_acquired"`
	DateSold         time.Time       `json:"date_sold"`
	Proceeds         decimal.Decimal `json:"proceeds"`
	CostBasis        decimal.Decimal `json:"cost_basis"`
	InstructionCode  string          `json:"instruction_code"`
	AdjustmentAmount decimal.Decimal `json:"adjustment_amount"`
}
