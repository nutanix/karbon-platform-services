# Karbon Platform Services Python SDK

## Overview
The Karbon Platform Services (KPS) Python SDK.

# Steps to install KPS Python SDK
<pre>
tar -xvzf ./kps_api_1.0.1219.tar.gz
pushd kps_api
python3 setup.py install
python3 -c "import kps_api"
retval=$?
if [[ "$retval" -ne "0" ]] ; then
  exit $retval
fi
popd
</pre>

# Note
You can open the documentation at kps_api/README.md

# Sample Code running steps, update KPS_API_KEY in following files before running.
<pre>
python3 samples/service_domain_client.py
python3 samples/application_client.py
</pre>

