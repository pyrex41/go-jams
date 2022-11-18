# go-jams

implementing https://nikolai.fyi/jams/

simple project to try to learn Go

Two tests currently fail to match JSON: double-quote and quotes-never-fail. For both, I think the standard
JSON parsing library is not handling this correctly:

``` go
╭─░▒▓ ~/go-jams │ main ───────────────────────────────── ✘ 1 │ 20:42:22 ▓▒░
╰─ go test
**FAILED** --->  double-quote
JAMS:
str\"ing
JSON:
str"ing

-------------------------



**FAILED** --->  quotes-never-fail
JAMS:
map[backslashed:\\double \\backslashes \\go \\anywhere es1:\/ es2:\\ es3:\n es4:\r es5:\b es6:\f es7:\t es8:\u0024]
JSON:
map[backslashed:\double \backslashes \go \anywhere es1:/ es2:\ es3:
 es5 es6:
          es7:   es8:$]

-------------------------



--- FAIL: TestPass (0.00s)
    jams_test.go:45: FAIL -- JAMS != JSON --  double-quote
    jams_test.go:45: FAIL -- JAMS != JSON --  quotes-never-fail
FAIL
exit status 1
FAIL    GOJAM/go-jams   0.114s


```
