package main

import (
    "encoding/json"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getTasks(w, r)
    case http.MethodPost:
        createTask(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    // Veritabanından görevleri almak için kullanılacak kod buraya gelecek
}

func createTask(w http.ResponseWriter, r *http.Request) {
    // Yeni bir görev oluşturmak için kullanılacak kod buraya gelecek
}
