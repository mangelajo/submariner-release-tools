#!/bin/sh
set -e -x
GH_USER=${GH_USER:-mangelajo}

projects="admiral lighthouse submariner submariner-operator"

for project in $projects; do

	rm -rf $project || true
	git clone git@github.com:submariner-io/${project}.git
	pushd $project
	git remote add $GH_USER git@github.com:$GH_USER/${project}.git
	git checkout -B unpin-shipyard
	sed -i 's/shipyard-dapper-base:.*$/shipyard-dapper-base:devel/' Dockerfile.dapper
	git commit -a -s -S -m "Unpin shipyard after release"
	git push -f $GH_USER unpin-shipyard
	hub pull-request -m "Unpin shipyard after release" || true
	popd

done
