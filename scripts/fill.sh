#!/bin/sh -x
./jirno user signup --name "Pavel Cheklin" --nickname paulnopaul --email tscheklin@gmail.com --password passpass

./jirno user login --nickname paulnopaul --password passpass

./jirno task create "Task 1" --description "Task 1 Description" --dateto 2022-01-01
./jirno task create "Task 2" --description "Task 2 Description" --dateto 2022-01-02 
./jirno task create "Task 3" --description "Task 3 Description" --dateto 2022-01-03 
./jirno task create "Task 4" --description "Task 4 Description" --dateto 2022-01-04 
./jirno task create "Task 5" --description "Task 5 Description" --dateto 2022-01-05 

./jirno task filter

./jirno task complete 1
./jirno task filter

./jirno task update 1 --dateto 2022-02-02 --completed=False
./jirno task filter

