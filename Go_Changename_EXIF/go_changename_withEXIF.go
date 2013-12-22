package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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

func Inject_EXIF(Input_File string, TypeCount int) {
	//06651_Scenario_Abroad_Abroad1_L-35.476328N-139.533234E_D-201106171900_017.jpg
	File_Tags := strings.Split(filepath.Base(Input_File), "_")
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
	fmt.Println(TypeCount)
	ModifyDate := OrgDate.Add(2 * time.Minute)
	for i := 1; i <= TypeCount; i++ {
		ModifyDate = ModifyDate.Add(2 * time.Minute)
	}

	//fmt.Println(ModifyDate)
	Create_EXIF_CSV(Input_File, ModifyDate.String(), FloatToString(FLa), FloatToString(FLo), LaD, LoD)
	//output, err := exec.Command("cmd.exe", "/c", "exiftool", "-k").Output()
	output, err := exec.Command("cmd.exe", "/c", "exiftool", "-exif:all=", "-csv=tmp_exif.csv", Input_File).Output()
	check(err)
	fmt.Println(string(output))
	//del_file := Input_File + "_orig"
	_, err = exec.Command("cmd.exe", "/c", "rm", "-f", Input_File+"_original").Output()
	check(err)
}

func main() {
	FileCount := 0
	TotalCount := 0
	//OrgPath := os.Args[1]
	// "c:/HTC/test2/"
	//fmt.Println("input : ", OrgPath)
	OrgPath := "c:/HTC/test2/"
	Recurse := true
	var newpath string
	var TypeName []string
	var OldTypeName []string
	TypeCount := 0
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)

		replace_strings := strings.Replace(OrgPath, "/", "_", -1)
		if err != nil {
			return err
		}

		if stat.IsDir() && path != OrgPath && !Recurse {
			fmt.Println("skipping dir:", path)
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}
		if stat.IsDir() {
			//fmt.Println(path)
			FileCount = 0
			newpath = strings.Replace(path, "\\", "_", -1)
			newpath = strings.Replace(newpath, replace_strings, "", -1)
			TypeName = strings.Split(newpath, "_")
			if strings.Contains(newpath, "Scenario") && len(TypeName) > 4 {
				if len(OldTypeName) < 4 {
					OldTypeName = TypeName
				}
				if !stringInSlice(OldTypeName[3], TypeName) {
					TypeCount = 0
				}
				OldTypeName = TypeName

			}
			//fmt.Println(newpath)
		}
		if !stat.IsDir() {
			if strings.Contains(path, "jpg") {
				TotalCount += 1
				FileCount += 1

				newfilename := fmt.Sprintf("%05d_%s_%03d", TotalCount, newpath, FileCount)
				//fmt.Println(newfilename)
				fmt.Println("org file name : ", path)
				//fmt.Println(filepath.Base(path))
				//fmt.Println(filepath.Dir(path))
				newfullname := filepath.Dir(path) + "\\" + newfilename + ".jpg"
				fmt.Println("new file name : ", newfullname)
				os.Rename(path, newfullname)

				if strings.Contains(path, "Scenario") {
					TypeCount += 1

					Inject_EXIF(newfullname, TypeCount)
				}
			}

		}

		return nil
	}
	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {
		log.Fatal(err)
	}

}
