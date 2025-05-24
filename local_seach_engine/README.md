# Buffer awful script

- **This script allows you to find the words you need in text. This script was implemented through a cheesy, crooked solution via an equally crappy buffer. But overall, it looks beautiful.**

- **Author as always is a ruthless and evil linuxoid who didn't think to install windows on qemu and test the program there**
  

**By the way, the second program that is made exactly for the task is in the folder real_hw**

- But still please rate this one too :)
---

# About the program and its work (briefly)

- The program is implemented in the **GO** language using an additional library called golang.org/x/term.
- The program (if it can be called so) is designed to search for text entered by the user
- It works

# A little bit about problems and unexpected situations

## “Stupid problem.”

In fact, the program is limited to text files only (i.e. .txt .cpp .py .asm). Files (I call them office files) such as .docs .odt .pdf it will not be able to read, because the text is presented there in a completely different way.

Of course, if you try to pass in an argument such a file, your pc will not burn, and the processor will not melt from any not expected spin block, and the sectors on your disk will not split like pangaea, but still, you should not do this. Otherwise you will get in the terminal a colorful drawing of fractions in the form of Japanese origami.

### illustrative example


![image](https://github.com/user-attachments/assets/bb230f93-b0bc-492d-b769-453d8e0a674a)


## Problem with colors in the terminal.

If you are using hyprland, custom colors in a terminal that you made yourself, or if you are sitting in tty. Of course, the program will work, but you may find the colors a bit dull, too light, or dark (like me), so it's better to install any other terminal (not x-term, I pray God) and roll there some universal theme like base-16 dark, it will look perfect.

- Example of how it looks with custom colors


![image](https://github.com/user-attachments/assets/fc116063-7ff2-4f4b-8beb-fb0fc7fc3ed9)


- Example of how it looks with normal colors


![image](https://github.com/user-attachments/assets/6dc28487-0ef4-4cc0-9997-d1750aa0847f)

### I think there's a difference by eye !

---

Otherwise, if the program complains that it cannot index any function, just run a script called `_check_dependencies.sh` and you will be happy.

---

What about the second program ? Why, the only problem is that when you try to run it more than 2 times, the number doesn't go up, and it's just like in Jumble (ералаш).


THE END!
