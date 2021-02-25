# Bulls and cows
Set answer size and player number, and start the game.

## OOP approach
One Dealer type(class).
One Player interface type(interface).
One BasePlayer type(class).
Two XXXPlayer type(class).

## review of OOP approach
In Go, the only Is-A relationship between two classes is the implementation of the interface. So don't think about parent&child classes when desiging the architecture of the program.
Always use interface to describe a group of classes, and always use pointer receiver.
