sqlite3-db: clean 
	sqlite3 jirno.db < ./db/sql/init.sql

cobra-client: clean sqlite3-db
	go build -o jirno ./cmd/cobra-client/main.go

tview-client: clean sqlite3-db
	go build -o jirno ./cmd/tview-client/main.go

clean:
	rm -f jirno jirno.db

