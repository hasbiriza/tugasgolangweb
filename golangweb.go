package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Deklerasi Struct

type McDonald struct {
	Name     string
	Location string
	Menu     []string
}

type KFC struct {
	Name     string
	Location string
	Menu     []string
}

type AW struct {
	Name     string
	Location string
	Menu     []string
}

type PizzaHut struct {
	Name     string
	Location string
	Menu     []string
}

type Starbucks struct {
	Name     string
	Location string
	Menu     []string
}

func HandlerMcDonald(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mcd := []McDonald{
			{"McDonald's A", "Jl. McD 123", []string{"Big Mac", "McNuggets", "Fries"}},
			{"McDonald's B", "Jl. McD 456", []string{"Quarter Pounder", "Filet-O-Fish", "Apple Pie"}},
			{"McDonald's C", "Jl. McD 789", []string{"McFlurry", "Happy Meal", "McCafe"}},
		}
		res, err := json.Marshal(mcd)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerKFC(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		kfc := []KFC{
			{"KFC A", "Jl. KFC 123", []string{"Original Recipe Chicken", "Zinger Burger", "Fries"}},
			{"KFC B", "Jl. KFC 456", []string{"Twister Wrap", "Hot Wings", "Mashed Potatoes"}},
			{"KFC C", "Jl. KFC 789", []string{"Chicken Popcorn", "Corn on the Cob", "KFC Krushers"}},
		}
		res, err := json.Marshal(kfc)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerAW(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AW := []AW{
			{"AW A", "Jl. AW 123", []string{"Root Beer", "Coney Dog", "Curly Fries"}},
			{"AW B", "Jl. AW 456", []string{"Chicken Tenders", "Mozzarella Cheese Sticks", "A&W Float"}},
			{"AW C", "Jl. AW 789", []string{"Papa Burger", "Onion Rings", "A&W Waffle Sundae"}},
		}
		res, err := json.Marshal(AW)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerPizzaHut(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		PHD := []PizzaHut{
			{"Pizza Hut A", "Jl. PH 123", []string{"Supreme Pizza", "Pepperoni Lovers", "Hawaiian Pizza"}},
			{"Pizza Hut B", "Jl. PH 456", []string{"Stuffed Crust Pizza", "Veggie Lovers", "Chicken Wings"}},
			{"Pizza Hut C", "Jl. PH 789", []string{"Cheese Lovers", "Meat Lovers", "Garlic Breadsticks"}},
		}
		res, err := json.Marshal(PHD)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerStarbucks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sbuck := []Starbucks{
			{"Starbucks A", "Jl. Starbucks 123", []string{"Caramel Macchiato", "Java Chip Frappuccino", "Pumpkin Spice Latte"}},
			{"Starbucks B", "Jl. Starbucks 456", []string{"Flat White", "Iced White Chocolate Mocha", "Green Tea Latte"}},
			{"Starbucks C", "Jl. Starbucks 789", []string{"Nitro Cold Brew", "Chestnut Praline Latte", "Salted Caramel Mocha"}},
		}
		res, err := json.Marshal(sbuck)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/list-mcdonald", HandlerMcDonald)
	http.HandleFunc("/list-kfc", HandlerKFC)
	http.HandleFunc("/list-aw", HandlerAW)
	http.HandleFunc("/list-pizzahut", HandlerPizzaHut)
	http.HandleFunc("/list-starbucks", HandlerStarbucks)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
