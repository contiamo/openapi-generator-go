#!/bin/bash
set -e

for case in $(ls); do
	if [ -d $case ]; then
		echo "updating $case"
		cd $case
		rm -rf generated
		openapi-generator-go generate -o generated --spec api.yaml --package-name generatortest --log-level=info
		rm generated/router.go
		rm -rf expected
		cp -r generated expected
		echo "updated $case"
		cd ..
	fi
done
