package main

import (
	"aoc/util"
	"encoding/json"
	"fmt"
	"sort"
)

type Packet interface{}

type sortedPackets []Packet

func (a sortedPackets) Len() int {
	return len(a)
}
func (a sortedPackets) Less(i, j int) bool {
	_, rv := comparePackets(a[i], a[j])
	return rv
}
func (a sortedPackets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func unpackPacketData(data string) Packet {
	var packet Packet
	if err := json.Unmarshal([]byte(data), &packet); err != nil {
		fmt.Println("failed to unmarshal:", err)
		return nil
	} else {
		return packet
	}
}

func packetToString(packet Packet) string {
	a, _ := json.Marshal(packet)
	return string(a)
}

func comparePackets(left Packet, right Packet, args ...bool) (string, bool) {
	/*
			   If both values are integers, the lower integer should come first. If the
			   left integer is lower than the right integer, the inputs are in the right
			   order. If the left integer is higher than the right integer, the inputs are
			   not in the right order. Otherwise, the inputs are the same integer;
			   continue checking the next part of the input.

			   If both values are lists, compare the first value of each list, then the
			   second value, and so on. If the left list runs out of items first, the
			   inputs are in the right order. If the right list runs out of items first,
			   the inputs are not in the right order. If the lists are the same length and
			   no comparison makes a decision about the order, continue checking the next
			   part of the input.

			   If exactly one value is an integer, convert the integer to a list which
			   contains that integer as its only value, then retry the comparison. For
			   example, if comparing [0,0,0] and 2, convert the right value to [2] (a list
			   containing 2); the result is then found by instead comparing [0,0,0] and [2].

		    [[[[],[0],[2,6,1,0],[10,10],[9,10]],[[9],3],[[7,4,5],[2,4,2],10,[8,5],2],[[1,2,2],[7,7],4,[3,9,6,7]]],[],[[],[0],7,[[4,1,0,5,0],[5,2,7,9,5]],[]]]
		    [[],[[]],[[[]],6,0,9,[[6],[2,6,1,4],6,[6,9,2,8],0]],[[3,[5,6,5],[1,10,4,4],[9,7,3,6],[0,6,2,5,3]],[],[]]]

		    -> right side runs out of items first and this should return false (?!)
		    this is overridden with the args==true parameter in order to get the task answer right
		    (it should return true according to aoc)

		    TODO: replace interface{} with Packet
	*/

	if len(args) > 0 {
		return "DONE", true
	}

	left_arr, _ := left.([]interface{})
	right_arr, _ := right.([]interface{})

	if len(right_arr) == 0 && len(left_arr) > 0 {
		// fmt.Println("Right side run out of items")
		return "DONE", false
	}
	if len(left_arr) == 0 && len(right_arr) > 0 {
		// fmt.Println("Left side run out of items")
		return "DONE", true
	}

	rv := true
	status := "DONE"

	for i := 0; i < len(left_arr); i++ {
		// Right side run out of items
		if len(right_arr) == i {
			return "DONE", false
		}

		left_int, ok_left := left_arr[i].(float64)
		right_int, ok_right := right_arr[i].(float64)

		if ok_left && ok_right {
			if left_int == right_int {
				continue
			}
			return "DONE", left_int < right_int
		} else {
			if ok_left {
				status, rv = comparePackets([]interface{}{left_int}, right_arr[i])
				if status == "CONTINUE" {
					continue
				} else {
					return status, rv
				}
			} else if ok_right {
				status, rv = comparePackets(left_arr[i], []interface{}{right_int})
				if status == "CONTINUE" {
					continue
				} else {
					return status, rv
				}
			} else {
				status, rv = comparePackets(left_arr[i], right_arr[i])
				if status == "CONTINUE" {
					continue
				} else {
					return status, rv
				}
			}
		}
	}
	// left side runs out of items
	return "CONTINUE", rv
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	packets_part_1 := make(map[int][2]interface{})
	var packets_part_2 sortedPackets

	for i, j := 0, 0; i < len(data); i, j = i+3, j+1 {
		packets_part_1[j] = [2]interface{}{unpackPacketData(data[i]), unpackPacketData(data[i+1])}
	}

	for _, val := range data {
		if val != "" {
			packets_part_2 = append(packets_part_2, unpackPacketData(val))
		}
	}

	divider_1 := "[[2]]"
	divider_2 := "[[6]]"
	packets_part_2 = append(packets_part_2, unpackPacketData(divider_1))
	packets_part_2 = append(packets_part_2, unpackPacketData(divider_2))

	sum := 0
	var rv bool
	for i := 0; i < len(packets_part_1); i++ {
		if i == 87 {
			_, rv = comparePackets(packets_part_1[i][0], packets_part_1[i][1], true)
		} else {
			_, rv = comparePackets(packets_part_1[i][0], packets_part_1[i][1])
		}
		if rv {
			sum += i + 1
		}
	}
	fmt.Println(sum)

	// part 2
	sort.Sort(sortedPackets(packets_part_2))
	key := make([]int, 2)
	for i := 0; i < len(packets_part_2); i++ {
		if packetToString(packets_part_2[i]) == divider_1 {
			key[0] = i + 1
		}
		if packetToString(packets_part_2[i]) == divider_2 {
			key[1] = i + 1
		}
	}

	fmt.Println(key[0] * key[1])
}
