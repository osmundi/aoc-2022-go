
### Fetch input data
for i in {1..24}; do curl -v  "https://adventofcode.com/2022/day/${i}/input" -H 'Cookie: session=<sessionid>' > day_${i}/data.csv ; done

### Developer journal
Base experience := nil

#### Day 1
Before this I had only build the "Hello, World!" example with go and read about half an hour of the [tour of go](https://go.dev/tour/welcome/1)(thumbs up for this). The first challenges are luckily quite trivial. Nevertheless it took almost one hour(!) executing this succesfully. Of course there were some overhead of setting the modules and fetching data and so on. But with go it's quite straightforward to get things started.

#### Day 2
The whole time writing this thing I had the feeling that there must be a better way of doing this. But with the little spare time I have on my hands, I just went and wrote the game logic with simple if...else statements. At least learning some new data structes in the language (map, struct).

#### Day 3
This taught me how arrays (or slices, like they say in the go world) are built. It wasn't too easy to find the index of nth element from the slice! At least for an seasoned Python developer like me, who has used to a little too easy way of doing things. On the other hand, feels empowering to know more of the internals of the data structures while using them. This starts to feel like a tool for grown ups.

Note to self: similarities between the two functions (find_common_value(string, string) int, find_common_value_from_three_el(string, string, string) int) needs to be reafactored to single function taking slice of strings as a input if you want to sleep next night well...

#### Day 4
This was quite quickly done. Liked the fmt.Sscanf -method a lot.

#### Day 5
Not too proud of how this came off. And still not sure if I like how go format things. Dont like the readibility with this one at all:
```
// remove n elements from a stack
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

Notes after the first week:  I'm already flying with this tool without constantly having the need to refer the [documentation](https://pkg.go.dev/std)(which is OK)/google. I still don't have too strong opinion about the language. Really like the simplicity of it and coding with a statically typed language brings a sense of relaibility and robustness to the code. At first handling the pointer logic was a shock (after a 10 year break from writing last bits of C -code) but with this last challenge, it seemed the right way. No need to pollute the memory with unnecessary stuff whether there is a garbage collector or not.

#### Day 8


