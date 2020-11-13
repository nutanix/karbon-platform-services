#!/bin/bash

curl -kv -i -X POST http://debezium-connect:8083/connectors \
 -H "Accept:application/json" \
 -H "Content-Type:application/json" \
 -d '{ "name": "wordpress-connector", "config": { "connector.class": "io.debezium.connector.mysql.MySqlConnector", "tasks.max": "1", "database.hostname": "mysql", "database.port": "3306", "database.user": "root", "database.password": "R00tMysql", "database.server.id": "184055", "database.server.name": "wordpress_db", "database.whitelist": "wordpress_db", "database.history.kafka.bootstrap.servers": "sherlock-kafka-svc:9092", "database.history.kafka.topic": "dbhistory.wordpress", "database.allowPublicKeyRetrieval": "true" } }'

