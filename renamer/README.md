# Stupid Rename

Stupid Rename is a program on go that is designed to change files in one way or another.
--
# Description 
Stupid Renamer Flag Combinations and file Changes
This document provides all possible flag combinations for the Stupid Renamer program, their effects on files in the /home/tenserflow/gopls/renamer/tests directory (containing img1.png and img2.png), and example outputs. The program is run with go run main.go from ~/gopls/renamer using Go 1.24.4 in the Fish shell, requiring quoted wildcard patterns (e.g., "*.png").
Program Overview


# Available Flags

--

## -dir - Directory to rename files in

String
./test_files/
Set to /home/tenserflow/gopls/renamer/tests in examples


## -pattern
File pattern (e.g., *.png)
String
*
Quote in Fish (e.g., "*.png")


## -p
Ad  prefix to filenames
String
""
Action flag


## -s
Add suffix to filenames (before extension)
String
""
Action flag


## -r
String to replace in filenames
String
""
Action flag; requires -w


## -w
Replacement string for -r
String
""
Used with -r


## -e
New extension (without dot)
String
""
Action flag


## -l
Convert filenames to lowercase
Bool
false
Action flag


## -u
Convert filenames to uppercase
Bool
false
Action flag


## -d
Enable dry run mode (preview without renaming)
Bool
false
Bonus feature


## -i
Enable interactive confirmation (prompt for each rename)
Bool
false
Bonus feature


Flag Combination Rules

Flag: Exactly one action flag for one programm call
Dependencies: -r requires -w.
Optional: -dir, -pattern, -d, -i.
Invalid Combinations:




Example Combinations and File Changes
Each example uses -dir /home/tenserflow/gopls/renamer/tests 

1. Prefix Action (-p)
Add a prefix to filenames.
1.1 Basic Prefix
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -p NEW_

File Changes:

img1.png → NEW_img1.png
img2.png → NEW_img2.pngOutput:

=== Stupid  Renamer  ===

Found 2 files

/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/NEW_img1.png
/home/tenserflow/gopls/renamer/tests/img2.png -> /home/tenserflow/gopls/renamer/tests/NEW_img2.png

Successfully renamed: 2 files

1.2 Prefix with Dry Run
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -p NEW_ -d

File Changes: None (dry run).Output: Same as above, but files remain img1.png, img2.png.
1.3 Prefix with Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -p NEW_ -i

File Changes (if y for both):

img1.png → NEW_img1.png
img2.png → NEW_img2.pngOutput (input: y, n):

===Stupid  Renamer ===

Found 2 files

Rename /home/tenserflow/gopls/renamer/tests/img1.png to /home/tenserflow/gopls/renamer/tests/NEW_img1.png? (y/n): y
Rename /home/tenserflow/gopls/renamer/tests/img2.png to /home/tenserflow/gopls/renamer/tests/NEW_img2.png? (y/n): n
/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/NEW_img1.png

Successfully renamed: 1 files

1.4 Prefix with Dry Run and Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -p NEW_ -d -i

File Changes: None (dry run).Output: Same as 1.3, but no changes applied.
2. Suffix Action (-s)
Add a suffix before the extension.
2.1 Basic Suffix
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -s _backup

File Changes:

img1.png → img1_backup.png
img2.png → img2_backup.pngOutput:

=== Stupid  Renamer ===

Found 2 files

/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/img1_backup.png
/home/tenserflow/gopls/renamer/tests/img2.png -> /home/tenserflow/gopls/renamer/tests/img2_backup.png

Successfully renamed: 2 files

2.2 Suffix with Dry Run
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -s _backup -d

File Changes: None.Output: Same as 2.1, but files unchanged.
2.3 Suffix with Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -s _backup -i

File Changes (if y for first, n for second):

img1.png → img1_backup.png
img2.png unchangedOutput:

=== Stupid  Renamer ===

Found 2 files

Rename /home/tenserflow/gopls/renamer/tests/img1.png to /home/tenserflow/gopls/renamer/tests/img1_backup.png? (y/n): y
Rename /home/tenserflow/gopls/renamer/tests/img2.png to /home/tenserflow/gopls/renamer/tests/img2_backup.png? (y/n): n
/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/img1_backup.png

Successfully renamed: 1 files

2.4 Suffix with Dry Run and Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -s _backup -d -i

File Changes: None.Output: Same as 2.3, but no changes applied.
3. Replace Action (-r with -w)
Replace a substring in filenames.
3.1 Basic Replace
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -r img -w photo

File Changes:

img1.png → photo1.png
img2.png → photo2.pngOutput:

=== Stupid  Renamer ===

Found 2 files

/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/photo1.png
/home/tenserflow/gopls/renamer/tests/img2.png -> /home/tenserflow/gopls/renamer/tests/photo2.png

Successfully renamed: 2 files

3.2 Replace with Dry Run
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -r img -w photo -d

File Changes: None.Output: Same as 3.1, but files unchanged.
3.3 Replace with Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -r img -w photo -i

File Changes (if y for both):

img1.png → photo1.png
img2.png → photo2.pngOutput (input: y, y):

=== Stupid  Renamer ===

Found 2 files

Rename /home/tenserflow/gopls/renamer/tests/img1.png to /home/tenserflow/gopls/renamer/tests/photo1.png? (y/n): y
Rename /home/tenserflow/gopls/renamer/tests/img2.png to /home/tenserflow/gopls/renamer/tests/photo2.png? (y/n): y
/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/photo1.png
/home/tenserflow/gopls/renamer/tests/img2.png -> /home/tenserflow/gopls/renamer/tests/photo2.png

Successfully renamed: 2 files

3.4 Replace with Dry Run and Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -r img -w photo -d -i

File Changes: None.Output: Same as 3.3, but no changes applied.
4. Extension Action (-e)
Change the file extension.
4.1 Basic Extension Change
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -e jpg

File Changes:

img1.png → img1.jpg
img2.png → img2.jpgOutput:

=== Stupid  Renamer ===

Found 2 files

/home/tenserflow/gopls/renamer/tests/img1.png -> /home/tenserflow/gopls/renamer/tests/img1.jpg
/home/tenserflow/gopls/renamer/tests/img2.png -> /home/tenserflow/gopls/renamer/tests/img2.jpg

Successfully renamed: 2 files

4.2 Extension with Dry Run
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -e jpg -d

File Changes: None.Output: Same as 4.1, but files unchanged.
4.3 Extension with Interactive
go run main.go -dir /home/tenserflow/gopls/renamer/tests -pattern "*.png" -e jpg -i

File Changes (if y for first, n for second):

