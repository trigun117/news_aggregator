package main

import (
	"fmt"
	"github.com/trigun117/news_aggregator/fetch"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"runtime"
	"time"
)

var (
	emailFrom         = os.Getenv("EF")
	emailTo           = os.Getenv("ET")
	emailFromLogin    = os.Getenv("EFL")
	emailFromPassword = os.Getenv("EFP")
)

func newsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("templates/index.html"))
		t.Execute(w, &fetch.NewsArticles)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("templates/contact.html"))
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		body := fmt.Sprintf("Name: %s\n Email: %s\n Message: %s", r.Form["name"][0], r.Form["email"][0], r.Form["message"][0])
		msg := fmt.Sprintf("From: %s \nTo: %s \nSubject: Contact\n\n %s", emailFrom, emailTo, body)
		smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", emailFromLogin, emailFromPassword, "smtp.gmail.com"), emailFrom, []string{emailTo}, []byte(msg))
		http.Redirect(w, r, "/", 301)
	}
}

func server() {
	go func() {
		for {
			defer runtime.GC()
			fetch.FreshNews()
			time.Sleep(10 * time.Minute)
		}
	}()
	http.HandleFunc("/", newsHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":80", nil)
}
