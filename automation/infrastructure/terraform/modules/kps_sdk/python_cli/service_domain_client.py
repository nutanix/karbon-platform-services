from __future__ import absolute_import

import logging
import uuid
import sys
import time
logging.basicConfig( stream=sys.stderr, format='%(funcName)s:%(levelname)s:%(message)s', level=logging.DEBUG )

import kps_api
from kps_api.api.application_api import ApplicationApi
from kps_api.rest import ApiException
from pprint import pprint

KPS_CLOUD_ENDPOINT='https://karbon.nutanix.com'
KPS_API_KEY="<KPS API TOKEN>"

class ServiceDomainClient():

    def __init__(self, kps_cloud_endpoint, api_key):
        self.configuration, self.authorization = self._get_auth_config(kps_cloud_endpoint, api_key)
        # create an instance of the API class
        self.api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(self.configuration))

    def _get_auth_config(self, kps_cloud_endpoint, api_key):
        # Configure API key authorization: BearerToken
        configuration = kps_api.Configuration()
        configuration.host = kps_cloud_endpoint
        configuration.api_key['Authorization'] = api_key
        configuration.api_key_prefix['Authorization'] = 'Bearer'
        configuration.debug = False

        authorization = configuration.get_api_key_with_prefix('Authorization')
        return configuration, authorization

    def service_domain_create(self, body):
        try:
            # Create service domain.
            api_response = self.api_instance.service_domain_create(body, self.authorization)
            pprint(api_response)
            return api_response
        except ApiException as e:
            print("Exception when calling ServiceDomainApi->service_domain_create: %s\n" % e)

    def service_domain_get(self, svc_domain_id):
        try:
            # Get a service domain by its ID.
            api_response = self.api_instance.service_domain_get(svc_domain_id, self.authorization)
            pprint(api_response)
            return api_response
        except ApiException as e:
            print("Exception when calling ServiceDomainApi->service_domain_get: %s\n" % e)

    def service_domain_delete(self, svc_domain_id):
        try:
            # Delete a service domain as specified by its ID.
            api_response = self.api_instance.service_domain_delete(svc_domain_id, self.authorization)
            pprint(api_response)
        except ApiException as e:
            print("Exception when calling ServiceDomainApi->service_domain_delete: %s\n" % e)

def main():
    sdClient = ServiceDomainClient(KPS_CLOUD_ENDPOINT, KPS_API_KEY)

    body = kps_api.ServiceDomain(
        name = "sddemo1",
        description = "Demo service domain object created using python sdk apis"
    )
    sdCreateResp = sdClient.service_domain_create(body)

    svcDomain = sdClient.service_domain_get(sdCreateResp.id)

    sdDeleteResp = sdClient.service_domain_delete(sdCreateResp.id)

if __name__ == "__main__":
  main()
