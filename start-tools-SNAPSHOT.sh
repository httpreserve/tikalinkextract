#!/usr/bin/env bash

gnome-terminal -e 'java -mx1000m -jar tools/tika-server-1.15-SNAPSHOT.jar --port=9998'  #> /dev/null 2>&1 &
