
### Fetch input data
for i in {1..24}; do curl -v  "https://adventofcode.com/2022/day/${i}/input" -H 'Cookie: session=<sessionid>' > day_${i}/data.csv ; done

### Developer journal
Base experience := nil

#### Day 1
Before this I had only build the "Hello, World!" example with go and read about half an hour of the [tour of go](https://go.dev/tour/welcome/1)(thumbs up for this). The first challenges are luckily quite trivial. Nevertheless it took almost one hour(!) executing the first one succesfully. Of course there were some overhead of setting the modules and fetching data and so on. But with go it's quite straightforward to get things started.

#### Day 2
The whole time writing this thing I had the feeling that there must be a better way of doing this. But with the little spare time I have on my hands, I just went and wrote the game logic with simple if...else statements. At least learning some new data structes on the go (map, struct).

#### Day 3
This taught me how arrays (or slices, like they say in the go world) are built. It wasn't too easy to find the index of nth element from the slice! At least for an seasoned Python developer like me, who has used to a little too easy way of doing things. On the other hand, feels empowering to know more of the internals of the data structures while using them. This starts to feel like a tool for grown ups.

Note to self: similarities between the two functions (find_common_value(string, string) int, find_common_value_from_three_el(string, string, string) int) needs to be reafactored to single function taking slice of strings as a input if you want to sleep next night well...

#### Day 4
This was quite quickly done. Liked the fmt.Sscanf -method a lot.

#### Day 5
Not too proud of how this came off. And still not sure if I like how go format things. Dont like the readibility with this one at all:
```
// remove last n elements from a stack
stacks[move_from-1] = stacks[move_from-1][:len(stacks[move_from-1])-move_n]
```
Or maybe this should be done differently in the first place...

Note to self: 
 - Should I use byte or rune as a data type for single letter?
 - First time using a generic type in a method (for reversing an array, not sure if this is ok).
 - Write more functions for verbosity, now everything is done in the main function.

#### Day 6
Still a piece of cake.  

Note to self: is it ok to use copy-paste code (util.Unique -method) if you're trying to really learn this stuff? Sometimes it might be justified timewise. Mostly not.

#### Day 7
Pointers, (kind of?) linked list, recursion. This was so far maybe the most sophisticated answer. First time felt like I was doing something right. At least on the part where the filesystem is initialized. The part 2 answer was hacked together in a quick and dirty way. 

Notes after the first week:  I'm already flying with this tool without constantly having the need to refer the [documentation](https://pkg.go.dev/std)(which is OK)/google. I still don't have too strong opinion about the language. Really like the simplicity of it and coding with a statically typed language brings a sense of relaibility and robustness to the code. At first handling the pointer logic was a shock (after a 10 year break from writing last bits of C -code) but with this last challenge, it seemed the right way to do it. No need to pollute the memory with unnecessary stuff whether there is a garbage collector or not.

#### Day 8
Figuring out the edge cases took more time than implementing the algorithm itself. 

Note to self: Read specs with care. I spend too much time debugging my algorithm when the wanted answer was a multiple of the scenic views instead of a sum.
Note to self: compare performance with a python program made with the same slow algorithm.

#### Day 9
Learned a constructor pattern in go when dynamically creating ropes in the task. Otherwise nothing new, just spammed 70 lines of if-else logic.

#### Day 10
Sounded trivial task, but took suprisingly long...cycle step where the register must be mutated caused some headache. Second part was easier as it has been so far in other tasks also. It seems that already at day 10, solving the algorithm takes more time than managing the quirks of the language i'm learning.

#### Day 11
First time I had too look for external help. The wording in part two was little too misleading for me and I didn't understand too look for common denominator between the divisible rules. I first tought there was some kind of optimal pattern so only the monkeys with addition operations could pass the items between each other, so the worry levels wouldn't increase with so fast pace. Nothing too fancy codewise.

#### Day 12
Hardest so far. Reminded of the university CS studies back in the day. Implemented the pathfinding algorithm with A\*. Came out pretty neat. I used the pseudocode from [wikipedia](https://en.wikipedia.org/wiki/A*_search_algorithm) as a reference (some of the rows didn't even need any alterations :D). Had to make simple visualization of the path, but it would be nice to make an animation of the pathfinding in work as a bonus (test https://github.com/gdamore/tcell ?).

#### Day 13
This was actually quite usefull. Learned about interfaces. How to read ambigious (e.g. JSON) data into an empty interface and how to use the sorting interface (need to define Len, Less and Swap -methods to a type in order to use the general Sort function from sort package). I used type assertions to get the underlying concrete value of Packet interface (list/int). In python or similar plain old type() -check would have been enough and a lot more verbose than this.

Note to self: maybe you should check how someone else has unmarshalled the list data and what kind of datastructure has been used there.

#### Day 14
Good looking cave. It was fun filling it with sand. There is probably lot of room to optimize the algorithm, but nevertheless it run the program and visualization pretty swiftly (~1s). 

#### Day 15
For the first time part 2 took most of the time. It came clear quite soon that this wouldn't be reasonable to brute force, so the algorithm searchers the southermost intersection point of every possible sensor coverage area and tests only the points one step lower of this intersecion point (561 points needed to be looked instead of 1,6×10¹³). Theres probably a multitude of ways to solve this, so it would be interesting to see some other solutions.

#### Day 16
I really tried to figure this out on my own. Tried to read about different (weighted) graph traversal alhgorithms and travelling salesman problem, but soon figured out that this is too much for me. Or at least too much compared to the time im ready to invest into. So after checking couple neat python solutions, reading about Floyd–Warshall algorithm, dynamic programming and watching (c++) explanation video from youtube, I had a some kind of idea of how to do this and implemented recursive dynamic programming algorithm. In the end it looked easy and run quickly, but still felt like i didn't earn even a half start out of this.

#### Day 17
Part 1 was quite straightforward. Tried to optimize code for the part 2, but soon realized that this couldn't be brute forced in a reasonable time. At least learned to do some basic code profiling with go. Finally found the pattern by saving wind gusts that has moved rock and the amount of moves rock has taken. This was so exhausting that I just ended doing lot of the calcuting with python intepreter so the go program wont give straight answer to part 2.

#### Day 18
Managed to get only one start again. I tried to solve part 2 by trying to find outer limits of the droplet from every open side of the single cubes. Algorithm worked on the test data and with a bit more complex test data I generated myself (because the initial one was too simple). I didn't find out the edge case that brokes the DFS algorithm which traverses through the airpocket and tries to find out of the droplet. Lots of room for optimization also.. part 2 ran over 5s.

#### Day 19
Did with the same idea like day 16 (recursive DP). The implementation came out pretty long (>300 lines). Probably went little bit too far with the OOP and interfaces. At least the code came out somewhat readable. Did some optimization with limiting the amount of unnecessary robots and memoizing allready visited states but still it run quite long (~ 30s). Without these, it would have ran indefinetly. Code spends currently most of the time hashing the state and accesing it. Luckily the recursion depth didnt went too deep in this one, so I dont need to rethink some optional solution. This will do.

#### Day 20
No fancy algorithms, recursion or anything like that. Enjoyed a lot in a long time. This task could be a goood primer to learn about array indexing. I ended up reading [this](https://go.dev/blog/slices-intro) pretty well written introduction to arrays/slices. Struggled little bit with slicing a pointer (slice expression binds stronger than a dereference, so you must surround the pointer with parenthesis before slicing: (*arr)[n:n] ). Modulo arithmetics works differently (-2 % 7 != 5) in go than for example in python, which was bit suprising.  Also struggled a bit with off by one error (finally!). 

#### Day 21
Algorithm solves recursively the value of branch from binary syntax tree where the 'humn' variable cannot be found and buffers the operations that connects it to the another branch, to a list until the 'humn' leaf has been found. After that the buffered inversed operations are applied to the the right side of the root equation. At first I had problems with precision values and resorted to using even the Decimal package, but soon enough figured out that I had bug in the inversed arithmetics. 

#### Day 22

















