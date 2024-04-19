
# pnginx


golang + redis/memory + php 

客户端 -> golang -> redis/memory -> fcgi -> php

客户端 <- golang <- redis/memory <- fcgi <- php

```
// Middleware is the HTTP cache middleware handler.
func (c *Client) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := generateKey(r.URL.String())   //缓存key   r.URL
		next.ServeHTTP(w, r)
	})
}
```