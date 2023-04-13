package t1

import (
	"bufio"
	"strconv"
	"strings"
)

type CallCentre struct {
	servers []bool
	countA  int
	countR  int
}

type CallCentres map[int]*CallCentre

func (c CallCentres) Disable(n, m string) {
	ni, _ := strconv.Atoi(n)
	mi, _ := strconv.Atoi(m)
	if c[ni-1].servers[mi-1] {
		c[ni-1].servers[mi-1] = false
		c[ni-1].countA--
	}
}

func (c CallCentres) RESET(n string) {
	ni, _ := strconv.Atoi(n)
	ni--
	for j := range c[ni].servers {
		c[ni].servers[j] = true
	}
	c[ni].countR++
}

func (c CallCentres) GETMAX() int {
	max := c[0].countA * c[0].countR
	cIndex := 0
	for i := range c {
		p := c[i].countA * c[i].countR
		if p > max {
			max = p
			cIndex = i
		}
	}
	return cIndex
}

func (c CallCentres) GETMIN() int {
	min := c[0].countA * c[0].countR
	cIndex := 0
	for i := range c {
		p := c[i].countA * c[i].countR
		if p < min {
			min = p
			cIndex = i
		}
	}
	return cIndex
}

func Kolya(scanner *bufio.Scanner) (res []int) {
	// init
	var n, m, q int
	{
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		n, _ = strconv.Atoi(s[0])
		m, _ = strconv.Atoi(s[1])
		q, _ = strconv.Atoi(s[2])
	}
	centres := CallCentres{}
	for i := 0; i < n; i++ {
		s := make([]bool, m)
		for j := range s {
			s[j] = true
		}
		centres[i] = &CallCentre{servers: s, countA: m, countR: 0}
	}

	// read command
	for i := 0; i < q; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		switch s[0] {
		case "DISABLE":
			centres.Disable(s[1], s[2])
		case "RESET":
			centres.RESET(s[1])
		case "GETMAX":
			res = append(res, centres.GETMAX()+1)
		case "GETMIN":
			res = append(res, centres.GETMIN()+1)
		}
	}
	return
}
