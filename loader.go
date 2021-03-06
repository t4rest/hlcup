package main

import (
	"archive/zip"
	"github.com/mailru/easyjson"
	"hl/models"
	"io/ioutil"
	"strings"
)

func ImportDataFromZip() error {
	var r *zip.ReadCloser
	var err error
	r, err = zip.OpenReader("/tmp/data/data.zip")
	if err != nil {
		r, err = zip.OpenReader("/home/andrey/go/src/hlcupdoc/data/FULL/data/data.zip")

		if err != nil {
			return err
		}
	}
	defer r.Close()

	for _, f := range r.File {
		name := f.Name

		println(name)

		// cut off folder part
		i := strings.LastIndex(name, "/")
		if i != -1 {
			name = name[i:]
		}
		// cut off extension part
		i = strings.LastIndex(name, ".")
		if i != -1 {
			name = name[:i]
		}

		parts := strings.Split(name, "_")
		if len(parts) != 2 {
			continue
		}

		// add concurrent processing
		rc, err := f.Open()
		bytes, err := ioutil.ReadAll(rc)
		if err != nil {
			return err
		}
		err = rc.Close()
		if err != nil {
			return err
		}

		switch parts[0] {
		case "users":
			importUsers(bytes)
		case "locations":
			importLocations(bytes)
		case "visits":
			importVisits(bytes)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func importUsers(b []byte) error {

	var users models.Users

	err := easyjson.Unmarshal(b, &users)
	if err != nil {
		return err
	}

	models.InsertUsers(users)

	return nil
}

func importLocations(b []byte) error {
	var locations models.Locations
	err := easyjson.Unmarshal(b, &locations)
	if err != nil {
		return err
	}

	models.InsertLocations(locations)

	return nil
}

func importVisits(b []byte) error {
	var visits models.Visits
	err := easyjson.Unmarshal(b, &visits)
	if err != nil {
		return err
	}

	models.InsertVisits(visits)

	return nil
}
