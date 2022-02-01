# session-gin

### session biasa yg disimpan kedlm memory server
<p>go mod init github.com/muhsufyan/session-gin</p><br>
<p>go mod tidy</p><br>
<p>go get github.com/gin-gonic/gin</p><br>
<p>go get github.com/gin-contrib/sessions</p>

### session dengan redis
<p>docker run --name rediscache -p 6379:6379 -d redis</p><br>
<p>go mod download github.com/boj/redistore</p><br>
<p>go mod download github.com/gomodule/redigo</p>