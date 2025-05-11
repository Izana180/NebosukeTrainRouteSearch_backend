// 公共交通センターAPIから全国の駅データを取得
package odpt

import(
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
)

type odptStationResponse struct{
	Title string `json:"dc:title"`
	ID string `json:"@id"`
}

func FetchStationsFromOdpt() ([]*model.Station, error){
	apiKey := os.Getenv("ODPT_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("APIKEY is not set")
	}

	url := fmt.Sprintf("https://api.odpt.org/api/v4/odpt:Station?acl:consumerKey=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("odpt API did not respond: %w", err)
	}
	defer resp.Body.Close()

	var odptStations []odptStationResponse
	// レスポンスの成形
	if err := json.NewDecoder(resp.Body).Decode(&odptStations); err != nil {
		return nil, fmt.Errorf("failed decoding response: %w", err)
	}

	var stations []*model.Station
	for _, s := range odptStations {
		stations = append(stations, &model.Station{
			Name: s.Title,
			ID: s.ID,
		})
	}
	return stations, nil
}
