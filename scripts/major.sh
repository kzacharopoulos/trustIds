#!/usr/bin/env bash
#

source changelog.sh
awk -i inplace -F '.' '{print $1+1 ".0.0"}' .version
changelog