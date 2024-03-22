package controllers

import (
	//_ "controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"

	"github.com/boltdb/bolt"
)

// login handler

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	chk := Checklogin(name, pass)
	if chk == 1 {
		// .. check credentials ..
		SetSession(name, response)
		redirectTarget = "/internal"
		http.Redirect(response, request, redirectTarget, 302)
	} else {
		redirectTarget = "/errorlogin"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

func Checklogin(name, pass string) int {
	//check database of usernames

	db, err := bolt.Open("userDetails.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	cnt := 0
	flag := 0
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}
		c := bucket.Cursor()
		//names := []byte(name)
		for keyy, valuu := c.First(); keyy != nil; keyy, valuu = c.Next() {
			// retrieve the data
			cnt = cnt + 1
			if name == string(keyy) && pass == string(valuu) {
				fmt.Println("username & pwd exists")
				flag = 1
				return nil
			}

		}
		return nil
	})
	//fmt.Printf("%s does not exist out of %d usernames\n", name, cnt)
	return flag

}

// logout handler

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearSession(response)
	http.Redirect(response, request, "/", 302)
}

func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%s %s", frontend.IndexPage, "Login")
}
func IndexPageHandler1(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "%s %s", frontend.IndexPage, "Login Failed, Try again")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// internal page

func InternalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		fmt.Fprintf(response, "%s %s %s", frontend.InternalPage, "Login Successful", userName)
	} else {
		http.Redirect(response, request, "/", http.StatusFound)
	}
}
func InternalPageHandler1(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if userName != "" {
		fmt.Fprintf(response, "%s %s %s", frontend.InternalPage, "Successfully Registered", userName)
	} else {
		http.Redirect(response, request, "/", http.StatusFound)
	}
}
