package handlers

import (
	"encoding/json"
	"main/internal/tools"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

// When an account acquires a new item in-game
func GotItem(w http.ResponseWriter, r *http.Request) {
	var item tools.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad request."))
		return
	}

	var existingItem tools.Item
    result := tools.DB.Where("account = ?", item.Account).Where("item_id = ?", item.ItemID).First(&existingItem)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // Record not found, handle accordingly
			tools.DB.Create(&item)
			w.WriteHeader(201)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
        } else {
			log.Error(result.Error)
        }
        return
    }

	// This account already has at least one of that item, update it.
	tools.DB.Model(&existingItem).UpdateColumn("amount", gorm.Expr("amount + ?", item.Amount))
	w.WriteHeader(201)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []tools.Item
	tools.DB.Where("account = ?", chi.URLParam(r, "name")).Find(&items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(items)
}