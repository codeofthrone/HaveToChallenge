#!/usr/bin/python

import sys,re,os
from os import listdir
from os.path import isfile, join
import time,itertools

tStart = time.time()


print " =============== result =============== "
print " execute file : " + sys.argv[0]
print " Arg 1 : " + sys.argv[1]
InputArg = sys.argv[1]
#onlyfiles = [ f for f in listdir(mypath) if isfile(join(mypath,f)) ]
#print onlyfiles

output = open(os.path.abspath('.')+'/output.txt','w')
firstoutput = open(os.path.abspath('.')+'/firstoutput.txt','w')
finaloutput = open(os.path.abspath('.')+'/finaloutput.txt','w')
#file.write(str(list_cn)+'\n')  
#for item in list_cn:  
#    print >> file,'%s\n' % item  
  

filelist=[]
lastElapsedTime = 0
lastline=[]

file=open(InputArg)
for line in file:
    #print line
    buf = []
    if 'jpg' in line :
        #print lastline
        for item in lastline:
            finaloutput.write("%s " % item )
        finaloutput.write("\n")
 
        lastline=[]
        filename = line.strip('\n')
        filelist.append(filename)
        lastElapsedTime = 0
        #print line + " start !"
    elif 'Done' in line :
        output.write("\n")
        pass
        #print line + " end ! "
    else : 
        l=line.split(' ')
        [float(item) for item in l ]
        l[6]=l[6].strip('\n')
        l.insert(0,filename)
        Step = l[1]
        PointCOunt = l[2]
        AreaB = float(l[4])
        AreaC = float(l[3])
        AreaE = float(l[6])
        AreaF = float(l[5])
        AreaA =  float(AreaC) +float(AreaB) 
        AreaD =  float(AreaF) +float(AreaE) 
        FormulaA = format(((AreaF+(AreaB-AreaE))/AreaA)*100,'.5f' )
        FormulaB = format(((AreaF-AreaE)/AreaC)*100 , '.5f')
        FormulaC = format( (AreaF/AreaC)*100 , '.5f')
        ElapsedTimeDifference = float(l[7]) - float(lastElapsedTime)
        lastElapsedTime = float(l[7])
        buf.extend(l)
        buf.insert(7, FormulaA)
        buf.insert(8, FormulaB)
        buf.insert(9, FormulaC)
        buf.insert(11, format(ElapsedTimeDifference,'.5f'))
        lastline = buf
        
        if int(Step) == 1 :
            for item in buf:
                firstoutput.write("%s " % item.strip("\n") )
            firstoutput.write("\n")


        for item in buf:
            output.write("%s " % item )
        output.write("\n")
        #print buf
        #print buf        
        #output.write(l[0]+" "+l[1]+" "+l[2]+" "+l[3]+" "+l[4]+" "+l[5]+" "+l[6]+" "+)
#       output.write(str( float(l[2])+float(l[3])))
        
for item in lastline:
    finaloutput.write("%s " % item )
finaloutput.write("\n")
    


tStop = time.time()

print "total process time " ,tStop - tStart
output.close()  
finaloutput.close()  
