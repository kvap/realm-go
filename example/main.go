package main

import (
	"log"
	"github.com/kvap/realm-go"
)

func main() {
	sg, err := realm.NewSharedGroup("hello.realm", false)
	if err != nil {
		log.Printf("failed to open realm: %v", err)
		return
	}
	defer sg.Destroy()

	g, err := sg.BeginWrite()
	if err != nil {
		log.Printf("failed to begin write transaction: %v", err)
		return
	}

	hasTable, err := g.HasTable("Items")
	if err != nil {
		log.Printf("failed to check if table exists: %v", err)
		return
	}

	var t *realm.Table
	if !hasTable {
		t, err = g.AddTable("Items", true)
		if err != nil {
			log.Printf("failed to add table: %v", err)
			return
		}
		log.Print("added the table")
	} else {
		t, err = g.GetTableByName("Items")
		if err != nil {
			log.Printf("failed to get table: %v", err)
			return
		}
		log.Print("got the table")
	}
	defer t.Destroy()

	if err := sg.Commit(); err != nil {
		log.Printf("failed to commit: %v", err)
		return
	}

	log.Print("done.")
}
