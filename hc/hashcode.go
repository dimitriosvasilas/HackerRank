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

func removeIndex(s []photo, index int) []photo {
	return append(s[:index], s[index+1:]...)
}

func makeSlideshow(photos []photo) slideshow {
	slideSh := slideshow{}
	//photosMap := make(map[int]bool)

	//for _, p := range photos {
	//	photosMap[p.id] = false
	//}

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(photos), func(i, j int) {
		photos[i], photos[j] = photos[j], photos[i]
	})
	//fmt.Println(photos)
	slidesH := make([]slide, 0)
	slidesV := make([]slide, 0)

	for _, ph := range photos {
		if ph.orientation == "H" {
			slide := makeHorizontalSlide(ph)
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

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(slidesH), func(i, j int) {
		slidesH[i], slidesH[j] = slidesH[j], slidesH[i]
	})

	slideSh.slides = slidesH
	return slideSh
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

		for i := 0; i <
			int(tagsN); i++ {
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
	runInput("in_b.txt", "out_b.txt")
	runInput("in_c.txt", "out_c.txt")
	runInput("in_d.txt", "out_d.txt")
	runInput("in_e.txt", "out_e.txt")
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
