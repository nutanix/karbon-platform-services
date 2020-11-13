import requests
import urllib3
import json

# remove some of the noise
requests.packages.urllib3.disable_warnings()
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


# performs a simple get request against the Nutanix API to get a list of Nodes Info
def get_nodesinfo(kps_token, url):
  api_url = ("%s/v1.0/nodesinfo" % (url))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  return j


def get_nodesinfo_byprojectid(kps_token, url, projectid):
  api_url = ("%s/v1.0/projects/%s/nodesinfo" % (url, projectid))
  headers = { 'Authorization' : 'Bearer ' + kps_token }

  r = requests.get(api_url, headers=headers, verify=False)
  j = json.loads(r.text)

  return j
