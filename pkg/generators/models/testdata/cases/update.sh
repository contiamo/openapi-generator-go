#!/bin/bash
set -e

for case in $(ls)
do
	if [ -d $case ];
	then
		cd $case
		rm -rf expected
		cp -r generated expected
		echo "updated $case"
		cd ..
	fi
done
