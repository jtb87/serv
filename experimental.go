package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	draw2 "golang.org/x/image/draw"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func (a *App) InitExperimental() {
	exp := a.Router.PathPrefix("/exp").Subrouter()
	// exp.Use(authMiddleWareJWT)
	exp.HandleFunc("/select", a.selectHandler).Methods("GET")
	exp.HandleFunc("/context", a.contextHandler).Methods("GET")
	exp.HandleFunc("/interface", a.interfaceHandler).Methods("GET")
	exp.HandleFunc("/mutex", a.mutexHandler).Methods("GET")
	exp.HandleFunc("/file", a.fileUploadHandler).Methods("POST")
	exp.HandleFunc("/jsoninterface", a.jsonInterface).Methods("POST")
	exp.HandleFunc("/randompicture", a.randomPicture).Methods("GET")
	exp.HandleFunc("/logocombine", a.logoCombine).Methods("GET")
	exp.HandleFunc("/params", a.textMessage).Methods("POST")

}

func (a *App) textMessage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println(params)
	var v map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&v); err != nil {
		fmt.Println(err)
		respondWithError(w, "Invalid request payload")
		return
	}
	v["params"] = params
	fmt.Println(v)
	respondWithJSON(w, 200, v)
}

// logging request middleware
func authMiddleWareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Auth") != "testingKey" {
			respondWithJSON(w, 401, "[forbidden]")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *App) logoCombine(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	filename := wd + "/img/monk_logo.png"
	img, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer img.Close()
	pngImage, err := png.Decode(img)
	if err != nil {
		fmt.Println(err)
	}

	// Scale down by a factor of 3.
	sb := pngImage.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, sb.Dx()/3, sb.Dy()/3))
	draw2.BiLinear.Scale(dst, dst.Bounds(), pngImage, sb, draw.Over, nil)

	m := image.NewRGBA(pngImage.Bounds())
	transp := color.RGBA{100, 255, 200, 255}
	n := m.Bounds().Max.X - dst.Bounds().Max.X
	p := image.Rect(667, 0, (sb.Dx()/3 + 667), sb.Dy()/3)

	fmt.Println(n)
	draw.Draw(m, m.Bounds(), &image.Uniform{transp}, image.ZP, draw.Src)
	draw.Draw(m, p, dst, image.ZP, draw.Over)
	out, err := os.Create(wd + "/img/converted_image.png")
	if err != nil {
		fmt.Println(err)
	}
	err = png.Encode(out, m)
	if err != nil {
		fmt.Println(err)
	}
	respondWithJSON(w, 200, "all good!")

}

func (a *App) randomPicture(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	wd = wd + "/tmp"
	wdt, err := os.Open(wd)
	if err != nil {
		fmt.Println(err)
	}
	defer wdt.Close()
	rdir, err := wdt.Readdir(0)
	if err != nil {
		fmt.Println(err)
	}
	var filelist []string
	for _, f := range rdir {
		if strings.HasSuffix(f.Name(), ".png") {
			filelist = append(filelist, f.Name())

		}
	}
	if len(filelist) == 0 {
		respondWithError(w, "no files")
		return
	}

	rand.Seed(time.Now().Unix())
	file := filelist[rand.Intn(len(filelist))]
	filename := wd + "/" + file
	img, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	enc := base64.StdEncoding
	b64 := enc.EncodeToString(img)
	data := fmt.Sprintf("data:image/png;base64,%s", b64) // b := byte[]
	bod := map[string]string{"imgname": file, "img": data}
	respondWithJSON(w, 200, bod)
}

func (a *App) jsonInterface(w http.ResponseWriter, r *http.Request) {
	var v map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		fmt.Println(err)
		respondWithError(w, "Invalid request payload")
		return
	}

	fmt.Println(v)
	respondWithJSON(w, 200, "json received and printed")
}

type Upload struct {
	Name    string
	Content []byte
}

func (a *App) fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	v := Upload{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		fmt.Println(err)
		respondWithError(w, "Invalid request payload")
		return
	}
	filename := fmt.Sprintf("./tmp/%s", v.Name)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, "Could not save image")
		return
	}
	defer f.Close()
	f.Write(v.Content)
	respondWithJSON(w, 200, "file upload succesfull")
}

func (a *App) mutexHandler(w http.ResponseWriter, r *http.Request) {
	var mutex = sync.Mutex{}

	val := 0
	go func() {
		mutex.Lock()
		// time.Sleep(time.Second * 3 )
		fmt.Println(val)
		val += 100
		mutex.Unlock()
	}()
	go func() {
		mutex.Lock()
		fmt.Println(val)
		val += 1000
		mutex.Unlock()
	}()
	fmt.Println(val)
	time.Sleep(time.Second)
	fmt.Println(val)
	fmt.Println("call to mutex")
	respondWithJSON(w, 200, "mutex workings")
}

func (a *App) selectHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)

	ab := true
	ba := false
	aa := make(chan bool)
	bb := make(chan bool)

	go func() {
		// time.Sleep(time.Millisecond)
		bb <- ba
	}()
	go func() {
		aa <- ab
	}()
	go func() {
		aa <- ab
	}()
	fmt.Println(<-aa)
	select {
	case <-aa:
		respondWithJSON(w, 200, "a ==  true")
	case <-bb:
		fmt.Println(bb)
		respondWithJSON(w, 200, "b == true")
	}
}

func (a *App) contextHandler(w http.ResponseWriter, r *http.Request) {
	tch := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 3)
		tch <- true
	}()
	// if timeout occurs before channel tch returns the request will stop
	select {
	case <-r.Context().Done():
		fmt.Println("request has stopped")
		return
	case <-tch:
		fmt.Println("context point1")
		respondWithJSON(w, 200, "request completed")
	}
}

func (a *App) interfaceHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	v := type1{
		"Gregorius",
	}
	InterfaceExample(&v)
	v2 := type2{
		1245,
		"TMIB",
	}
	InterfaceExample(&v2)
	respondWithJSON(w, 200, "all good")
}

type I interface {
	Method1() string
}

func InterfaceExample(i I) {
	fmt.Println(i.Method1())
}

type type1 struct {
	Name string
}

func (t *type1) Method1() string {
	return t.Name

}

type type2 struct {
	Age  int
	Name string
}

func (t *type2) Method1() string {
	tmp := fmt.Sprintf("my name is: %s and I'm %d old", t.Name, t.Age)
	return tmp
}
