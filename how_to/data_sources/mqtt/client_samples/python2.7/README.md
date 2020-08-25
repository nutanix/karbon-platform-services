# A simple MQTT publish subscribe example
mqtt-pub_sub-example.py is a simple example that shows how to
connect to an mqtt broker, publish a single message to a specific topic and
receive the published message back.

### Prerequisites
* A Nutanix edge with an IP address onboarded to Karbon Platform Services for IoT
* X509 certificates generated using Karbon Platform Services for IoT
* Python 2.7.10
* pip 10.0.1 (python 2.7)
* paho-mqtt. Install it for python 2.7.10 using the following command:
```
sudo pip2.7 install paho-mqtt
```

### Running the example
1. Download the certificates from Karbon Platform Services for IoT and store them locally under **certs**.
directory. Name the files as follows:
* ca.crt - Root CA certificate
* client.crt - client certificate
* client.key - client private key

2. Modify `broker_address` to point to the Karbon Platform Services for IoT edge IP address.

Run the example as follows:
```
$ python2.7 mqtt-example.py
```

Expected output:
> Connecting...  
> Connected to broker  
> Published!  
> New message received!  
> Topic:  test  
> Message:  Hello, World!
