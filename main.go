package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"

    ct "github.com/daviddengcn/go-colortext"
)

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

func getRecord(questUrl string, type1 string, id string) {
    resp, err := http.Get(questUrl)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusNotFound {
        b, _ := io.ReadAll(resp.Body)
        ct.Foreground(ct.Green, true)
        fmt.Println("查询成功，玩家 " + id + " 的 " + type1 + " 战绩如下：")
        ct.ResetColor()
        var formattedJSON bytes.Buffer
        err := json.Indent(&formattedJSON, b, "", "  ")
        if err != nil {
            log.Fatal("JSON格式化失败:", err)
        }
        fmt.Println(formattedJSON.String())
    } else {
        ct.Foreground(ct.Red, true)
        fmt.Println("未查询到战绩，此 ID 可能是新玩家")
        ct.ResetColor()
    }
}
