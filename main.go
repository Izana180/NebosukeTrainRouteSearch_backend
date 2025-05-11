package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/joho/godotenv"
)

type OdptStation struct {
    Title string `json:"dc:title"`
    ID    string `json:"@id"`
}

func main() {
    // .env 読み込み
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found. Falling back to system env.")
    }

    apiKey := os.Getenv("ODPT_API_KEY")
    if apiKey == "" {
        log.Fatal("ODPT_API_KEY is not set")
    }

    // ルーティング設定
    http.HandleFunc("/station", func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("name")
        if query == "" {
            http.Error(w, "Missing ?name= parameter", http.StatusBadRequest)
            return
        }

        // ODPT APIへリクエスト送信
        url := fmt.Sprintf("https://api.odpt.org/api/v4/odpt:Station?acl:consumerKey=%s", apiKey)
        resp, err := http.Get(url)
        if err != nil {
            http.Error(w, "Failed to request ODPT API", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        var stations []OdptStation
        if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
            http.Error(w, "Failed to decode ODPT response", http.StatusInternalServerError)
            return
        }

        // 駅名で部分一致検索
        var result []OdptStation
        for _, s := range stations {
            if strings.Contains(s.Title, query) {
                result = append(result, s)
            }
        }

        // 結果をJSONで返す
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
    })

    fmt.Println("🚀 Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}