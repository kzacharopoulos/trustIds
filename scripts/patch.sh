#!/usr/bin/env bash
#

source changelog.sh
awk -i inplace -F '.' '{print $1 "." $2 "." $3+1}' ../.version
changelog