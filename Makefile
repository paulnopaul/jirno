sqlite3-db: clean 
	sqlite3 jirno.db < ./db/sql/init.sql

cobra-client: clean sqlite3-db
	go build -o jirno ./cmd/cobra-client/main.go

gocui-client: clean sqlite3-db
	go build -o jirno ./cmd/gocui-client/main.go

clean:
	rm -f jirno jirno.db

