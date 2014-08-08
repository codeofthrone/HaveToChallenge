package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	r.ParseForm()
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println("vendor", r.FormValue("Vendor"))
	// SQLite precondition

	// Switch Filter Insert Here

	// Make Current Value

	// Store in SQLite

	// Generate Page
	vendor := r.FormValue("Vendor")
	if r.Method == "POST" {
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

	//p := &Page{ Vendor: "Vendor", Product: "Product", Model: "Model",    OS: "OS",    Version: "Version",    SDCardSize: "SDCardSize",    SIMCardSize: "SIMCardSize",    CPU: "CPU",    GPU: "GPU",    RAM: "RAM",    StorageSize: "StorageSize",    Wifi: "Wifi",    BlueTooth: "BlueTooth",    NFC: "NFC",    ScreenSize: "ScreenSize",    FrontCamera: "FrontCamera",    RearCamera: "RearCamera",    FlashLight: "FlashLight",    DPI: "DPI",    Resolution: "Resolution",    Gyroscope: "Gyroscope",    Gsensor: "Gsensor",    Barometer: "Barometer",    WifiCharge: "WifiCharge",    Comment: "Comment" }
}

func main() {
	http.HandleFunc("/New/", NewHandler)
	err := http.ListenAndServe(":22629", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
