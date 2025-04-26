@echo off

:: Installation directory
SET instDir="%~dp0\"

:: Location of Tika server jar
SET tikaServerJar=%instDir%\tools\tika-server-standard-3.1.0.jar

start java -mx1000m -jar %tikaServerJar% --port=9998 --noFork
