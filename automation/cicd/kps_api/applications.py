import requests
import urllib3
import json

# remove some of the noise
requests.packages.urllib3.disable_warnings()
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


# performs a simple get request against the Nutanix API to get a list of all the applications
def get_application(kps_token, url):
  api_url = ("%s/v1.0/applications" % (url))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  return j


def get_application_byprojectid(kps_token, url, projectid):
  api_url = ("%s/v1.0/projects/%s/applications" % (url, projectid))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  return j


# update an application in Xi IoT using the API and new parameters passed to the method
def update_application(kps_token, url, app_id, new_code, description):
  api_url = ("%s/v1.0/applications/%s" % (url, app_id))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  j["appManifest"] = new_code
  j["description"] = ("%s" % description)

  print("Update Body:")
  print(json.dumps(j, indent=2))


  api_url = ("%s/v1.0/applications/%s" % (url, app_id))
  headers = { 'Authorization' : 'Bearer ' + kps_token, 'Content-Type' : 'application/json' }

  r = requests.put(api_url, headers=headers, data=json.dumps(j), verify=False)
  j = json.loads(r.text)

  return j


# create a new application in Xi IoT using the API and parameters passed to the method
def create_application(kps_token, url, new_code, description, name, projectid):
  api_url = ("%s/v1.0/applications" % url)
  headers = { 'Authorization' : 'Bearer ' + kps_token, 'Content-Type' : 'application/json' }

  create_j = {
    "appManifest": new_code,
    "description": description,
    "name": name,
    "onlyPrePullOnUpdate": False,
    "projectId": projectid
  }
  print("Create Body:")
  print(json.dumps(create_j, indent=2))

  r = requests.post(api_url, headers=headers, data=json.dumps(create_j), verify=False)
  j = json.loads(r.text)

  return j


# delete application in Xi IoT based on the application ID
def delete_application(kps_token, url, app_id):
  api_url = ("%s/v1.0/applications/%s" % (url, app_id))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.delete(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  return j
