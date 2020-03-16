#!/bin/bash
DATE=$(date +"%Y-%m-%d")

mvn clean
tar --exclude='*.jar' --exclude="samples/" --exclude=".git" --exclude="*.class" -czvf ~/Backup/openapi-generator-$DATE.tar.gz .
