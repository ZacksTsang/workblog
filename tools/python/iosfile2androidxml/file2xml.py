import sys
import os

toxml = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
toxml += "<resources>\n"

print(sys.argv[0])

if len(sys.argv) == 1:
	print("please type file and output file name")
elif len(sys.argv) == 2:
	print("please type output file name")
else:
	with open(sys.argv[1],'r') as f:
	    for line in f:
	    	if len(line)>0 and (line[0]=="\"" or line[0]=="/"):
	    		line=line.replace(";","").replace("\n","")
	    		if line[0] == "/":
	    			line = line.replace("/","")
	    			toxml += "	<!--"+line+"-->\n"
	    		else:
	    			lists = line.split("=")

	    			content = lists[1].replace("\"","")

	    			conlist = list(content)

	    			index = 1 
	    			for i, c in enumerate(conlist):
	    				if c == "@":
	    					conlist[i] = `index` + "$s" 
	    					index+=1
	    			content=''.join(conlist)

	    			toxml+="	<string name="+lists[0]+">"+content+"</string>\n"
	
	toxml += "</resources>"

	print(toxml)
	print("export to xml")
	f = open(sys.argv[2]+'.xml','w')
	f.write(toxml) # python will convert \n to os.linesep
	f.close()

	print("success to xml")