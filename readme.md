
# pnginx 


golang + redis/memory + fastcgi/反向代理 + php 

## usage
```
Usage of /tmp/go-build3712817361/b001/exe/main:
  -cache_refresh_key string
    	refreshKey key (default "key")
  -cache_time int
    	 cache 20s Second (default 20)
  -cache_type string
    	cache type is redis or memory or  not cache, default  not cache (default "nocache")
  -ext comma separated list
    	the fastcgi file extension(s) comma separated list (default "php")
  -fcgi string
    	the fcgi unix:///tmp/php-cgi-72.sock, you can pass more fcgi related params as query params (default "unix:///var/run/php/php7.0-fpm.sock")
  -http string
    	the http address to listen on (default ":6065")
  -index comma separated list
    	the default index file comma separated list (default "index.php,index.html")
  -listing
    	whether to allow directory listing or not
  -redis_db int
    	redis db default 0
  -redis_host string
    	redis host 192.167.1.22:6379 (default "192.167.1.22:6379")
  -redis_pass string
    	redis password default '' 
  -root string
    	the document root (default "./")
  -router string
    	the router filename incase of any 404 error (default "index.php")
  -rtimeout int
    	the read timeout, zero means unlimited
  -wtimeout int
    	the write timeout, zero means unlimited

```

## 缓存key  r.URL
```
// Middleware is the HTTP cache middleware handler.
func (c *Client) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := generateKey(r.URL.String())   //缓存key   r.URL
		next.ServeHTTP(w, r)
	})
}
```