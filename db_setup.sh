#! /bin/bash

./osetools-cli initdb 
sqlite3 ./db_data/data.db < sql_setup_commands.txt