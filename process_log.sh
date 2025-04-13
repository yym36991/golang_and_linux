#!/bin/bash
git log -1 | awk 'NR > 4 { if (!/Author/ && !/Date/) { x = (x ? x "," : "") $1 }} END { print x }'
