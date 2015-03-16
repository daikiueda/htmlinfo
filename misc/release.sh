#!/bin/sh

rm -rf ./.release

gox -os "windows linux darwin" -output ".release/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}"


for DIRNAME in `ls ./.release`
do
	cd ./.release/${DIRNAME}
	tar -czf ../${DIRNAME}.zip ./*
	cd ../..
	rm -rf ./.release/${DIRNAME}
	echo ./.release/${DIRNAME}.zip
done

