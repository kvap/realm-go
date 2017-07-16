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

	err = sg.Write(func(g *realm.Group) bool {
		t, err := g.GetTableByName("Items")
		if err != nil {
			log.Printf("failed to get table: %v", err)
			return false
		}
		defer t.Destroy()

		cols, err := t.GetColumnCount()
		if err != nil {
			log.Printf("failed to get column count: %v", err)
			return false
		}
		if cols > 0 {
			return false
		}

		if _, err = t.AddColumn(realm.TypeInt, "a", false); err != nil {
			log.Printf("failed to add column: %v", err)
			return false
		}

		return true
	})
	if err != nil {
		log.Printf("write transaction failed: %v", err)
		return
	}

	err = sg.Write(func(g *realm.Group) bool {
		t, err := g.GetTableByName("Items")
		if err != nil {
			log.Printf("failed to get table: %v", err)
			return false
		}
		defer t.Destroy()

		col, err := t.GetColumnIndex("a")
		if err != nil {
			log.Printf("failed to get column index: %v", err)
			return false
		}

		row, err := t.AddEmptyRow()
		if err != nil {
			log.Printf("failed to add empty row: %v", err)
			return false
		}

		if err := t.SetInt(col, row, int64(1000 + row)); err != nil {
			log.Printf("failed to fill the empty row: %v", err)
			return false
		}

		return true
	})
	if err != nil {
		log.Printf("write transaction failed: %v", err)
		return
	}

	err = sg.Read(func(g *realm.Group) {
		t, err := g.GetTableByName("Items")
		if err != nil {
			log.Printf("failed to get table: %v", err)
			return
		}
		defer t.Destroy()

		col, err := t.GetColumnIndex("a")
		if err != nil {
			log.Printf("failed to get column index: %v", err)
			return
		}

		empty, err := t.IsEmpty()
		if err != nil {
			log.Printf("failed to check if table is empty: %v", err)
			return
		}

		if empty {
			log.Print("table is empty")
			return
		}

		size, err := t.GetSize()
		if err != nil {
			log.Printf("failed to get table size: %v", err)
			return
		}

		for row := 0; row < size; row++ {
			value, err := t.GetInt(col, row)
			if err != nil {
				log.Printf("failed to get int from table: %v", err)
				return
			}
			log.Printf("value[col=%d,row=%d] = %d", col, row, value)
		}
	})

	log.Print("done.")
}
