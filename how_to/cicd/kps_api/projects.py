import requests
import urllib3
import json

# remove some of the noise
requests.packages.urllib3.disable_warnings()
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# performs a simple get request against the Nutanix API to get a list of all projects
def get_projects(kps_token, url):

  api_url = ("%s/v1.0/projects" % (url))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers)
  j = json.loads(r.text)

  return j
