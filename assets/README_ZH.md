# 後端系統 - 廣告獲取與投遞
## 需求分析
### 人物狀態
1. 廣告投遞商
   ```
   為廣告的提供者，可以自行設計希望投遞的目標客群各種細節
   不過不會直接讓廣告商使用內部API，因為廣告宜先經過審核再由內部人員將廣告放置系統內
   ```
2. 管理者
   ```
   為系統管理者，預想狀態為廣告投遞商與管理者進行討論，
   確認投遞廣告項目與目標受眾後，將該項廣告資訊加入系統當中
   ```
3. 廣告接收者
   ```
   為一般使用人員，預想狀態透過查詢獲取其相關的廣告內容
   ```

### 需求項目
1. 新增廣告
   * ⽤於產⽣管理廣告資源 (僅為CREATE)
   * 使⽤ RESTful
   * 廣告屬性
     * 標題 Title
     * 開始與結束時間 StartAt、EdnAt
     * 年齡、性別、國家、平台 (Age、Gender、Country、Platform) (均為多選)
2. 查詢廣告
   * 列出符合條件的活躍廣告 (StartAt < NOW < EndAt)
   * 使⽤ RESTful
   * 具分頁功能 透過 offset & limit
   * 查詢參數 : Age、Gender、Country、Platform

```
備註 :
    RESTful 為一種API設計風格(路由即資源) 而通常使指http伺服器。
    因此該專案使用http作為整體後端的溝通協議。
```
```
備註 :
    在新增廣告環節當中，若是標題、開始時間、結束時間等等資訊都相同，是否該視為同一筆廣告。
    這項資訊關連到後續的資料庫設計。
    雖然題意無特別強調，不過我的答案是否定。我認為可能出現恰好資訊完全相同不同公司的狀態。
    為此呼應我後續的資料庫設計。
```

### 效能規範
Q、提供 Public API 能超過 10,000 Requests Per Second 的設計

首先對於每秒10000則請求，相信當中**勢必會出現續多相同條件請求**，
因此系統宜採用**快取服務以達到同時更多相同的請求**，當中快取位置可放於反向代理伺服器或是後端系統。

再來假設已經對於資料獲取方式做優化仍無法達到要求，則應採取"**垂直擴展**"或是"**水平擴展**"。
1. 垂直擴展 則需要行能更強大的主機，於軟體設計較無法控制 
2. 水平擴展 則需要使用分散式技術，可以使用kubernetes的方式進行部屬，並連接多台主機以提升效能

最後對於系統使否能承受指定流量仍需要做測試採能知曉，所以必須撰寫其相關腳本，達到流量要求並且產生隨機的詢問內容。
(此項專案使用我個人熟悉的K6作為負載測試)

### 其餘規範

* 可以參考 API 範例，也能⾃⾏設計 API
  * 會基於題目所規範的API進行設計
* 如果需要，請隨意使⽤任何外部函式庫
  * 基於安全考量與各項題目要求會使用gin作為伺服器框架，同時會使用多種連接各項資料庫的相關套件。
* ⾃由選擇外部儲存⽅式
  * 對於各項資訊儲存會採用不同方式
* 請對這兩組 API 進⾏合理的參數驗證和錯誤處理
  * 為此在middleware時會建立其相關檢查機制，並對其做相關處理。
* 請撰寫適當的 test
  * 為此會新增test資料夾，做單元測試、API測試、各項負載測試
* 不需要考慮⾝份驗證
  * 好的。表示所有人均可透過API進行互動。
* 同時存在系統的總活躍廣告數量 (也就是 StartAt < NOW < EndAt) < 1000
  * 可以在投遞廣告時進行檢查，若是該時段數量已超過1000則拒絕該廣告投遞請求
* 每天 create 的廣告數量 不會超過 3000 個
  * 為此可以為每個廣告建立created_at屬性，用以檢查當日投遞廣告數量，同時為提高效能可以將今日以投遞數量放置快取當中，整體規畫路徑設計如下，
  1. 進來後先於快取檢查當日投遞廣告數量
  2. 進入資料庫 再次檢查今日投遞數量
  3. 若已超過3000則廣告，則拒絕該項請求，並同時更改快取資訊數量為資料庫紀錄數值
  4. 若低於3000則廣告，則新增該項請求內容，並同時更改快取資訊數量為資料庫紀錄數值
* 請在 Readme 中描述想法和設計上的選擇
  * 好的。

### 附加功能

Q、對於達到題目要求的設計

對於題目規範，若欲達成最合適設計，最佳解法為設計實驗並觀察其狀態。

為此必須要建立其觀測系統以記錄並觀察數據，因此會於外增加Grafana、Prometheus、Influxdb等等工具，以觀測實驗成效。

## API 設計

由於是採用RESTful設計風格進行設計，因此使用OPENAPI(舊為Swagger)作為API規範文件非常合適。

[API文檔內容](../docs/api_specification.yaml)

```
備註 :
    此次開發採用文檔先行的方式，先設計文檔再撰寫後端程序內容。
```