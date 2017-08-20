package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Student struct {
	email string
	batch string
	cgpa  float64
	wdev  float64
	wdep  float64
	exp   float64
	quiz  float64
	score float64
}

type Students []Student

func (slice Students) Len() int {
	return len(slice)
}

func (slice Students) Less(i, j int) bool {
	return slice[i].score < slice[j].score
}

func (slice Students) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func MinIntVarible(v1 int, vn ...int) (m int) {
	m = v1
	for i := 0; i < len(vn); i++ {
		if vn[i] < m {
			m = vn[i]
		}
	}
	return
}

var expMap = make(map[string]float64)
var batchMap = make(map[string]string)

func init() {
	expMap["0"] = 0.0
	expMap["0-6"] = 1.0
	expMap["6-12"] = 2.0
	expMap["12-18"] = 3.0
	expMap["18-24"] = 4.0
	expMap["more than 2 years"] = 5.0

	batchMap["B.Tech 3rd Year"] = "B3"
	batchMap["B. Tech 4th Year"] = "B4"
	batchMap["M. Tech 1st Year"] = "M1"
	batchMap["M. Tech 2nd Year"] = "M2"
}

func get_slice(r io.Reader) Students {
	s := make(Students, 0, 100)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) < 9 {
			continue
		}

		email := fields[1]
		batch := batchMap[fields[2]]
		cgpa, _ := strconv.ParseFloat(fields[3], 64)
		wdev, _ := strconv.ParseFloat(fields[4], 64)
		wdep, _ := strconv.ParseFloat(fields[5], 64)
		exp := expMap[fields[6]]
		quiz, _ := strconv.ParseFloat(fields[7], 64)

		score := 0.65*cgpa + 0.20*wdev + 0.15*wdep + 0.05*quiz + 0.15*exp

		t := Student{
			email: email,
			batch: batch,
			cgpa:  cgpa,
			wdev:  wdev,
			wdep:  wdep,
			exp:   exp,
			quiz:  quiz,
			score: score,
		}

		s = append(s, t)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Invalid input: %s", err)
	}
	return s
}

func print_slice(students []Student, b bool) {
	if b {
		var sum float64
		for i := range students {
			sum += students[i].score
		}
		fmt.Println("Score: ", sum)
	}
	w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, '\t', 0)

	for i := range students {
		fmt.Fprintf(w, "%s\t%s\t%f\n", students[i].email, students[i].batch, students[i].score)
	}
	w.Flush()
}

func main() {
	f, err := os.Open("fcs.tsv")
	if err != nil {
		log.Fatalln("Could not open file")
	}
	defer f.Close()

	// A buffered reader has better performance
	r := bufio.NewReader(f)

	students := get_slice(r)

	b3 := make(Students, 0)
	b4 := make(Students, 0)
	m1 := make(Students, 0)
	m2 := make(Students, 0)

	for st := range students {
		switch students[st].batch {
		case "B3":
			b3 = append(b3, students[st])
		case "B4":
			b4 = append(b4, students[st])
		case "M1":
			m1 = append(m1, students[st])
		case "M2":
			m2 = append(m2, students[st])
		}
	}
	sort.Sort(sort.Reverse(b3))
	sort.Sort(sort.Reverse(b4))
	sort.Sort(sort.Reverse(m1))
	sort.Sort(sort.Reverse(m2))

	groups := make(map[int]Students)

	min := MinIntVarible(len(b3), len(b4), len(m1), len(m2))

	var i int
	var b3_min, b3_min2, b3_min3, b3_max, b4_min, m1_max, m2_min Student
	for i = 0; i < min; i++ {
		b3_min, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_max, b3 = b3[0], b3[1:]

		b4_min, b4 = b4[len(b4)-1], b4[:len(b4)-1]
		m1_max, m1 = m1[0], m1[1:]
		m2_min, m2 = m2[len(m2)-1], m2[:len(m2)-1]

		groups[i] = append(groups[i], b3_min)
		groups[i] = append(groups[i], b3_max)
		groups[i] = append(groups[i], b4_min)
		groups[i] = append(groups[i], m1_max)
		groups[i] = append(groups[i], m2_min)
	}

	for j := 0; j < len(m1)+2; j++ {
		b3_min, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_min2, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_max, b3 = b3[0], b3[1:]

		m1_max, m1 = m1[0], m1[1:]
		m2_min, m2 = m2[len(m2)-1], m2[:len(m2)-1]

		groups[i] = append(groups[i], b3_min)
		groups[i] = append(groups[i], b3_max)
		groups[i] = append(groups[i], b3_min2)
		groups[i] = append(groups[i], m1_max)
		groups[i] = append(groups[i], m2_min)

		// increment i : as it is the group index
		i++
	}

	for j := 0; j < len(m2)+2; j++ {
		b3_min, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_min2, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_min3, b3 = b3[len(b3)-1], b3[:len(b3)-1]
		b3_max, b3 = b3[0], b3[1:]

		m2_min, m2 = m2[len(m2)-1], m2[:len(m2)-1]

		groups[i] = append(groups[i], b3_min)
		groups[i] = append(groups[i], b3_max)
		groups[i] = append(groups[i], b3_min2)
		groups[i] = append(groups[i], b3_min3)
		groups[i] = append(groups[i], m2_min)

		// increment i : as it is the group index
		i++

	}
	for j := 0; j < 5; j++ {
		b3_min, b3 = b3[len(b3)-1], b3[:len(b3)-1]

		groups[i] = append(groups[i], b3_min)

	}
	// increment i : as it is the group index
	i++
	for j := 0; j < 4; j++ {
		b3_min, b3 = b3[len(b3)-1], b3[:len(b3)-1]

		groups[i] = append(groups[i], b3_min)

	}
	i++

	index := i

	fmt.Println("")
	for i := 0; i < index; i++ {
		fmt.Printf("GRP# : %d\n", i)
		print_slice(groups[i], false)
		fmt.Println("")

	}
	//	fmt.Println("Remaining b3", len(b3))
	//	print_slice(b3, false)
	//	fmt.Println("Remaining b4", len(b4))
	//	print_slice(b4, false)
	//	fmt.Println("Remaining m1", len(m1))
	//	print_slice(m1, false)
	//
	//	fmt.Println("Remaining m2", len(m2))
	//	print_slice(m2, false)
}
