package repositoryimpl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/repository"
)

type routeRepository struct{}

func NewRouteRepository() repository.RouteRepository {
	return &routeRepository{}
}

// NAVITIME APIからのレスポンスを格納する構造体
type navitimeResponse struct {
	Items []struct {
		Summary struct {
			Move struct {
				Time int                `json:"time"`
				Fare map[string]float64 `json:"fare"`
			} `json:"move"`
		} `json:"summary"`
		Sections []struct {
			Type     string `json:"type"`
			Name     string `json:"name,omitempty"`
			FromTime string `json:"from_time,omitempty"`
			ToTime   string `json:"to_time,omitempty"`
			Move     string `json:"move,omitempty"`
		} `json:"sections"`
	} `json:"items"`
}

func (r *routeRepository) FetchRouteWithNodeid(from, to, datetime string, via []string, isArrivalTime bool) (*model.Route, error) {
	endPoint := "https://navitime-route-totalnavi.p.rapidapi.com/route_transit"

	params := url.Values{}
	params.Add("start", from)
	params.Add("goal", to)
	if len(via) > 0 {
		viaNodes := []string{}
		for _, v := range via {
			viaNodes = append(viaNodes, fmt.Sprintf(`{"node":"%s"}`, v))
		}
		params.Add("via", fmt.Sprintf("[%s]", strings.Join(viaNodes, ",")))
	}
	if isArrivalTime {
		params.Add("end_time", datetime)
	} else {
		params.Add("start_time", datetime)
	}

	requestURL := fmt.Sprintf("%s?%s", endPoint, params.Encode())

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("x-rapidapi-host", os.Getenv("NAVITIME_API_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("NAVITIME_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get response from navitimeAPI: status %d", resp.StatusCode)
	}

	// stream型のresp.Body要素をbyte列に変換
	// メモリ消費激しい処理なので、スケールを考えて後ほどnewDecoderに変更する
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp navitimeResponse
	// json解析して構造体に代入
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if len(apiResp.Items) == 0 {
		return nil, fmt.Errorf("no route found")
	}

	// 一旦一つの経路のみについて考える
	item := apiResp.Items[0]
	steps := []model.RouteStep{}

	var prevPoint string
	for i := 0; i < len(item.Sections); i++ {
		section := item.Sections[i]
		// NAVITIME APIのレスポンス ->
		// 基本的にpoint(主に駅), move(移動手段)が交互に来る
		switch section.Type {
		case "point":
			prevPoint = section.Name
		case "move":
			// moveの次に来るpointが到着駅
			arrival := ""
			for j := i + 1; j < len(item.Sections); j++ {
				if item.Sections[j].Type == "point" {
					arrival = item.Sections[j].Name
					break
				}
			}
			if section.Move != "walk" && prevPoint != "" && arrival != "" {
				steps = append(steps, model.RouteStep{
					DepartureStation: prevPoint,
					ArrivalStation:   arrival,
					Fromtime:         section.FromTime,
					Totime:           section.ToTime,
					Movetype:         section.Move,
				})
			}
		}
	}

	totalFare := 0
	if fare, ok := item.Summary.Move.Fare["unit_0"]; ok {
		totalFare = int(fare)
	}

	route := &model.Route{
		Steps:     steps,
		TotalTime: item.Summary.Move.Time,
		TotalFare: totalFare,
	}

	return route, nil
}