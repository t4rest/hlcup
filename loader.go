package main

import (
	"archive/zip"
	"highload/models"
	"io/ioutil"
	"strings"
	"github.com/mailru/easyjson"
)

func ImportDataFromZip() error {
	r, err := zip.OpenReader("/tmp/data/data.zip")
	if err != nil {
		return err
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
			err = importUsers(bytes)
		case "locations":
			err = importLocations(bytes)
		case "visits":
			err = importVisits(bytes)
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
