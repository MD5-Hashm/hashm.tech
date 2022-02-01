#!/bin/bash

PID=`ps -ef | grep hashmtech | grep -v grep | awk '{print $2}'`
kill $PID
/projects/hashmweb/bin/uweb &
python3 /projects/hashmweb/bin/updater.py
PID=`ps -ef | grep uweb | grep -v grep | awk '{print $2}'`
kill $PID
/projects/hashmweb/bin/hashmtech &