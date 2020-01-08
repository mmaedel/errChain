package main

import (

	//"io"
	"log"
	"os"
	//"reflect"
	"strconv"
	//"sync"
)

var errChain []Result = make([]Result, 0)

type check interface {
	errCheck(interface{}, error) bool
}

type Result struct {
	f   interface{}
	err error
}

func (er *Result) errCheck(f interface{}, e error) bool {

	if e != nil {

		errChain = append(errChain, Result{f, e})

		return false

	}

	errChain = append(errChain, Result{f, e})

	return true

}

func assign(f interface{}, e error) *Result {

	var c check

	r := new(Result)

	c = r

	if !c.errCheck(f, e) {

		return nil

	}

	return &Result{f, nil}

}

func main() {
	
	var r *Result

	for i := 0; i < 10; i++ {

		if r = assign(os.Open("test" + strconv.Itoa(i))); r != nil {

			f := (r.f).(*os.File)

			defer f.Close()

		}

	}

	for i, result := range errChain {

		if err := result.err; err != nil {

			log.Println(i, result.err)

			if r = assign(os.Create("test" + strconv.Itoa(i))); r != nil {

				f := (r.f).(*os.File)

				defer f.Close()

				log.Println(i, "write success")

				if r = assign(os.Open("test" + strconv.Itoa(i))); r != nil {

					f := (r.f).(*os.File)

					defer f.Close()

					log.Println(i, "read success")

				}

				continue

			}

		} else {

			log.Println(i, "read success")

		}

	}

}
