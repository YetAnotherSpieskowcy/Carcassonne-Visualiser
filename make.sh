#!/bin/sh

run () {
    echo "Running application..."
    go run . $2 -tags sdl
}

update () {
    echo "Updating Carcassonne-Engine..."
	go get -u github.com/YetAnotherSpieskowcy/Carcassonne-Engine@main
	echo "Updating raylib-go..."
	go get -v -u github.com/gen2brain/raylib-go/raylib
}

if [ $1 == "run" ]; then
    if [ $# -eq 2 ]; then
        run
    else
        echo "Missing log file name"
    fi
elif [ $1 == "update" ]; then
    update
else
    echo "Unknown command "
fi