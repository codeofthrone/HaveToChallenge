package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Key   string
	Title string
}

type New struct {
	Vendor      string
	Product     string
	Model       string
	OS          string
	Version     string
	SDCardSize  string
	SIMCardSize string
	CPU         string
	GPU         string
	RAM         string
	StorageSize string
	Wifi        string
	BlueTooth   string
	NFC         string
	ScreenSize  string
	FrontCamera string
	RearCamera  string
	FlashLight  string
	DPI         string
	Resolution  string
	Gyroscope   string
	Gsensor     string
	Barometer   string
	WifiCharge  string
	Comment     string
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println("vendor", r.FormValue("Vendor"))
	// SQLite precondition

	// Switch Filter Insert Here

	// Make Current Value

	// Store in SQLite

	// Generate Page
	vendor := r.FormValue("Vendor")
	//	newpage:=&New{r.FormValue("Vendor"),r.FormValue("Product"), r.FormValue("Model"), r.FormValue("OS"), r.FormValue("Version"), r.FormValue("SDCardSize"), r.FormValue("SIMCardSize"), r.FormValue("CPU"), r.FormValue("GPU"),, r.FormValue("RAM"), r.FormValue("Storage"), r.FormValue("Wifi"), r.FormValue("BlueTooth"), r.FormValue("NFC"), r.FormValue("ScreenSize"), r.FormValue("FrontCamera"), r.FormValue("RearCamera"), r.FormValue("FlashLight"), r.FormValue("DPI"), r.FormValue("Resolution"), r.FormValue("Gyroscope"), r.FormValue("G_Sensor"), r.FormValue("Barometer"), r.FormValue("WifiCharge"), r.FormValue("Comment")}

	if r.Method == "POST" {
		//Write in DB
		db, _ := sql.Open("sqlite3", "./DeviceManagement.s3db")
		if r.FormValue("Vendor") != "" {
			stmt, _ := db.Prepare("INSERT INTO Type_Phone(Vendor,Product,Model,OS,Version,SDCardSize,SIMCardSize,CPU,GPU,RAM,StorageSize,DataPlan,Wifi,BlueTooth,NFC,ScreenSize,FrontCamera,RearCamera,FlashLight,DPI,Resolution,Gyroscope,Gsensor,Barometer,WifiCharge,Comment) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
			res, _ := stmt.Exec(r.FormValue("Vendor"), r.FormValue("Product"), r.FormValue("Model"), r.FormValue("OS"), r.FormValue("Version"), r.FormValue("SDCardSize"), r.FormValue("SIMCardSize"), r.FormValue("CPU"), r.FormValue("GPU"), r.FormValue("RAM"), r.FormValue("StorageSize"), r.FormValue("DataPlan"), r.FormValue("Wifi"), r.FormValue("BlueTooth"), r.FormValue("NFC"), r.FormValue("ScreenSize"), r.FormValue("FrontCamera"), r.FormValue("RearCamera"), r.FormValue("FlashLight"), r.FormValue("DPI"), r.FormValue("Resolution"), r.FormValue("Gyroscope"), r.FormValue("Gsensor"), r.FormValue("Barometer"), r.FormValue("WifiCharge"), r.FormValue("Comment"))
			id, _ := res.LastInsertId()
			fmt.Println(id)

		} else if r.FormValue("Country") != "" {
			stmt, _ := db.Prepare("INSERT INTO Type_SIM(Country,Telecom,Number,PinCode,SIMCardSize,Comment) values(?,?,?,?,?,?)")
			res, _ := stmt.Exec(r.FormValue("Country"), r.FormValue("Telecom"), r.FormValue("Number"), r.FormValue("PinCode"), r.FormValue("SIMCardSizes"), r.FormValue("Comments"))
			id, _ := res.LastInsertId()
			fmt.Println(id)
		} else {
			stmt, _ := db.Prepare("INSERT INTO Type_Account(Company,Service,Register,Account,Password,Accessory,Comment) values(?,?,?,?,?,?,?)")
			res, _ := stmt.Exec(r.FormValue("Company"), r.FormValue("Service"), r.FormValue("Register"), r.FormValue("Account"), r.FormValue("Password"), r.FormValue("Accessory"), r.FormValue("Commenta"))
			id, _ := res.LastInsertId()
			fmt.Println(id)
		}
		db.Close()

		fmt.Println("empty")
		p := &Page{Key: "empty"}
		t, _ := template.ParseFiles("New.html")
		t.Execute(w, p)

	} else {
		fmt.Println("not empty")
		p := &New{Vendor: vendor}
		t, _ := template.ParseFiles("New.html")
		t.Execute(w, p)
	}
	//p := &Page{ Vendor: "Vendor", Product: "Product", Model: "Model",	OS: "OS",	Version: "Version",    SDCardSize: "SDCardSize",    SIMCardSize: "SIMCardSize",    CPU: "CPU",    GPU: "GPU",    RAM: "RAM",    StorageSize: "StorageSize",    Wifi: "Wifi",    BlueTooth: "BlueTooth",    NFC: "NFC",    ScreenSize: "ScreenSize",    FrontCamera: "FrontCamera",    RearCamera: "RearCamera",    FlashLight: "FlashLight",    DPI: "DPI",    Resolution: "Resolution",    Gyroscope: "Gyroscope",    Gsensor: "Gsensor",    Barometer: "Barometer",    WifiCharge: "WifiCharge",    Comment: "Comment" }
}

func main() {
	http.HandleFunc("/New/", NewHandler)
	err := http.ListenAndServe(":22629", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
