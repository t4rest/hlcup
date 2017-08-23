#!/usr/bin/env bash

#nohup mysqld -u root &
#echo | ls /tmp/data/
#unzip -j /tmp/data/data.zip '*.json' -d data
#(echo y | nohup mysqld -u root) &
mysqld -u root --initialize-insecure
nohup mysqld -u root &
sleep 10
#mysqladmin -u root status

#mysql -u root < echo "SET PASSWORD FOR 'root'@'localhost' = PASSWORD(1234);"
mysql -u root < go/src/highload/storage/scheme.sql

./go/bin/highload 80 data