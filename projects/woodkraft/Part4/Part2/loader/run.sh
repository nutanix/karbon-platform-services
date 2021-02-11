#!/bin/bash

echo "Starting loader"

locust -f /loadgen.py --headless -u $USER -r 1 -H $WP_HOST
