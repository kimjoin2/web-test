module main

go 1.17

require (
	web-test/configFormat v0.0.0
)

replace (
	web-test/configFormat => ./configFormat
)