id = GRL_1_C
no = 1
base = https://judgedat.u-aizu.ac.jp/testcases
dir = src/$(id)/cases/$(no)
in = src/$(id)/cases/$(no)/input.txt
out = src/$(id)/cases/$(no)/output.txt

.PHONY: download
download: $(dir) $(in) $(out)

$(dir):
	mkdir -p $(dir)

$(in):
	curl -o $(in) ${base}/$(id)/$(no)/in

$(out):
	curl -o $(out) ${base}/$(id)/$(no)/out

.PHONY: testgo
testgo: download
	echo "case $(no)"
	cat src/$(id)/cases/$(no)/input.txt | go run src/$(id)/main.go | ~/bin/cliassert -v -stdout-contain "`cat src/$(id)/cases/$(no)/output.txt`"

.PHONY: testpy
testpy: download
	echo "case $(no)"
	cat src/$(id)/cases/$(no)/input.txt | python3 src/$(id)/main.py | ~/bin/cliassert -v -stdout-contain "`cat src/$(id)/cases/$(no)/output.txt`"

.PHONY: opengo
opengo:
	cat src/$(id)/main.go | pbcopy
	open -a Safari http://judge.u-aizu.ac.jp/onlinejudge/submission.jsp#submit/$(id)

.PHONY: openpy
openpy:
	cat src/$(id)/main.py | pbcopy
	open -a Safari http://judge.u-aizu.ac.jp/onlinejudge/submission.jsp#submit/$(id)
