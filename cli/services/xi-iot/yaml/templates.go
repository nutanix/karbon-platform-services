package yaml

const (
	DataSrc = `
	kind: datasource
	name: datasource-cli-test
	protocol: MQTT
	type: Sensor
	authType: CERTIFICATE
	fields:
	- name: clifield3
	  topic: mytopic1
	  fieldType: Custom
	- name: clifield4
	  topic: mytopic2
	  fieldType: Custom
	ifcInfo:
	  class: DATAINTERFACE
	  img: dataifc/mockifc:6.out
	  kind: IN
	  protocol: rtmp
	  ports:
		- name: http
		  port: 9090
	selectors:
	  - categoryName: testyaml
		categoryValue: datasource
		scope:
		  - clifield3
		  - clifield4
	`
)
