# 概要
![image](https://github.com/user-attachments/assets/3cc72708-f1a8-4244-a71c-e67a48881a2c)  
ねぼすけあんないのバックエンドプロジェクトです。  
フロント:  
https://github.com/Izana180/onGoing_Nebosuke_trainRouteSearch

# 使用技術
- Go with Gin
- AWS EC2
- NAVITIME API
- ODPT API

# 提供機能
- 全駅情報の取得  
  フロントでの検索時サジェスト表示用に、全駅情報を取得します。
- 経路検索  
  発着駅をはじめとするデータを受け取り、経路データを取得します。

# 詳細仕様
## API仕様書(swagger)  
https://app.swaggerhub.com/apis/izana-42a/nebosuke-route_api/1.0.0  
※表示に時間がかかる場合があります。  
### /stations  
メソッド: GET   
パラメータ: なし  
レスポンス形式:  
```
[
  {
    "id": "urn:ucode:_00001C0000000000000100000341B4E7",
    "name": "渋谷"
  }, 
  ...
]
```
### /routesearch
メソッド: GET  
パラメータ:  
| 名前 | 随意性 | 説明 |
| ---- | --- | ---- |
| from | 必須 | 出発駅のNodeID(string) |
| to | 必須 | 到着駅のNodeID(string) |
| datetime | 必須 | 出発あるいは到着時刻(string) |
| isArrivalTime | 任意 | datetimeが到着時刻かどうか(bool)  ※デフォルトはfalse |
| via | 任意 | 経由駅のNodeID(string[]) |
レスポンス形式:  
```
{
  "steps": [
    {
      "arrival_station": "新川崎",
      "departure_station": "新橋",
      "from_time": "2025-05-28T10:01:00+09:00",
      "movetype": "local_train",
      "to_time": "2025-05-28T10:13:00+09:00"
    }
  ],
  "total_fare": 270,
  "total_time": 30
}
```
