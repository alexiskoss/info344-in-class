package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//struct and fields must be exported to be
//used outside of models so must be capitlized
type Zip struct {
	Code  string
	City  string
	State string
}

//* points to a type; slice of pointers to Zip structs
type ZipSlice []*Zip

//maps of strings: zipslice
type ZipIndex map[string]ZipSlice

//returns a Zipslice or an error; can have multiple return types
//put at the end in parens; common to include error
func LoadZips(fileName string) (ZipSlice, error) {
	//returns file or error to variable
	f, err := os.Open(fileName)
	//if an error occurred...
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	//csv reader for zip codes; f is an input stream, can read one line at a time
	reader := csv.NewReader(f)
	//discard header/column row, don't need it
	//_ means don't keep value callable
	//don't need := because nothing is being declared
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	//preallocating the underlying array to be 43000 elements long
	//don't have to reallocate later, spending time
	zips := make(ZipSlice, 0, 43000)
	for {
		fields, err := reader.Read()
		//end of stream, signals the input is over
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		//creating a new *Zip every iteration to put in the zips Slice
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		//append to old zips the new zipcode z
		zips = append(zips, z)
	}
}
