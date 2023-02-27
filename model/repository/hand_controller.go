packeage repository

import (
    "log"

    "github.com/tkanata/go-training-v2/model/entity"
)

// 外部パッケージに公開するインターフェース
type HandRepository interface {
    GetCards() (cards []entity.HandEntity, err error)
}

// 非公開のHandRepostory
type handRepository struct {
}


// HandRepositoryのコンストラクタ。HandRepository構造体のポインタを返却する。
func NewHandRepository() HandRepository {
    return &handRepository{}
}

// Card取得処理
func (hr *handRepository) GetCards() (cards []entity.HandEntity, err error) {
    cards = []entity.HandEntity{}

    cands := entity.HandEntity{}
    
    // どこかからcardを取得する処理を書く。
    hogehoge

    return "テスト(repository/hand_controller.go)"
}



