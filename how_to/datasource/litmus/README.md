## Litmus

The litmus data interface can be used to connect the programmable logic controllers you have set up in Litmus LoopEdge to Xi IoT. To start setting up this datasource, 
use the following example YAML and configure it to your LoopEdge.

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
  img: dataifc/litmus2nats:lit3
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
  topic: "natshost-<LoopEdge IP>"
- name: port
  topic: "natsport-4222"
```

Run the following command to create an instance of the datasource:
```console
xi-iot create -f litmus.yaml
```
* If the datasource was configured correctly, you should now see a datasource in your UI called litmus.