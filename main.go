package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"

    ct "github.com/daviddengcn/go-colortext"
)

type BedWarsRecord struct {
    MvpNum    int     `json:"mvpNum"`
    PlayNum   int     `json:"playNum"`
    WinRate   float64 `json:"winRate"`
    BeddesNum int     `json:"beddesNum"`
    KillDead  float64 `json:"killDead"`
}

type SkyWarsRecord struct {
    KillNum float64 `json:"killNum"`
    PlayNum int     `json:"playNum"`
    WinRate float64 `json:"winRate"`
}

func main() {
    for {
        fmt.Print("请输入玩家名称来查询战绩：")
        var id = "两秒十七发"
        fmt.Scanln(&id)
        encodeId := url.QueryEscape(id)
        BedWarsQuestUrl := "http://mc-api.16163.com/search/bedwars.html?uid=" + encodeId
        fmt.Println("正在发送请求到 " + BedWarsQuestUrl)
        getRecord(BedWarsQuestUrl, "起床战争", id)
        fmt.Println()
        SkyWarsQuestUrl := "http://mc-api.16163.com/search/skywars.html?uid=" + encodeId
        fmt.Println("正在发送请求到 " + SkyWarsQuestUrl)
        getRecord(SkyWarsQuestUrl, "空岛战争", id)
        fmt.Println()
    }
}

func getRecord(questUrl string, gameType string, id string) {
    resp, err := http.Get(questUrl)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusNotFound {
        b, _ := io.ReadAll(resp.Body)
        ct.Foreground(ct.Green, true)
        fmt.Println("查询成功，玩家 " + id + " 的 " + gameType + " 战绩如下：")
        ct.ResetColor()

        switch gameType {
        case "起床战争":
            var record BedWarsRecord
            if err := json.Unmarshal(b, &record); err != nil {
                log.Fatal("JSON解析失败:", err)
            }
            fmt.Printf("总场次：%d\n", record.PlayNum)
            winNum := (float64)(record.PlayNum) * record.WinRate
            fmt.Printf("胜场/胜率：%d/%d%%\n", (int)(winNum), int(record.WinRate*100))
            mvpRate := float64(record.MvpNum) / (winNum) * 100
            fmt.Printf("MVP次数/MVP率：%d/%.2f%%\n", record.MvpNum, mvpRate)
            fmt.Printf("K/D：%.2f\n", record.KillDead)
            fmt.Printf("破坏的床的数量：%d\n", record.BeddesNum)
        case "空岛战争":
            var record SkyWarsRecord
            if err := json.Unmarshal(b, &record); err != nil {
                log.Fatal("JSON解析失败:", err)
            }
            fmt.Printf("总场次：%d\n", record.PlayNum)
            fmt.Printf("胜率：%.2f\n", record.WinRate)
            fmt.Printf("击杀数：%.2f\n", record.KillNum)
        }
    } else {
        ct.Foreground(ct.Red, true)
        fmt.Println("未查询到战绩，此 ID 可能是新玩家")
        ct.ResetColor()
    }
}
