package entities

import (
	"gorm.io/gorm"
	"time"
)

type ReportingTransaction struct {
	gorm.Model
	TransDate      time.Time
	ModifiedAt     time.Time
	OrderID        string
	Currency       string
	Rrn            string
	Amount         float64
	Balance        float64
	Description    string
	TransFee       float64
	AccountNo      string
	ProfileID      string
	ProfileName    string
	RecordType     string
	ProviderRef    string
	TransactionRef string
	ServiceType    string
	Provider       string
	Status         string
	MinuteAudited  bool `gorm:"default:false"`
	HourlyAudited  bool `gorm:"default:false"`
	DailyAudited   bool `gorm:"default:false"`
}
