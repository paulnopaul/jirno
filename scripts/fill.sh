#!/bin/sh -x
./jirno user signup --name "Pavel Cheklin" --nickname paulnopaul --email tscheklin@gmail.com --password passpass

./jirno user login --nickname paulnopaul --password passpass

./jirno task create t1
./jirno task create t2
./jirno task create t3
./jirno task create t4

./jirno task filter 

./jirno task complete 2

./jirno task filter 

