#!/bin/bash

DATE=$(date +%Y-%m-%d)
FILENAME="/backup/backup-$DATE.gz"

mongodump --host=mongodb --port=27017 --archive=$FILENAME --gzip
