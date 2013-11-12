#!/usr/bin/python


import flickr_api
import urllib, urlparse
import os
import sys



API_KEY = 'a74076cd639a153906f93a086e9b4bc9'
API_SECRET = 'bedb151aeb6c762c'

flickr_api.set_keys(api_key = API_KEY , api_secret = API_SECRET)





if len(sys.argv)>1:
    tag = sys.argv[1]
else:
    print 'no tag specified'

# downloading image data
f = flickr.photos_search(tags=tag)
urllist = [] #store a list of what was downloaded

# downloading images
for k in f:
    url = k.getURL(size='Original', urlType='source')
    urllist.append(url) 
    image = urllib.URLopener()
    image.retrieve(url, os.path.basename(urlparse.urlparse(url).path)) 
    print 'downloading:', url

'''
If you also want to write the list of urls to a text file, add the following lines at the end.
'''

# write the list of urls to file       
fl = open('urllist.txt', 'w')
for url in urllist:
    fl.write(url+'\n')
fl.close()
