#!/bin/sh

clear
echo "Hello, there is a crutch script for main.go. Please execute it to create file with text you want"
echo "------------------------------------------------------------------------------------------------>"
echo ""
echo "Input a filename which will be used as the name of the file from where go will read the text"
echo ""
read file_name

if [ -z "$file_name" ]; then
    echo "Error: Filename cannot be empty"
    exit 1
fi

if [ -f "$file_name" ]; then
    echo "File with this name already exists!"
    exit 1
fi

touch "$file_name"
error_code=$?

if [ $error_code -eq 0 ]; then
    echo "File created successfully!"
else
    echo "Check folder permissions and make sure that you do not execute it in / sub-directories"
    echo "Error code is >> $error_code"
    exit $error_code
fi

clear
echo "Cool! Now, please enter a text, this text will be allocated in the text file you created"
echo ""
read user_text

echo "$user_text" >> "$file_name"
error_code=$?

if [ $error_code -eq 0 ]; then
    echo "Text successfully written to the file!"
    echo "File contents:"
    cat "$file_name"
else
    echo "Error: Failed to write text to the file"
    echo "Error code is >> $error_code"
    exit $error_code
fi
