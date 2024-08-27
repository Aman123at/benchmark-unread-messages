package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Benchmark test for newly unread message")
	fetchUnreadMessageCount()

	// one time insert process
	// bulkInsertMessagesIntoSQL()
}

func NewConn() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/indicator")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func fetchUnreadMessageCount() {
	db := NewConn()
	defer db.Close()
	query := "SELECT COUNT(DISTINCT fromuser) FROM messages WHERE touser='A' AND timestamp>'2024-08-27 18:42:24' limit 1"
	startTime := time.Now()
	_, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	log.Println("fetchUnreadMessageCount time:", time.Since(startTime))

	// without index ~ 200ms
	// index on touser ~ 150ms
	// index on touser and timestamp ~ 120ms
	// index on touser , timestamp and fromuser ~ 100ms
}

// var users = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
// var dumMsgs = []string{"Hey What are you doing ?", "All good", "How are you?", "How's your day?", "Great", "That sounds interesting!", "Just working on a project.", "What are you up to today?"}

// func getRandomFromAndToUser() (string, string) {
// 	rand.Seed(time.Now().UnixNano())
// 	fromIdx := rand.Intn(len(users))
// 	toIdx := rand.Intn(len(users))
// 	if fromIdx == toIdx {
// 		if toIdx == len(users)-1 {
// 			toIdx = toIdx - 1
// 		} else {
// 			toIdx = toIdx + 1
// 		}
// 	}
// 	randFrom := users[fromIdx]
// 	randTo := users[toIdx]
// 	return randFrom, randTo
// }

// func getRandomMsg() string {
// 	rand.Seed(time.Now().UnixNano())
// 	return dumMsgs[rand.Intn(len(dumMsgs))]
// }

// func getCurrentTimeStamp() string {
// 	return time.Now().Format("2006-01-02 15:04:05")
// }
// func bulkInsertMessagesIntoSQL() {
// 	dbconn := NewConn()
// 	stmt, err := dbconn.Prepare("INSERT INTO messages (id, msg, fromuser, touser, timestamp) VALUES (?, ?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatalf("Error on query exec : %v", err)
// 	}
// 	defer stmt.Close()

// 	for i := 0; i < 1000000; i++ {
// 		fromUser, toUser := getRandomFromAndToUser()
// 		randMsg := getRandomMsg()
// 		currTimestamp := getCurrentTimeStamp()
// 		_, err = stmt.Exec(i+1, randMsg, fromUser, toUser, currTimestamp)
// 		if err != nil {
// 			log.Fatalf("Error on query exec : %v", err)
// 		}
// 	}
// 	log.Println("Rows Inserted")
// }
