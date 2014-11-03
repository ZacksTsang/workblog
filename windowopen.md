`window.open` 不能打开一个新标签问题
=========
##什么情形下是不可以的呢？

我发现在 angularjs 中的 `$http.get/put/post` 使用就会被浏览器拦截，经过多次尝试发现异步请求的都会被当成是弹出窗口，会被拦截。 但是直接使用 `button`的 `click` 事件则是在新的标签页打开的。

故在异步请求中不能使用 `window.open` 打开一个新的标签页，解决方案是：

```
  var win=window.open(""); //预先打开一个新的标签页
  
  $http.get('').success(function(data, status, headers, config){
      
      win.location.href=""; //需要打开的超链接
      
  }).error(function(data, status, headers, config){
  })

```
