# Creating a Log Collector

Read about log collectors details and requirements in the [Xi IoT Admin Guide](https://portal.nutanix.com/#/page/docs/details?targetId=Xi-IoT-Infra-Admin-Guide:edg-iot-logs-c.html).

Each sample YAML file defines a log collector. Log collectors can be:
- Infrastructure-based: collect infrastructure-related (service domain) information
- Project-based: collect project-related information (applications, data pipelines, and so on)

You can connect each log collector to the destination cloud's logging infrastructure:
- *AWS CloudWatch*
- *AWS Kinesis Data Stream*
- *AWS Kinesis Data Firehose*
- *GCP StackDriver*

## Notes

1. Log conversion with `code` is not currently supported.
2. *AWS CloudWatch* is supported at the moment.

## Example Usage

Create a log collector defined in a YAML file:

`user@host$ xi-iot create -f log-collector.yaml`

### infra-logcollector-cloudwatch.yaml

This sample infrastructure log collector collects logs for a specific tenant, then forwards them to : *AWS CloudWatch*.

``` yaml
kind: logcollector
name: infra-log-name
type: infrastructure
destination: cloudwatch
cloudProfile: cloud-profile-name
awsRegion: monitoring.us-west-2.amazonaws.com
cloudwatchGroup: cloudwatch-group-name
cloudwatchStream: cloudwatch-stream-name
filterSourceCode: ""
```

| Field Name          | Value or Subfield Name / Description | Value or Subfield Name / Description                |
|---------------------|--------------------------------------|-----------------------------------------------------|
| kind                | `logcollector`                       | Specify the resource type                           |
| name                | `infra-log-name`                     | Specify the unique log collector name               |
| type                | `infrastructure`                     | Log collector for infrastructure                    |
| destination         | `cloudwatch`                         | Cloud destination type                              |
| cloudProfile        | `cloud-profile-name`                 | Specify the AWS cloud profile                       |
| awsRegion           | `monitoring.us-west-2.amazonaws.com` | A valid AWS region name or cloudwatch endpoint FQDN |
| cloudwatchGroup     | `cloudwatch-group-name`              | Specify the log group name                          |
| cloudwatchStream    | `cloudwatch-stream-name`             | Specify the log stream name                         |
| filterSourceCode    | ` `                                  | Specify the log conversion code                     |

Need to specify all `awsRegion`, `cloudwatchStream` and `cloudwatchGroup` to enable `cloudwatch` log streaming.

Names and limitations:
- `awsRegion` must be a valid AWS region name (`us-west-2`) or cloudwatch endpoint FQDN (`monitoring.us-west-2.amazonaws.com`)
- `cloudwatchGroup` and `cloudwatchStream` can be between 1 and 512 characters long. Allowed characters include a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.'

### infra-logcollector-kinesis.yaml

This sample infrastructure log collector collects logs for a specific tenant, then forwards them to: *AWS Kinesis*.

``` yaml
kind: logcollector
name: infra-log-name
type: infrastructure
destination: kinesis
cloudProfile: cloud-profile-name
awsRegion: us-west-2
kinesisStream: kinesis-stream-name
filterSourceCode: ""
```

| Field Name       | Value or Subfield Name / Description | Value or Subfield Name / Description  |
|------------------|--------------------------------------|---------------------------------------|
| kind             | `logcollector`                       | Specify the resource type             |
| name             | `infra-log-name`                     | Specify the unique log collector name |
| type             | `infrastructure`                     | Log collector for infrastructure      |
| destination      | `kinesis`                            | Cloud destination type                |
| cloudProfile     | `cloud-profile-name`                 | Specify the AWS cloud profile         |
| awsRegion        | `us-west-2`                          | A valid AWS region name               |
| kinesisStream    | `kinesis-stream-name`                | Specify the AWS Kinesis stream name   |
| filterSourceCode | ` `                                  | Specify the log conversion code       |

Both `awsRegion` and `kinesisStream` are required for `kinesis` log collector destination.

Names and limitations:
- `kinesisStream` can be between 1 and 512 characters long. Allowed characters include a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.'

Note: *AWS Kinesis Data Firehose* is not currently supported. This sample is for information purposes only.

### project-logcollector-firehose.yaml

This sample project log collector collects logs for a specific project, then forwards them to *AWS Kinesis Data Firehose*).

``` yaml
kind: logcollector
name: project-log-name
type: project
project: project-name
destination: firehose
awsRegion: us-west-2
kinesisStream: firehose-delivery-stream-name
filterSourceCode: ""
```

| Field Name       | Value or Subfield Name / Description | Value or Subfield Name / Description  |
|------------------|--------------------------------------|---------------------------------------|
| kind             | `logcollector`                       | Specify the resource type             |
| name             | `project-log-name`                   | Specify the unique log collector name |
| type             | `project`                            | Log collector for specific project    |
| project          | `project-name`                       | Specify the project name              |
| destination      | `firehose`                           | Cloud destination type                |
| cloudProfile     | `cloud-profile-name`                 | Specify the GCP cloud profile         |
| awsRegion        | `us-west-2`                          | A valid AWS region name               |
| kinesisStream    | `firehose-delivery-stream-name`      | Specify the delivery stream name      |
| filterSourceCode | ` `                                  | Specify the log conversion code       |

Both `awsRegion` and `kinesisStream` are required for `firehose` log collector destination.

Names and limitations:
- `kinesisStream` can be between 1 and 512 characters long. Allowed characters include a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.'

Note: *AWS Kinesis Data Firehose* is not currently supported. This sample is for information purposes only.

### project-logcollector-stackdriver.yaml

This sample project log collector collects logs for a specific project, then forwards them to *GCP StackDriver*).

``` yaml
kind: logcollector
name: project-log-name
type: project
project: project-name
destination: stackdriver
cloudProfile: cloud-profile-name
filterSourceCode: ""
```

| Field Name       | Value or Subfield Name / Description | Value or Subfield Name / Description  |
|------------------|--------------------------------------|---------------------------------------|
| kind             | `logcollector`                       | Specify the resource type             |
| name             | `project-log-name`                   | Specify the unique log collector name |
| type             | `project`                            | Log collector for specific project    |
| project          | `project-name`                       | Specify the project name              |
| destination      | `stackdriver`                        | Cloud destination type                |
| cloudProfile     | `cloud-profile-name`                 | Specify the GCP cloud profile         |
| filterSourceCode | ` `                                  | Specify the log conversion code       |

The `stackdriver` integration does not require additional parameters.

Note: *GCP Stackdriver* is not currently supported. This sample is for information purposes only.
