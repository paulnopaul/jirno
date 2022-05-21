sqlite3-db:
	rm -f jirno.db
	sqlite3 jirno.db < ./db/sql/init.sql

cobra-client: sqlite3-db
	go build -o jirno ./cmd/cobra-client/main.go

clean: 
	rm -f jirno jirno.db
