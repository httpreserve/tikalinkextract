@echo off

:: Installation directory
SET instDir="%~dp0\"

:: Location of Tika server jar
SET tikaServerJar=%instDir%\tools\tika-server-1.16.jar

stop java -jar  %tikaServerJar%

exit
