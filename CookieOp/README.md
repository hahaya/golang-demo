CookieOp --- 使用golang简单操作cookie  

注意：  
当第一次在浏览器中输入`http://localhost:8080/`进行访问时，浏览器中会提示`http: named cookie not present`，这是因为第一次访问时r *http.Request中没有cookie，刷新浏览器第二次访问时，就会正常显示cookie中的值。