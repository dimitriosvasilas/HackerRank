package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type photo struct {
	id          int
	orientation string
	tags        []string
}

type slide struct {
	photosN int
	photo1  photo
	photo2  photo
	tags    []string
}

type slideshow struct {
	slides []slide
}

func makeASlide1(photos []photo) slide {
	return slide{photosN: 2, photo1: photos[0], photo2: photos[1]}
}

func makeASlide2(photos []photo) slide {
	return slide{photosN: 2, photo1: photos[1], photo2: photos[2]}
}

func makeHorizontalSlide(ph photo) slide {
	return slide{photosN: 1, photo1: ph}
}

func makeVerticallSlide(p1 photo, p2 photo) slide {
	return slide{photosN: 2, photo1: p1, photo2: p2}
}

func removeIndex(s []slide, index int) []slide {
	return append(s[:index], s[index+1:]...)
}

func makeSlideshow(photos []photo) slideshow {
	slideSh := slideshow{}

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(photos), func(i, j int) {
		photos[i], photos[j] = photos[j], photos[i]
	})

	slidesH := make([]slide, 0)
	slidesV := make([]slide, 0)

	for _, ph := range photos {
		if ph.orientation == "H" {
			slide := makeHorizontalSlide(ph)
			slide.getTags()
			slidesH = append(slidesH, slide)
		}

	}
	complete := 0
	var temp photo
	for _, ph := range photos {
		if ph.orientation == "V" {
			if complete == 0 {
				temp = ph
				complete++
			} else if complete == 1 {
				slide := makeVerticallSlide(temp, ph)
				slide.getTags()
				slidesV = append(slidesV, slide)
				complete = 0
			} else {
				log.Fatal("oups")
			}
		}
	}

	for _, sl := range slidesV {
		slidesH = append(slidesH, sl)
	}
	/*
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(slidesH), func(i, j int) {
			slidesH[i], slidesH[j] = slidesH[j], slidesH[i]
		})
	*/
	result := slidesH
	max := 0
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(slidesH), func(i, j int) {
			slidesH[i], slidesH[j] = slidesH[j], slidesH[i]
		})
		tempScore := 0
		for i := 0; i < len(slidesH)-1; i++ {
			tempScore = tempScore + scoringFunction(slidesH[i], slidesH[i+1])
			if tempScore > max {
				result = slidesH
			}
		}
	}

	slideSh.slides = result
	return slideSh

	/*
		result := make([]slide, 0)
		result = append(result, slidesH[0])
		slidesH = removeIndex(slidesH, 0)

		l := len(slidesH)
		for i := 0; i < l; i++ {
			max := 0
			maxI := -1
			for i, sl := range slidesH {
				score := scoringFunction(sl, result[len(result)-1])
				if score > max {
					max = score
					maxI = i
				}

			}
			result = append(result, slidesH[maxI])
			slidesH = removeIndex(slidesH, maxI)

			fmt.Println("result")
			for _, sl := range result {
				fmt.Println(sl)
			}

		}
		/*
			fmt.Println("result")
			for _, sl := range result {
				fmt.Println(sl)
			}
			fmt.Println("remaining")
			for _, sl := range slidesH {
				fmt.Println(sl)
			}
			fmt.Println("end")
	*/
}

func runInput(in string, out string) {
	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	//n, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	photos := make([]photo, 0)

	i := 0
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		p := photo{id: i, orientation: line[0]}
		tags := make([]string, 0)

		tagsN, _ := strconv.ParseInt(line[1], 10, 0)

		for i := 0; i < int(tagsN); i++ {
			tags = append(tags, line[2+i])
		}

		p.tags = tags
		photos = append(photos, p)
		i++
	}

	//for _, p := range photos {
	//	fmt.Println(p)
	//}
	//fmt.Println()

	slideSh := makeSlideshow(photos)
	//fmt.Println(slideSh)

	slideSh.print(out)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runInput("in_a.txt", "out_a.txt")
	fmt.Println("a")
	runInput("in_b.txt", "out_b.txt")
	fmt.Println("b")
	runInput("in_c.txt", "out_c.txt")
	fmt.Println("c")
	runInput("in_d.txt", "out_d.txt")
	fmt.Println("d")
	runInput("in_e.txt", "out_e.txt")
	fmt.Println("e")

}

func scoringFunction(sl1 slide, sl2 slide) int {
	return minOf(intresectionLen(sl1.tags, sl2.tags), intresectionLen(sl1.tags, sl2.tags), diffLen(sl2.tags, sl1.tags))
}

func (sl *slide) getTags() {
	if sl.photosN == 1 {
		sl.tags = sl.photo1.tags
	} else {
		sl.tags = union(sl.photo1.tags, sl.photo2.tags)
	}
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func intresectionLen(a, b []string) int {
	count := 0
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			count++
		}
	}
	return count
}

func diffLen(a, b []string) int {
	count := 0
	mA := make(map[string]bool)
	mB := make(map[string]bool)

	for _, item := range a {
		mA[item] = true
	}
	for _, item := range b {
		mB[item] = true
	}
	for _, item := range a {
		_, okA := mA[item]
		_, okB := mB[item]
		if okA && !okB {
			count++
		}
	}
	return count
}

func (sls *slideshow) print(out string) {
	f, err := os.Create(out)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	fmt.Fprintln(w, len(sls.slides))
	for _, slide := range sls.slides {
		if slide.photosN == 1 {
			fmt.Fprintln(w, slide.photo1.id)
		} else {
			fmt.Fprintln(w, slide.photo1.id, slide.photo2.id)
		}
	}
	w.Flush()
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
