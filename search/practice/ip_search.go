package practice

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IPTable struct {
	// int of startIP and endIP all in the index
	index []int
	// startIP map to IPInfo[endIP and place]
	ipPlace map[int]IPInfo
}

type IPInfo struct {
	EndIP int
	Place string
}

func NewIPTable(ipPlaceMap map[string]string) (*IPTable, error) {
	index := make([]int, 0, len(ipPlaceMap))
	ipPlace := make(map[int]IPInfo)

	for rstr, place := range ipPlaceMap {
		r := strings.SplitN(rstr, ",", 2)
		if len(r) < 2 {
			return nil, fmt.Errorf("range format error, len %d", len(r))
		}
		startIP, endIp := IPToInt(r[0]), IPToInt(r[1])
		// 这里省略了校验ip的逻辑...
		index = append(index, startIP)
		ipPlace[startIP] = IPInfo{
			Place: place,
			EndIP: endIp,
		}
	}

	sort.Ints(index)

	return &IPTable{
		index:   index,
		ipPlace: ipPlace,
	}, nil
}

// @ip target ip
func (it *IPTable) ToPlace(ip string) (place string) {

	ipInt := IPToInt(ip)
	l, r := 0, len(it.index)-1
	mid := 0

	place = "notfound"
	for l <= r {
		mid = l + (r-l)>>1
		if ipInt < it.index[mid] {
			r = mid - 1
		} else {
			if mid == len(it.index)-1 || ipInt < it.index[mid+1] {
				ipInfo := it.ipPlace[it.index[mid]]
				if ipInt <= ipInfo.EndIP {
					return ipInfo.Place
				}
				return
			}
			l = mid + 1
		}
	}

	return
}

func IPToInt(ip string) int {
	ipSlice := strings.Split(ip, ".")

	ipInt := 0
	for i, v := range ipSlice {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return -1
		}
		vInt = vInt << ((len(ipSlice) - i - 1) * 8)
		ipInt = ipInt | vInt
	}

	return ipInt
}

func IntToIP(ipInt int) string {

	var sb strings.Builder
	slice := 0
	for i := 0; i < 4; i++ {
		slice = ipInt & (255 << i * 8)
		sb.WriteString(strconv.Itoa(slice))
		if i < 3 {
			sb.WriteByte('.')
		}
	}

	return sb.String()
}
