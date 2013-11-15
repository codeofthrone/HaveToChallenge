#!/bin/bash

##
## exiftool  -GPSLongitude="7.422809"  -GPSLatitude="48.419973" IMAG0083.jpg 
##

EXIFTools=`which exiftool`

Source_Path=$1


cd $Source_Path
File_List=`ls -l $Source_Path |grep 'jpg' | awk '{print $9}'`
for file_name in $File_List
do
	echo $file_name
    cp /tmp/tmp_exif.csv /tmp/tmp_exif_bak.csv
    sed -i -e "s,EXIF_TMP.jpg,$file_name," /tmp/tmp_exif_bak.csv
	$EXIFTools -exif:all= -csv=/tmp/tmp_exif_bak.csv $file_name
    rm -f $file_name"_original"
    

done 

