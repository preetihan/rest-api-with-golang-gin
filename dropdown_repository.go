package main

import (
	"database/sql"
	"log"
	"strings"
)

var env = getEnv("ENV", Environment)

func connectDb() *sql.DB {
	con := getEnv("DB_CON", Dbcon)
	log.Println("DB_CON = " + con)

	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		log.Panicln("connectDb error: ", err)
	}

	return db
}

func getHcDropdown(path string) (dropdowns []DropdownRes, err error) {

	splitPath := strings.Split(path, "-")

	switch splitPath[2] {
	case "txnType":
		dropdowns = []DropdownRes{
			{Value: "1", Label: "1"},
			{Value: "2", Label: "2"},
			{Value: "3", Label: "3"},
			{Value: "4", Label: "4"},
			{Value: "5", Label: "5"},
			{Value: "6", Label: "6"},
		}
	case "hFrom":
		dropdowns = []DropdownRes{
			{Value: "00", Label: "00"},
			{Value: "01", Label: "01"},
			{Value: "02", Label: "02"},
			{Value: "03", Label: "03"},
			{Value: "04", Label: "04"},
			{Value: "05", Label: "05"},
			{Value: "06", Label: "06"},
			{Value: "07", Label: "07"},
			{Value: "08", Label: "08"},
			{Value: "09", Label: "09"},
			{Value: "10", Label: "10"},
			{Value: "11", Label: "11"},
			{Value: "12", Label: "12"},
			{Value: "13", Label: "13"},
			{Value: "14", Label: "14"},
			{Value: "15", Label: "15"},
			{Value: "16", Label: "16"},
			{Value: "17", Label: "17"},
			{Value: "18", Label: "18"},
			{Value: "19", Label: "19"},
			{Value: "20", Label: "20"},
			{Value: "21", Label: "21"},
			{Value: "22", Label: "22"},
			{Value: "23", Label: "23"},
		}
	case "mFrom":
		dropdowns = []DropdownRes{
			{Value: "00", Label: "00"},
			{Value: "01", Label: "01"},
			{Value: "02", Label: "02"},
			{Value: "03", Label: "03"},
			{Value: "04", Label: "04"},
			{Value: "05", Label: "05"},
			{Value: "06", Label: "06"},
			{Value: "07", Label: "07"},
			{Value: "08", Label: "08"},
			{Value: "09", Label: "09"},
			{Value: "10", Label: "10"},
			{Value: "11", Label: "11"},
			{Value: "12", Label: "12"},
			{Value: "13", Label: "13"},
			{Value: "14", Label: "14"},
			{Value: "15", Label: "15"},
			{Value: "16", Label: "16"},
			{Value: "17", Label: "17"},
			{Value: "18", Label: "18"},
			{Value: "19", Label: "19"},
			{Value: "20", Label: "20"},
			{Value: "21", Label: "21"},
			{Value: "22", Label: "22"},
			{Value: "23", Label: "23"},
			{Value: "24", Label: "24"},
			{Value: "25", Label: "25"},
			{Value: "26", Label: "26"},
			{Value: "27", Label: "27"},
			{Value: "28", Label: "28"},
			{Value: "29", Label: "29"},
			{Value: "30", Label: "30"},
			{Value: "31", Label: "31"},
			{Value: "32", Label: "32"},
			{Value: "33", Label: "33"},
			{Value: "34", Label: "34"},
			{Value: "35", Label: "35"},
			{Value: "36", Label: "36"},
			{Value: "37", Label: "37"},
			{Value: "38", Label: "38"},
			{Value: "39", Label: "39"},
			{Value: "40", Label: "40"},
			{Value: "41", Label: "41"},
			{Value: "42", Label: "42"},
			{Value: "43", Label: "43"},
			{Value: "44", Label: "44"},
			{Value: "45", Label: "45"},
			{Value: "46", Label: "46"},
			{Value: "47", Label: "47"},
			{Value: "48", Label: "48"},
			{Value: "49", Label: "49"},
			{Value: "50", Label: "50"},
			{Value: "51", Label: "51"},
			{Value: "52", Label: "52"},
			{Value: "53", Label: "53"},
			{Value: "54", Label: "54"},
			{Value: "55", Label: "55"},
			{Value: "56", Label: "56"},
			{Value: "57", Label: "57"},
			{Value: "58", Label: "58"},
			{Value: "59", Label: "59"},
		}
	case "sFrom":
		dropdowns = []DropdownRes{
			{Value: "00", Label: "00"},
			{Value: "01", Label: "01"},
			{Value: "02", Label: "02"},
			{Value: "03", Label: "03"},
			{Value: "04", Label: "04"},
			{Value: "05", Label: "05"},
			{Value: "06", Label: "06"},
			{Value: "07", Label: "07"},
			{Value: "08", Label: "08"},
			{Value: "09", Label: "09"},
			{Value: "10", Label: "10"},
			{Value: "11", Label: "11"},
			{Value: "12", Label: "12"},
			{Value: "13", Label: "13"},
			{Value: "14", Label: "14"},
			{Value: "15", Label: "15"},
			{Value: "16", Label: "16"},
			{Value: "17", Label: "17"},
			{Value: "18", Label: "18"},
			{Value: "19", Label: "19"},
			{Value: "20", Label: "20"},
			{Value: "21", Label: "21"},
			{Value: "22", Label: "22"},
			{Value: "23", Label: "23"},
			{Value: "24", Label: "24"},
			{Value: "25", Label: "25"},
			{Value: "26", Label: "26"},
			{Value: "27", Label: "27"},
			{Value: "28", Label: "28"},
			{Value: "29", Label: "29"},
			{Value: "30", Label: "30"},
			{Value: "31", Label: "31"},
			{Value: "32", Label: "32"},
			{Value: "33", Label: "33"},
			{Value: "34", Label: "34"},
			{Value: "35", Label: "35"},
			{Value: "36", Label: "36"},
			{Value: "37", Label: "37"},
			{Value: "38", Label: "38"},
			{Value: "39", Label: "39"},
			{Value: "40", Label: "40"},
			{Value: "41", Label: "41"},
			{Value: "42", Label: "42"},
			{Value: "43", Label: "43"},
			{Value: "44", Label: "44"},
			{Value: "45", Label: "45"},
			{Value: "46", Label: "46"},
			{Value: "47", Label: "47"},
			{Value: "48", Label: "48"},
			{Value: "49", Label: "49"},
			{Value: "50", Label: "50"},
			{Value: "51", Label: "51"},
			{Value: "52", Label: "52"},
			{Value: "53", Label: "53"},
			{Value: "54", Label: "54"},
			{Value: "55", Label: "55"},
			{Value: "56", Label: "56"},
			{Value: "57", Label: "57"},
			{Value: "58", Label: "58"},
			{Value: "59", Label: "59"},
		}
	default:
		dropdowns = []DropdownRes{}
	}

	return
}
