#!/bin/sh -x
./ji user signup --name "Pavel Cheklin" --nickname paulnopaul --email tscheklin@gmail.com --password passpass

./ji user login --nickname paulnopaul --password passpass

./ji task create t1
./ji task create t2
./ji task create t3
./ji task create t4

./ji task filter
