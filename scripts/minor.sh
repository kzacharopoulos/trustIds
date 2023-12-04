#!/usr/bin/env bash
#

source changelog.sh
awk -i inplace -F '.' '{print $1 "." $2+1 ".0"}' .version
changelog