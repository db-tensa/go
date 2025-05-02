#!/bin/bash
clear
echo "Hello ! This is just a little script to ensure that  all go dependencies and go itself are installed "
echo ""
echo "---------------------------------"
echo ""
echo "Please choose your system : "
echo ""
echo "1. -- MacOs  -- "
echo "2. -- Arch   -- "
echo "3. -- Ubuntu -- "
echo "4. -- Fedora -- "
echo ""
echo "INPUT -->  "

read choise

if [ "$choise" = "1" ]; then
    if brew list go >/dev/null 2>&1; then
        echo "Go is already installed via Homebrew."
    else
        echo "Go is not installed. Installing via Homebrew..."
        brew install go
    fi
elif [ "$choise" = "2" ]; then
    if pacman -Q go >/dev/null 2>&1; then
        echo "Go is installed on your system"
    else 
        echo "Go is not installed on your system"
        sudo pacman -S go 
    fi
elif [ "$choise" = "3" ]; then
    if apt list --installed 2>/dev/null | grep -q golang; then
        echo "Go is installed via apt."
    else
        echo "Go is not installed via apt."
        sudo apt install -y golang-go
    fi
elif [ "$choise" = "4" ]; then
    if dnf list installed golang >/dev/null 2>&1; then
        echo "Go is installed via dnf."
    else
        echo "Go is not installed via dnf."
        sudo dnf install -y golang
    fi
fi


echo "---------------------------------"
echo ""
echo "Now let's create go.mod and install tcell !"
echo ""
echo "---------------------------------"
echo ""
go mod init re_homework
go get github.com/gdamore/tcell/v2

echo ""
echo "ENJOY !"
echo ""
echo "You can and you should type << make >> to run go project !"

