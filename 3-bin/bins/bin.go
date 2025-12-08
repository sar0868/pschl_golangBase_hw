package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (bin *Bin) ToString() string {
	return fmt.Sprintf("id: %s\nprivate: %t\nname: %s\ncreated: %v", bin.Id, bin.Private, bin.Name, bin.CreatedAt)
}
