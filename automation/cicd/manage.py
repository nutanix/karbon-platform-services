import json, warnings, errno, os, requests
import shutil, traceback, urllib, urllib3
import uuid, time, sys, re, getopt
from datetime import datetime
from kps_api import projects, applications

# Check for errors in API response 
def check_api_return(resp, expect_body):
	if "statusCode" in resp:
		print("API call failed with status code %s and error message:\n%s" % (resp["statusCode"], resp["message"]))
		sys.exit(1)
	# make sure there is data in the return body, only if result is set to one, otherwise we arent expecting a return body
	else:
		if expect_body:
			if "result" not in resp:
				print("Result data was empty")
				sys.exit(1) 

	return resp

# get the arguments passed to the script
url = ''
project = ''
token = ''
app = ''

try:
	opts, args = getopt.getopt(sys.argv[1:],"h:u:p:t:a:",['url=', 'project=', 'token=', 'app='])
except getopt.GetoptError:
	print ('main.py -u <xiiot url to connect to> -p <project name within xiiot> -t <token>')
	sys.exit(2)
for opt, arg in opts:
	if opt == '-h':
		print ('main.py -u <xiiot url to connect to> -p <project name within xiiot> -t <token>')
		sys.exit()
	elif opt in ('-u', '--url'):
		url = arg 
	elif opt in ('-p', '--project'):
		project = arg
	elif opt in ('-t', '--token'):
		token = arg
	elif opt in ('-a', '--app'):
		app = arg
	else:
		print ('manage.py -u <kps url to connect to> -p <project name within kps> -t <token>')
		sys.exit(2)

# get a list of all projects on the kps cluster
resp = projects.get_projects(token, url)
project_data = check_api_return(resp, expect_body=True)

# get project ID that matches our project name
for proj in project_data["result"]:
	if project == proj["name"]:
		project_id = proj["id"]

try:
	project_id
except NameError:
	print("Couldnt get ID for %s, make sure the Project exists in Karbon Platform Services" % project)
	sys.exit(1)

# create the list required to update/create an application in Xi IoT
local_appdata = {}
# read the deployment config, so we have all the code stored that needs to be compared/updated
f = open("deployment.yaml", "r")
new_code = f.read()
file_byline = f.readlines()
f.close()
local_appdata['appManifest'] = new_code
local_appdata['description'] = app
for line in file_byline:
  line = line.rstrip()
  if re.match("#\sdescription\:\s.+", line):
    local_appdata['description'] = line.split(': ')[1]

# loop through remote apps to check if it already exists
resp = applications.get_application_byprojectid(token, url, project_id)
application_data = check_api_return(resp, expect_body=True)

match = False
# loop through the existing applications for this project to see if it exists already
for curr_app in application_data["result"]:
  if curr_app["name"] == app:
    match = True
    app_id = curr_app["id"]

# check if app needs to be created or updated 
if match:
	print("%s currently exists in Karbon Platform Services for project %s" % (app, project))
	print("\nUpdating App %s in Karbon Platform Services" % app)
	app_update_data = applications.update_application(token, url, app_id, local_appdata['appManifest'], local_appdata['description'])
	print("Response Body:")
	print(json.dumps(app_update_data))
	if "statusCode" not in app_update_data:
		print("\nApp update successful")
	elif app_update_data["statusCode"] == 400:
		print("\nApp update failed with message: %s." % app_update_data["message"])
		sys.exit(1)
else:
	print("%s currently does not exist in Karbon Platform Services, creating it for project %s" % (app, project))
	resp = applications.create_application(token, url, local_appdata['appManifest'], local_appdata['description'], app, project_id)
	check_api_return(resp, expect_body=False)