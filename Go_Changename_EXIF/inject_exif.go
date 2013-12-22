package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func Create_EXIF_CSV(FileName, CreateDate, La, Lo, LaD, LoD string) {
	file, err := os.Create("tmp_exif.csv")
	check(err)

	defer file.Close()

	msg1 := "SourceFile,DateTimeOriginal,CreateDate,GPSLatitudeRef,GPSLongitudeRef,GPSAltitudeRef,GPSProcessingMethod,GPSAltitude,GPSLatitude,GPSLongitude,GPSPosition\n"
	msg2 := fmt.Sprintf("%s,%s,%s,North,East,Above Sea Level,ASCII,0 m Above Sea Level,\"%s\"\" %s\",\"%s\"\" %s\",\"%s\"\" %s\",\" %s\"\" %s\"", FileName, CreateDate, CreateDate, La, LaD, Lo, LoD, La, LaD, Lo, LoD)
	_, err = file.WriteString(msg1)
	_, err = file.WriteString(msg2)

	file.Sync()
}

func Inject_EXIF(Input_File string) {
	//06651_Scenario_Abroad_Abroad1_L-35.476328N-139.533234E_D-201106171900_017.jpg
	File_Tags := strings.Split(Input_File, "_")
	GPS := File_Tags[5]
	StrOrgDate := File_Tags[6]
	tempGPS := strings.Split(GPS, "-")
	Sla := tempGPS[1]
	Slo := tempGPS[2]
	LaD := Sla[len(Sla)-1:]
	LoD := Slo[len(Slo)-1:]
	Sla_num := Sla[:len(Sla)-2]
	Slo_num := Slo[:len(Slo)-2]
	//fmt.Println(Sla_num, LaD, Slo_num, LoD, StrOrgDate)

	GPSRandomA, _ := strconv.ParseFloat(strconv.Itoa(random(0, 6)), 32)
	FLa, _ := strconv.ParseFloat(Sla_num, 32)
	FLo, _ := strconv.ParseFloat(Slo_num, 32)
	FLa += GPSRandomA / 100000
	FLo += GPSRandomA / 100000

	fmt.Println(GPSRandomA, FLa, LaD, FLo, LoD)

	tempDate := strings.Split(StrOrgDate, "-")
	StrDate := tempDate[1]
	OrgDate, _ := time.Parse("200601021504", StrDate)
	ModifyDate := OrgDate.Add(2 * time.Minute).String()

	fmt.Println(ModifyDate)
	Create_EXIF_CSV(Input_File, ModifyDate, FloatToString(FLa), FloatToString(FLo), LaD, LoD)
	//output, err := exec.Command("cmd.exe", "/c", "exiftool", "-k").Output()
	output, err := exec.Command("cmd.exe", "/c", "exiftool", "-exif:all=", "-csv=tmp_exif.csv", Input_File).Output()
	check(err)
	fmt.Println(string(output))
	//del_file := Input_File + "_orig"
	_, err = exec.Command("cmd.exe", "/c", "rm", "-f", Input_File+"_original").Output()
	check(err)
}

func main() {
	//Input_File := os.Args[1]
	Input_File := "06651_Scenario_Family01_Abroad_Abroad1_L-35.476328N-139.533234E_D-201106171900_017.jpg"
	Inject_EXIF(Input_File)
}
