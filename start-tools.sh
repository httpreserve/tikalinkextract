#!/usr/bin/env bash

# Installation directory
instDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Location of Tika server jar
tikaServerJar="$instDir"/tools/tika-server-standard-3.1.0.jar

gnome-terminal -e "java -mx1000m -jar $tikaServerJar --port=9998 --noFork"
