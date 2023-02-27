package controller

import (
    "encoding/json"
    "net/http"
    "path"
    "strconv"

    "github.com/tkanata/go-training-v2/model/entity"
    "github.com/tkanata/go-training-v2/model/repository"
)

// 外部パッケージに公開するインターフェース
type HandController interface {
    GetCards(w http.ResponseWriter, r *http.Request)
    PostHand(w http.ResponseWriter, r *http.Request)
    PutHand(w http.ResponseWriter, r *http.Request)
    DeleteHand(w http.ResponseWriter, r *http.Request)
}

// 非公開のPokerController構造体
type handController struct {
    cr repository.HandRepository
}

// HandControllerのコンストラクタ。
// 引数にHandRepositoryを受け取り、HandController構造体のポインタを返却する。
func NewHandController(hr repository.HandRepository) HandController {
    return & handController{hr}
}

// cardの取得
func (hc *handController) GetCards(w http.ResponseWriter, r *http.Request) {
    // リポジトリの取得処理呼び出し
    cards, err := hc.hr.GetCards()
    if err != nil {
        w.WriterHeader(500)
        return
    }

    // jsonに変換
    output, _ := json.MarchalIndent(handResponse.Cards, "", "\t\t")

    // jsonを返却
    w.Header().Set("Content-Type", "application/json")
    w.Writer(output)
}





