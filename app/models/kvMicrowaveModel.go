package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type KvMicrowave struct {
	DataNo            int       `gorm:"primaryKey;autoIncrement:false"`
	DataDate          time.Time `gorm:"primaryKey;autoIncrement:false;not null;type:date"`
	DataTime          string    `gorm:"primaryKey;autoIncrement:false"`
	MwPin1            int
	MwPin2            int
	MwPin3            int
	MwPin4            int
	MwPf1             int
	MwPf2             int
	MwPf3             int
	MwPf4             int
	MwPr1             int
	MwPr2             int
	MwPr3             int
	MwPr4             int
	StopCV            int
	WaterLevel        int
	FarInfraredHeater int
	MwOutput          int
	MwTotalOutput     int
	TempRight         int
	TempLeft          int
	TempTop           int
	TempMirror        int
	TempFloor         int
	TempDoor          int
	InRoomPressure    int
}
type KvMicrowaveResult struct {
	DataDate          string
	StopCV            float64
	WaterLevel        float64
	FarInfraredHeater float64
	MwOutput          float64
	MwTotalOutput     float64
	InRoomPressure    float64
}

func (d KvMicrowave) GetKvMicrowave(config *Config, selDate string) (microWave []*KvMicrowaveResult, err error) {

	db, _ := gorm.Open(mysql.New(config.SetDbConfig()), &gorm.Config{})

	db.AutoMigrate(&KvMicrowave{})
	strSelect := "CONCAT(SUBSTRING(data_time,1,2),'時')  as data_date, avg(stop_cv) as stop_cv,avg(water_level) as water_level"
	strSelect += ",avg(far_infrared_heater) as far_infrared_heater,avg(mw_output) as mw_output"
	strSelect += ",avg(mw_total_output) as mw_total_output,avg(in_room_pressure) as in_room_pressure"
	rows, err := db.Table("kv_microwaves").Select(strSelect).Where("Data_Date = ?", selDate).Group("SUBSTRING(data_time,1,2)").Rows()
	if err != nil {
		return microWave, err
	}

	for rows.Next() {
		var newRow KvMicrowaveResult
		db.ScanRows(rows, &newRow)
		microWave = append(microWave, &newRow)
	}
	defer rows.Close()

	return microWave, err
}
