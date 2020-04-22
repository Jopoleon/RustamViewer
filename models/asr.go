package models

type ASR struct {
	ID             *int    `json:"id" db:"id"`
	ExternalID     *int    `json:"external_id" db:"external_id"`
	MenuName       *string `json:"menu_name" db:"menu_name"`
	ProjectName    *string `json:"project_id" db:"project_id"`
	Ani            int     `json:"ani" db:"ani"`
	CallID         *string `json:"callid" db:"callid"`
	Seq            int     `json:"seq" db:"seq"`
	Utterance      *string `json:"utterance" db:"utterance"`
	Interpretation *string `json:"interpretation" db:"interpretation"`
	Confidence     float64 `json:"confidence" db:"confidence"`
	Inputmode      *string `json:"inputmode" db:"inputmode"`
	Grammaruri     *string `json:"grammaruri" db:"grammaruri"`
	Waverecord     []byte  `json:"waverecord" db:"waverecord" csv:"-"`
	CreatedOn      MyTime  `json:"created_on" db:"created_on"`
	AssignedClass  *string `json:"assignedclass" db:"assignedclass"`
}
