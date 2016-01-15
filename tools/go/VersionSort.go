package versionsort

import (
	// "fmt"
	"sort"
	"strconv"
)

type List []string

func (list List) Len() int {
	return len(list)
}

func (list List) Less(i, j int) bool {

	left := list[i][:3]
	right := list[j][:3]

	fl, _ := strconv.ParseFloat(left, 16)
	fr, _ := strconv.ParseFloat(right, 16)
	if fl < fr {
		return true
	} else if fl > fr {
		return false
	} else {
		if len(list[i]) >= 5 {
			fl, _ = strconv.ParseFloat(list[i][4:5], 16)
		} else {
			fl = 0
		}
		if len(list[j]) >= 5 {
			fr, _ = strconv.ParseFloat(list[j][4:5], 16)
		} else {
			fr = 0
		}
		if fl < fr {
			return true
		} else if fl > fr {
			return false
		} else {
			return list[i] < list[j]
		}
	}
}

func (list List) Swap(i, j int) {
	var temp string = list[i]
	list[i] = list[j]
	list[j] = temp
}

func VersionSort(versions []string) []string {
	sort.Sort(List(versions))

	return versions
}

func Test() {
	// fmt.Println("...", VersionSort([]string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}))
}
