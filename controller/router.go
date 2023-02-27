package controller

import (
    "net/http"
)

// 外部パッケージに公開するインターフェース
type Router interface {
    HandleCardsRequest(w  http.ResponseWriter, r *http.Request)
}

// 非公開のRouter構造体
type router struct {
    hc HandController
}

// Routerのコンストラクタ。引数にCardControllerを受け取り、Router構造体のポインタを返却する。
func NewRouter(hc HandController) Router {
    return &router{hc}
}

func (ro *router) HandleCardsRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            ro.hc.GetCards(w, r)
    }
}

