# Karbon Platform Services Python SDK

## Overview
The Karbon Platform Services (KPS) Python SDK.

# Steps to install KPS Python SDK

## PyPi

The Karbon Platform Services SDK can also be found on [PyPi](https://pypi.org/project/kps-api/)

```bash
pip install kps-api
```

## Manually
<pre>
tar -xvzf ./kps_api_1.0.<version>.tar.gz
pushd kps_api
python3 setup.py install
python3 -c "import kps_api"
retval=$?
if [[ "$retval" -ne "0" ]] ; then
  exit $retval
fi
popd
</pre>

## Manage API Keys

Prior to using the sdk, read through the [docs](https://portal.nutanix.com/page/documents/details?targetId=Karbon-Platform-Services-Project-User-Guide:ks-ks-manage-api-keys-c.html) on how to manage API keys in Karbon Platform Services.

## API Endpoints

The documentation for the Nutanix Karbon Platform Services API endpoints can be found on [Nutanix.dev](https://www.nutanix.dev/reference/karbon-platform-services/).

## Sample Code running steps, update KPS_API_KEY in following files before running.
<pre>
python3 samples/service_domain_client.py
python3 samples/application_client.py
</pre>

