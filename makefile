# カレントディレクトリ内のテストを実行する
test :
	go test ./...

# カバレッジの測定
cover :
	go test -cover -coverprofile cover.out ./...
	go tool cover -html cover.out -o tmp/cover.html

# ベンチマークテストの実行
bench :
	go test -bench ./...
