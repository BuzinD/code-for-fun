#!/bin/bash

CASE1='1
23
{
"dir": "root",
"files": [".zshrc"],
"folders": [
{
"dir": "desktop",
"files": ["config.yaml"]
},
{
"dir": "downloads",
"files": ["cat.png.hack"],
"folders": [
{
"dir": "kta",
"files": [
"kta.exe",
"kta.hack"
]
}
]
}
]
}'

echo $CASE1 | go run main.go