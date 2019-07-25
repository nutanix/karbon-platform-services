#!/usr/bin/env node

/* jslint node: true */
'use strict';

var nats = require('nats');
const datatream_pb = require('./node-env/datastream_pb')
 
var natsURL = "nats://nats:4222";
//Topic name is same as datapipeline endpoint
var topic = "datapipeline-demo";
nats.DEFAULT_URI
// Connect to NATS server.
var nc = nats.connect({
    'encoding' : 'binary',
    'servers' : [natsURL]
});
nc.on('connect', function() {
    nc.subscribe(topic, function(msg) {
        var buf = Buffer.from(msg, 'binary');
        var arr = new Uint8Array(buf);
        var dsMsg = datatream_pb.DataStreamMessage.deserializeBinary(arr);
        var dsPayload = dsMsg.getPayload();
        var data = Buffer.from(dsPayload);
        console.log('Received "' + data.toString() + '"');
    });
    console.log('Listening on [' + topic + ']');
});

nc.on('error', function(e) {
    console.log('Error [' + nc.currentServer + ']: ' + e);
    process.exit();
});

nc.on('close', function() {
    console.log('CLOSED');
    process.exit();
});