package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type KvHighFrequency struct {
	DataNo            int       `gorm:"primaryKey;autoIncrement:false"`
	DataDate          time.Time `gorm:"primaryKey;autoIncrement:false;not null;type:date"`
	DataTime          string    `gorm:"primaryKey;autoIncrement:false"`
	Interval          int
	CurrentNow        int
	TuningNow         int
	D00108            int
	TempUp            int
	TempDown          int
	WeldingTime       int
	CoolingTime       int
	TuningSpeed       int
	CurrentValue1     int
	TuningValue1      int
	WeldingTime1      int
	CoolingTime1      int
	WeldingTime2      int
	TuningValue2      int
	CurrentValue2     int
	WeldingSW2        int
	SetTempUp         int
	HighVoltageMemory int
	SetTempDown       int
}
type KvHighFrequencyResult struct {
	DataDate     string
	CurrentNow   float64
	TuningNow    float64
	WeldingTime  float64
	CoolingTime  float64
	WeldingTime1 float64
	CoolingTime1 float64
	WeldingTime2 float64
}

func (d KvHighFrequency) GetKvHighFrequency(config *Config, selDate string) (highFrequency []*KvHighFrequencyResult, err error) {

	db, _ := gorm.Open(mysql.New(config.SetDbConfig()), &gorm.Config{})

	db.AutoMigrate(&KvHighFrequency{})
	strSelect := "CONCAT(SUBSTRING(data_time,1,2),'時')  as data_date, avg(current_now) as current_now,avg(tuning_now) as tuning_now"
	strSelect += ",avg(welding_time) as welding_time,avg(cooling_time) as cooling_time"
	strSelect += ",avg(welding_time1) as welding_time1,avg(cooling_time1) as cooling_time1"
	strSelect += ",avg(welding_time2) as welding_time2"
	rows, err := db.Table("kv_high_frequencies").Select(strSelect).Where("Data_Date = ?", selDate).Group("SUBSTRING(data_time,1,2)").Rows()

	if err != nil {
		return highFrequency, err
	}
	for rows.Next() {
		var newRow KvHighFrequencyResult
		db.ScanRows(rows, &newRow)
		highFrequency = append(highFrequency, &newRow)
	}
	defer rows.Close()

	return highFrequency, err
}
