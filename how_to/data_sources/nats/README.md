## NATS

This data interface can be used to connect the NATS devices you have set up on the edge to Xi IoT. To start setting up this datasource, use the following example YAML 
and create your data source via the Xi IoT CLI.

**nats.yaml**
```yaml
kind: dataSource
name: ex_nats
svcDomain: <Service Domain>
protocol: DATAINTERFACE
type: Sensor
authType: CERTIFICATE
ifcInfo:
  class: DATAINTERFACE
  img: dataifc/nats:v1.0
  kind: IN
  protocol: nats
  ports:
    - name: rtmp
      port: 1935
edge: myprovideredge
fields:
- name: topic
  topic: "<Device Topic>"
- name: host
  topic: "natshost-<IP Address>"
- name: port
  topic: "natsport-<Port>"
```

**Note:** Your *Device Topic* is the NATS address your device communicates with. 

Run the following command to create an instance of the datasource:
```console
xi-iot create -f nats.yaml
```
* If the yaml was configured correctly, you should now see a datasource in your UI called *ex_nats*.
* This device can now be leveraged as a data source in entities such as Kubernetes Applications and Data Pipelines.

### Litmus

This data interface also has the ability to integrate with Litmus LoopEdge and you can use the same interface to communicate with your PLC devices with 
minor adjustments. The following YAML file is configured to connect with your LoopEdge.

**litmus.yaml**
```yaml
kind: dataSource
name: litmus
svcDomain: <Service Domain>
protocol: DATAINTERFACE
type: Sensor
authType: CERTIFICATE
ifcInfo:
  class: DATAINTERFACE
  img: dataifc/nats:v1.0
  kind: IN
  protocol: nats
  ports:
    - name: rtmp
      port: 1935
edge: myprovideredge
fields:
- name: secret
  topic: "secret-<LoopEdge Secret>"
- name: host
  topic: "natshost-<LoopEdge Address>"
- name: port
  topic: "natsport-<LoopEdge Port>"
```

If your LoopEdge connection was successful, your device names and topics should now be in the cloud as artifacts.
To see your newly created artifacts run the following command:
```console
xi-iot get datasource --show-artifacts
```

You should be met with a row entry labelled *litmus*. Inside this row entry you should see a list of your device tags and corresponding device topics:
```console
NAME         	ENDPOINTS                                             CLIENT SECRET  
litmus        	<names>:                                                   
              		devicehub.raw.<topic>.*                 
                <names>:                                                          
               		devicehub.raw.<topic>.*
```

You can now proceed to add whichever device topic you choose, and can handle this directly in the UI. Proceed to **Infrastructure** → **Data Sources** and click on *litmus*. 
Select **Edit** and you should be brought to a menu where you can define your data source topics. You should see your some topics such as port and secret already present. Now you 
can add the device topic of your choice and select **Update**.
* Once you have updated your data source, you will start ingesting data from your selected device and use this entity in Kubernetes Applications and Data Pipelines.

## FAQs

What happens if I remove the device topic and add a new one?
* The stream will get cancelled and your data source will start ingesting data from your new device topic.

Can I add multiple device topics to the same datasource?
* No, we recommend you create multiple data sources for each unique device topic. You can also configure multiple correlated sensors into one device topic in LoopEdge or create a Gateway in Xi IoT to aggregate data from multiple sources. 

Is there a way I can ingest multiple PLC sources under one data source?
* Yes, you can configure multiple devices under one device topic in LoopEdge.