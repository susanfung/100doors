# 100 Doors

100 doors in a row are all initially closed. You make 100 passes by the doors. The first time through, you visit every door and toggle the door (if the door is closed, you open it; if it is open, you close it).
The second time you only visit every 2nd door (door #2, #4, #6, ...).
The third time, every 3rd door (door #3, #6, #9, ...), etc, until you only visit the 100th door.

Question: What state are the doors in after the last pass? Which are open, which are closed?
Answer: As per the source, only doors that remain open are those whose numbers are perfect squares. The solution was used to verify the answer in the [final test](100doors_test.go#32).

[Source http://rosettacode.org]
