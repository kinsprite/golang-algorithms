package greedy

// Activity Activity
type Activity struct {
	Name   string
	Start  int
	Finish int
}

// ActivitySelector Greedy Activity Selector
func ActivitySelector(activities []Activity) []Activity {
	n := len(activities)
	A := make([]Activity, 0, n)
	A = append(A, activities[0])
	k := 0

	for m := 1; m < n; m++ {
		if activities[m].Start >= activities[k].Finish {
			A = append(A, activities[m])
			k = m
		}
	}

	return A
}
