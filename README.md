# go-jams

``` go
╭─░▒▓ ~/go-jams │ main !1 ?1 ───────────────────────────── ✔ │ 20:34:04 ▓▒░
╰─ go run .
JAMS:
map[basic_key:basic_value list_key:[item1 item2] nested:[map[key1:val1 key2:val2] map[key3:[val3 val4]]] str_key:superfluous nesting]

----------------
JSON:
map[basic_key:basic_value list_key:[item1 item2] nested:[map[key1:val1 key2:val2] map[key3:[val3 val4]]] str_key:superfluous nesting]

```


Two tests currently fail to match JSON: double-quote and quotes-never-fail. For double-quote,
not sure which implementation is correct. For quotes-never-fail, I think the standard
JSON parsing library is not handling this correctly:

``` go
╭─░▒▓ ~/go-jams │ main !1 ?1 ───────────────────────────── ✔ │ 20:34:04 ▓▒░
╰─ go test
JAMS:
str\"ing

-------------------------
JSON:
str"ing



JAMS:
map[backslashed:\\double \\backslashes \\go \\anywhere es1:\/ es2:\\ es3:\n es4:\r es5:\b es6:\f es7:\t es8:\u0024]

-------------------------
JSON:
map[backslashed:\double \backslashes \go \anywhere es1:/ es2:\ es3:
 es5 es6:
          es7:   es8:$]



--- FAIL: TestJams (0.00s)
    jams_test.go:45: parsed JAMS does not match parsed JSON for  double-quote
    jams_test.go:45: parsed JAMS does not match parsed JSON for  quotes-never-fail
FAIL
exit status 1
FAIL    GOJAM/go-jams   0.164s

```
