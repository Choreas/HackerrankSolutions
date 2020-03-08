package warmupchallenges

//JumpingOnClouds shows minimum jumps needed.
func JumpingOnClouds(c []int32) int32 {
	var jumps int32 = 0
	var idx int32 = 0
	// fmt.Println(len(c))
	for last := false; last == false; {
		idx, last, _ = jomp(c, idx)
		jumps++
	}

	return jumps
}

func jomp(c []int32, idx int32) (onCloud int32, last bool, err error) {
	last = false
	onCloud = idx
	err = nil
	if int32(len(c)-1)-2 >= onCloud {
		if c[onCloud+2] == 0 {
			onCloud = onCloud + 2
		}
	}
	if onCloud == idx {
		onCloud++
	}

	if int32(len(c)-1) == onCloud {
		last = true
	}

	return onCloud, last, err
}
