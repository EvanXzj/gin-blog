package main

import (
	"log"
	"time"

	"github.com/EvanXzj/gin-blog/models"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Cron Job Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
	// 试试goto?
}
