#!/bin/bash

clear
for i in {1..3}
do
    echo -e "\033[32m=== VALIDATOR BOOT [$i/3] ===\033[0m"
    echo -e "\033[34m
  
   _____
  /     \\
 /_______\\
 |  YO!  |  Validator v1.0
 |_______|\033[0m"
    sleep 1.2
    clear
done

echo -e "\033[33m[*] Compiling the monster and erasing your /etc/ and /boot sections....\033[0m"
go build -o main
if [ $? -ne 0 ]; then
    echo -e "\033[31mCrash! Compilation failed, dude!\033[0m"
    exit 1
fi

echo -e "\033[32m[*] Your /etc/ and /boot sections are erased, please reboot your system and build gentoo from binaries !!!!!\033[0m"
sleep 0.5
echo -e " \033[32m[*] Your ip is 127.0.0.1 . Am i not right  ? :))) \033[0m"
go run main.go
