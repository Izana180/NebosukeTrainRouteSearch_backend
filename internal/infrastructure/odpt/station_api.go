// 公共交通センターAPIから全国の駅データを取得
package odpt

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type odptStationResponse struct{
	Title string `json:"dc:title"`
	ID string `json:"@id"`
}

func FetchStationsFromOdpt(apiKey string) ([]odptStationResponse, error){
	if apiKey == "" {
		return nil, fmt.Errorf("APIKEY is not set")
	}

	url := fmt.Sprintf("https://api.odpt.org/api/v4/odpt:Station?acl:consumerKey=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("odpt API did not respond: %w", err)
	}
	defer resp.Body.Close()

	var raw []odptStationResponse
	// レスポンスの成形
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed decoding response: %w", err)
	}

	return raw, err
}
