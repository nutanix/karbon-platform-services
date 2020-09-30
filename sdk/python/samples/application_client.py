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


class ApplicationClient():

    def __init__(self, kps_cloud_endpoint, api_key):
        self.configuration, self.authorization = self._get_auth_config(kps_cloud_endpoint, api_key)
        # create an instance of the API class
        self.api_instance = kps_api.ApplicationApi(kps_api.ApiClient(self.configuration))
        self.project_api_instance = kps_api.ProjectApi(kps_api.ApiClient(self.configuration))

    def _get_auth_config(self, kps_cloud_endpoint, api_key):
        # Configure API key authorization: BearerToken
        configuration = kps_api.Configuration()
        configuration.host = kps_cloud_endpoint
        configuration.api_key['Authorization'] = api_key
        configuration.api_key_prefix['Authorization'] = 'Bearer'
        configuration.debug = False

        authorization = configuration.get_api_key_with_prefix('Authorization')
        return configuration, authorization

    def application_create_v2(self, body):
        try:
            api_response = self.api_instance.application_create_v2(body, self.authorization)
            # returns CreateDocumentResponseV2 object
            return api_response
        except ApiException as e:
            logging.error("Exception when calling ApplicationApi->application_create_v2: %s", e)
            raise

    def application_delete_v2(self, app_id):
        try:
            api_response = self.api_instance.application_delete_v2(self.authorization, app_id)
            # returns DeleteDocumentResponseV2 object
            return api_response
        except ApiException as e:
            logging.error("Exception when calling ApplicationApi->application_create_v2: %s", e)
            raise

    def project_list_v2(self, project_name=None):

        try:
            if project_name is not None:
                # int | 0-based index of the page to fetch results. (optional)
                page_index = 0
                page_size = 100  # int | Item count of each page. (optional)
                # order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
                # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
                filter = "name = \'%s\'" % project_name

                api_response = self.project_api_instance.project_list_v2(
                    self.authorization, page_index=page_index, page_size=page_size, filter=filter)
            else:
                api_response = self.project_api_instance.project_list_v2(self.authorization)
            # returns ProjectListPayload object
            return api_response
        except ApiException as e:
            logging.error("Exception when calling ProjectApi->project_list_v2: %s",e)
            raise

def main():
    """Test application_create_v2  and application_delete_v2 APIs

    Assumptions:
    - "Default Project" named project exists
    - Service Domain with one or more nodes exists
    """
    applicationClient = ApplicationClient(KPS_CLOUD_ENDPOINT, KPS_API_KEY)

    PROJECT_NAME="Default Project"
    projectListPayload = applicationClient.project_list_v2(PROJECT_NAME)
    for project in projectListPayload.result:
        logging.info("%s project id: %s", PROJECT_NAME, project.id)
        project_id = project.id
        break

    app_name = "flask-web-server-blog"
    with open("./data/flask-web-server.yaml", "r") as yamlFile:
        app_manifest = yamlFile.read()

    # ApplicationV2 | Describes the application creation request.
    body = kps_api.ApplicationV2(name=app_name, app_manifest=app_manifest, project_id=project_id)
    application = applicationClient.application_create_v2(body)
    pprint(application)

    appDeleteResp = applicationClient.application_delete_v2(application.id)
    pprint(appDeleteResp)


if __name__ == "__main__":
    main()
