# web-test

### cli tool for testing web apis

### environment
~~~
go 1.17
using go mod
~~~

### if you want to run with local web server, run container before run web-test
run with docker command
~~~
run :
docker run -d -p 8118:80 --rm --name nginx_for_test nginx:alpine

stop :
docker kill nginx_for_test
~~~

### how to build and run
~~~
git clone https://github.com/kimjoin2/web-test.git
cd web-test
go build -o web-test .
./web-test -f test.json
~~~


