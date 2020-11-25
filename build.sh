#!/bin/bash

for app in $*; do
    echo $app
done
read -p "set version: " -t 20 version
echo $version
