#!/usr/bin/env bash

# Installation directory
instDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Location of Tika server jar
tikaServerJar="$instDir"/tools/tika-server-1.16.jar

gnome-terminal -e "java -mx1000m -jar $tikaServerJar --port=9998"  #> /dev/null 2>&1 &
