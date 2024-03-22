package controllers

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"html/template"
	"image/jpeg"

	"github.com/boltdb/bolt"
	"github.com/eahydra/gouuid"
	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"
	"github.com/nfnt/resize"

	//_ "frontEnd"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var world = []byte("world")
var cmnt = []byte("cmnt")
var key [50][]byte
var value [50][]byte

// var val []byte
var i int

type Imagee struct {
	Username  string
	Imagename [50][]byte
	Imm       []string
	userB     [50][]byte
	userS     []string
	commentS  []string
	tempstr   []string
}

/*
type structuring struct {
	Username  string
	Imagename [50][]byte
	Imm       []string
	userB     [50][]byte
	userS     []string
	commentS  []string
	tempstr   []string
}
*/

var v Imagee
var t *template.Template

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {

	fmt.Println("v r at helloserver")
	file, handler, err := req.FormFile("id-file-d")

	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	currenttime := time.Now().Nanosecond()
	//fmt.Println("Current time : ", currenttime.Format("2006-01-02 15:04:05 +0800"))
	// change both atime and mtime to currenttime

	tempstr := handler.Filename
	str := string(tempstr)
	fname := strconv.Itoa(int(currenttime)) + "_" + str
	fmt.Println("current time is ", currenttime)
	fmt.Println("fname is ", fname)
	//fmt.Println(string(currenttime.Format("2006-01-02 15:04:05 ")) + str)
	/*
		out, err := os.Create("/static/server/pictures/" + fname)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		fmt.Println(out)
		fmt.Println(err)
	*/
	err = ioutil.WriteFile("static/server/pictures/"+fname, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("error")
	fmt.Println(err)
	file1, err := os.Open("static/server/pictures/" + fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("error")
	fmt.Println(err)
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file1)
	if err != nil {
		log.Fatal(err)
	}
	file1.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(600, 400, img, resize.Lanczos3)

	out, err := os.Create("static/server/scaledLoc/" + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	//writing to database filename.db
	db, err := bolt.Open("filename.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := gouuid.NewUUID()
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	fmt.Println(uuidStr)

	key[i] = []byte(uuidStr)
	value[i] = []byte(fname)

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key[i], value[i])
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	i = i + 1

	var dbf string
	if strings.HasSuffix(fname, ".jpg") {
		dbf = strings.Replace(fname, ".jpg", ".db", 1)
	} else if strings.HasSuffix(fname, ".jpeg") {
		dbf = strings.Replace(fname, ".jpeg", ".db", 1)
	} else if strings.HasSuffix(fname, ".png") {
		dbf = strings.Replace(fname, ".png", ".db", 1)
	}
	fmt.Println(dbf)

	dbpath := "static/comments/" + dbf
	fmt.Println(dbpath)
	dbz, err := bolt.Open(dbpath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer dbz.Close()
	err = dbz.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(cmnt)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("007"), []byte("admin:header"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Accessdb(w http.ResponseWriter, req *http.Request) {
	fmt.Println("after update")

	v = Imagee{}
	t = template.New("Imagee template")
	t, err := t.Parse(frontend.Templ)
	checkError(err)

	db, err := bolt.Open("filename.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	j := 0
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", world)
		}
		c := bucket.Cursor()

		for keyy, valuu := c.First(); keyy != nil; keyy, valuu = c.Next() {
			// retrieve the data

			//v.Imagename[j] = bucket.Get(keyy)

			v.Imagename[j] = valuu
			fmt.Println(string(valuu))
			//fmt.Fprint(w, string(key[j]))

			//v.Imagename[j] = val

			//v.Imm[j] = string("C:/mygo/gowiki/server/pictures/" + v.Imm[j])

			//fmt.Fprintln(w, "---", string(val))
			//fmt.Fprintln(w, "---", v.Imm[j])
			j = j + 1

		}
		fmt.Println("no of files is", j)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	v.Username = GetUserName(req)
	fmt.Println(v.Username)
	v.Imm = make([]string, j+1)
	for k := 0; k <= j; k++ {
		v.Imagename[k] = bytes.Trim(v.Imagename[k], "") // leading and trailing

		v.Imm[k] = string(v.Imagename[k])
		//fmt.Println(v.Imagename[k])

	}

	v.Imm = v.Imm[:j]

	err = t.Execute(w, v)
	checkError(err)

	//err = t.Execute(os.Stdout, person)
	//checkError(err)

}

func HandleImages(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(w, "v r at disp")

	files, _ := filepath.Glob("static/server/pictures/*")

	l := len(files)
	for i := 0; i < l; i++ {
		files[i] = filepath.Base(files[i])
	}
	for i := 0; i < l; i++ {
		fmt.Fprintln(w, files[i])
	}
	data, _ := json.Marshal(files)
	fmt.Fprint(w, string(data))

	//return
	//l = len(files)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < l; i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buff := make([]byte, 512000) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
		_, err = file.Read(buff)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filetype := http.DetectContentType(buff)

		switch filetype {
		case "image/jpeg", "image/jpg", "image/gif", "image/png":
			files[i] = filepath.Base(files[i])
			fmt.Fprintln(w, files[i])

		}
	}
	//return
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// index page

func WriteCmntToDb(response http.ResponseWriter, request *http.Request) {
	currentImage := request.FormValue("tesla")
	currentComment := request.FormValue("currentComment")
	//fmt.Println(currentImage)
	//fmt.Println(currentComment)

	var dbf string
	if strings.HasSuffix(currentImage, ".jpg") {
		dbf = strings.Replace(currentImage, ".jpg", ".db", 1)
	} else if strings.HasSuffix(currentImage, ".jpeg") {
		dbf = strings.Replace(currentImage, ".jpeg", ".db", 1)
	} else if strings.HasSuffix(currentImage, ".png") {
		dbf = strings.Replace(currentImage, ".png", ".db", 1)
	}
	fmt.Println(dbf)

	dbpath := "static/comments/" + dbf
	fmt.Println(dbpath)
	db, err := bolt.Open(dbpath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	/*
			u := gouuid.NewUUID()
			uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

		uuidStr := time.Now().Nanosecond()
		fmt.Println(uuidStr)
	*/

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(cmnt)
		if err != nil {
			return err
		}
		c := bucket.Cursor()
		//keyLast := make([]byte, 40)
		keyC := make([]byte, 40)

		keyLast, _ := c.Last()
		fmt.Println("keyLast []byte:")
		fmt.Println(keyLast)
		fmt.Println("keyLast string" + string(keyLast))
		/*
			keyTemp, eror := strconv.Atoi(string(keyLast))
			if eror != nil {
				return eror
			}
			fmt.Println("keyTemp int :")
			fmt.Println(keyTemp)
			keyC = []byte(strconv.Itoa(keyTemp + 1))
			fmt.Println("keyC []byte :")
			fmt.Println(keyC)
		*/
		keyTemp, noneeds := binary.Uvarint(keyLast)
		fmt.Println("noneeds" + string(noneeds))
		fmt.Println("keyTemp int :")
		fmt.Println(keyTemp)
		keyTemp = keyTemp + 1
		noneed := binary.PutUvarint(keyC, keyTemp)
		fmt.Println("noneed" + string(noneed))
		fmt.Println("keyC []byte :")
		fmt.Println(keyC)
		fmt.Println("keyC======> " + string(keyC))
		valueC := []byte(GetUserName(request) + ":" + currentComment) //stuffing
		err = bucket.Put(keyC, valueC)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
func ReadCmntFromDb(response http.ResponseWriter, request *http.Request) {
	fmt.Println("ganesh kumar")
	currentImage := request.FormValue("tesla")
	fmt.Println(currentImage)
	/*
		p := structuring{}
		g := template.New("Structuring template")
		g, err := g.Parse(templ)
		checkError(err)
	*/
	var dbf string
	if strings.HasSuffix(currentImage, ".jpg") {
		dbf = strings.Replace(currentImage, ".jpg", ".db", 1)
	} else if strings.HasSuffix(currentImage, ".jpeg") {
		dbf = strings.Replace(currentImage, ".jpeg", ".db", 1)
	} else if strings.HasSuffix(currentImage, ".png") {
		dbf = strings.Replace(currentImage, ".png", ".db", 1)
	}
	fmt.Println(dbf)

	dbpath := "static/comments/" + dbf

	db, err := bolt.Open(dbpath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	j := 0
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(cmnt)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", cmnt)
		}
		c := bucket.Cursor()

		for keyy, valuu := c.First(); keyy != nil; keyy, valuu = c.Next() {
			// retrieve the data

			v.userB[j] = valuu
			fmt.Println(string(keyy) + "--" + string(v.userB[j]))
			j = j + 1
		}
		fmt.Println("no of comments is", j)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	data := ""
	v.userS = make([]string, j+1, j+1)
	v.commentS = make([]string, j+1, j+1)
	v.tempstr = make([]string, 3, 3)
	for k := 1; k < j; k++ {
		v.userS[k] = string(v.userB[k])
		v.tempstr = strings.Split(v.userS[k], ":")
		v.userS[k] = "  " + v.tempstr[0]
		v.commentS[k] = v.tempstr[1]
		fmt.Println("user : " + v.userS[k])
		fmt.Println("comment : " + v.commentS[k])
		data = data + "<h3>" + v.userS[k] + "</h3>" + "  " + v.commentS[k] + "<br><hr>"

	}
	fmt.Println(data)
	fmt.Fprintf(response, data)

	//	err = t.Execute(response, v)
	//	checkError(err)

}
func RegisterInDb(response http.ResponseWriter, request *http.Request) {
	db, err := bolt.Open("userDetails.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := request.FormValue("name")
	pass := request.FormValue("password")
	/*
		u := gouuid.NewUUID()
		uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
		fmt.Println(uuidStr)
	*/
	check := Checklogin(name, pass)
	if check == 0 {
		keyR := []byte(name)
		valueR := []byte(pass)

		// store some data
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(world)
			if err != nil {
				return err
			}

			err = bucket.Put(keyR, valueR)
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
		//	j = j + 1
		//fmt.Fprintln(response, "Successfully Registered")
		http.Redirect(response, request, "/internalnew", 302)
	} else {
		fmt.Fprintln(response, "username already exists, Try a new one")
	}

}

func Uploadimage(response http.ResponseWriter, request *http.Request) {
	fmt.Println("at upload image")
	userName := GetUserName(request)
	if userName != "" {
		fmt.Println(userName)
		fmt.Fprintf(response, "%s %s", frontend.AppHtml, userName)
	} else {
		fmt.Fprintln(response, "Sorry buddy, Please Login..!")
	}

}

func Viewimage(response http.ResponseWriter, request *http.Request) {
	fmt.Println("at viewimage")
	userName := GetUserName(request)
	if userName != "" {
		fmt.Println(userName)
		Accessdb(response, request)
	} else {
		fmt.Fprintln(response, "Sorry buddy, Please Login..!")
	}

}

// server main method

//var router = mux.NewRouter()
