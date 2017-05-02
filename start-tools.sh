#!/usr/bin/env bash

gnome-terminal -e 'java -mx1000m -jar tools/tika-server-1.14.jar --port=9998'  #> /dev/null 2>&1 &
