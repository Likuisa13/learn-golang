package entity

type TodoListCategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name,omitempty" validate:"required" name:"Nama"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func (r *TodoListCategoryReq) SetID(ID int64) {
	r.ID = ID
}
