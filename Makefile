psql_docker: 
	# docker rm jirno_rb
	docker build -t jirno_postgres -f ./deploy/postgres.Dockerfile .
	docker run -p 5432:5432 --name jirno_db jirno_postgres

sqlite3-db: clean 
	sqlite3 jirno.db < ./db/sql/init.sql

cobra-client: sqlite3-db
	go build -o jirno ./cmd/cobra-client/main.go

tview-client: sqlite3-db
	go build -o jirno ./cmd/tview-client/main.go

clean:
	rm -f jirno jirno.db

