# Bulls and cows
Set answer size and player number, and start the game.  

## OOP approach
One Dealer type(class).  
One Player interface type(interface).  
One BasePlayer type(class).  
Two XXXPlayer type(class).  

## review of the OOP approach
In Go, the only Is-A relationship between two classes is the implementation of the interface. So don't think about parent&child classes when desiging the architecture of the program.  
Always composition over inheritance(there's no inheritance here).  
Always use interface to describe a group of classes, and always use pointer receiver.  

## FP approach  
Several functions related to game flow.  
Several functions related to player.  

Using closure to keep the player's info in the guess method of each player,
so we don't have to store the states in other functions and ask for them as arguments when taking a guess. In this way, we can make the function more independent and more easy to use.  

## review of the FP approach  
Try to think about which steps we'll need to make the job done, instead of trying to analogize the job into a story about some objects interactive with each others.  
Split one big functions into several smaller functions to reuse the logics.  
Use pure functions as much as possible.  
Use closure to keep states to simplify the logic and reduce the dependency between functions. Closure is very common in asynchronous codes because it can keep states for later use in callback functions.  
Design the data flow with pure functions and closures in your mind.   
 