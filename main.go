package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", func(c *gin.Context) {
		// membuat session
		session := sessions.Default(c)
		session.Set("hello", "world")
		session.Set("tes", "nama saya adalah sesuatu")
		session.Save()

		c.JSON(200, gin.H{
			"membuat session baru":      "membuat session baru gunakan .Set('nama_session':'data session') lalu akhiri dg .Save()",
			"mengakses data pd session": "mengakses data pada session gunakan .Get('nama_session_diakses':'data session yg diakses')",
			"session tes":               session.Get("tes"),
		})
	})
	r.GET("/logout", func(c *gin.Context) {
		// menghapus session
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1})
		session.Save()
	})
	r.GET("/data", func(c *gin.Context) {
		// mengakses session
		session := sessions.Default(c)
		fmt.Println(session.Get("tes"))
		c.JSON(200, gin.H{
			"session tes":   session.Get("tes"),
			"session hello": session.Get("hello"),
		})
	})
	r.Run(":8000")
}
