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

	err = sg.Write(func(g *realm.Group) bool {
		hasTable, err := g.HasTable("Items")
		if err != nil {
			log.Printf("failed to check if table exists: %v", err)
			return false
		}

		var t *realm.Table
		if !hasTable {
			t, err = g.AddTable("Items", true)
			if err != nil {
				log.Printf("failed to add table: %v", err)
				return false
			}
			log.Print("added the table")
		} else {
			t, err = g.GetTableByName("Items")
			if err != nil {
				log.Printf("failed to get table: %v", err)
				return false
			}
			log.Print("got the table")
		}
		defer t.Destroy()
		return true
	})
	if err != nil {
		log.Printf("write transaction failed: %v", err)
		return
	}

	log.Print("done.")
}
