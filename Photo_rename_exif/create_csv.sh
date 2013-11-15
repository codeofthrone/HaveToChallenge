#!/bin/bash

## OrignalData  =  2013:11:13 13:07:18,

msg1="SourceFile,DateTimeOriginal,CreateDate,GPSLatitudeRef,GPSLongitudeRef,GPSAltitudeRef,GPSProcessingMethod,GPSAltitude,GPSLatitude,GPSLongitude,GPSPosition"
msg2="EXIF_TMP.jpg,$OrignalDate,$OrignalDate,North,East,Above Sea Level,ASCII,0 m Above Sea Level,\"24 deg 58' 45.22\"\" N\",\"121 deg 32' 44.71\"\" E\",\"24 deg 58' 45.22\"\" N, 121 deg 32' 44.71\"\" E"

echo $msg1 > temp.csv
echo $msg2 >> temp.csv
