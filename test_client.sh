#!/bin/bash

if [ $2 = "create" ]; then
	curl -H'Content-type: application/json' -d '{"item":{"title":"Some thing to do","details":"This is very important!"}}' -XPOST $1/todos
elif [ $2 = "list" ]; then
	curl -H'Content-type: application/json' $1/todos
fi
