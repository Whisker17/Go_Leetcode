#!/bin/bash
cd src
for x in ./*.go; do
    mkdir "${x%.*}" && mv "$x" "${x%.*}"
done
