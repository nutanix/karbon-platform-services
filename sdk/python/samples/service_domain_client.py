from __future__ import absolute_import

import logging
import sys
logging.basicConfig( stream=sys.stderr, format='%(funcName)s:%(levelname)s:%(message)s', level=logging.DEBUG )

import kps_api
from kps_api.rest import ApiException
from pprint import pprint

KPS_CLOUD_ENDPOINT='https://karbon.nutanix.com'
# You can get following key from Manage API Keys in KPS Cloud UI top bar menu
KPS_API_KEY="<KPS_API_KEY>"

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
            return api_response
        except ApiException as e:
            pprint("Exception when calling ServiceDomainApi->service_domain_create: %s\n" % e)

    def service_domain_get(self, svc_domain_id):
        try:
            # Get a service domain by its ID.
            api_response = self.api_instance.service_domain_get(svc_domain_id, self.authorization)
            return api_response
        except ApiException as e:
            pprint("Exception when calling ServiceDomainApi->service_domain_get: %s\n" % e)

    def service_domain_delete(self, svc_domain_id):
        try:
            # Delete a service domain as specified by its ID.
            api_response = self.api_instance.service_domain_delete(svc_domain_id, self.authorization)
            return api_response
        except ApiException as e:
            pprint("Exception when calling ServiceDomainApi->service_domain_delete: %s\n" % e)

def main():
    sdClient = ServiceDomainClient(KPS_CLOUD_ENDPOINT, KPS_API_KEY)

    body = kps_api.ServiceDomain(
        name = "sddemo1",
        description = "Demo service domain object created using python sdk apis"
    )
    sdCreateResp = sdClient.service_domain_create(body)
    pprint(sdCreateResp)

    svcDomain = sdClient.service_domain_get(sdCreateResp.id)
    pprint(svcDomain)

    sdDeleteResp = sdClient.service_domain_delete(svcDomain.id)
    pprint(sdDeleteResp)

if __name__ == "__main__":
  main()
