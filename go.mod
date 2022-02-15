module tinygo.org/x/tinyterm

go 1.15

require (
	github.com/bgould/http v0.0.0-20190627042742-d268792bdee7
	github.com/valyala/fastjson v1.6.3
	tinygo.org/x/drivers v0.19.0
	tinygo.org/x/tinydraw v0.0.0-20220125063109-43cae6615eb5
	tinygo.org/x/tinyfont v0.2.1
)

replace tinygo.org/x/drivers => ../drivers
