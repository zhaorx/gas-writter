package models

type PointInfo struct {
	Point       string `db:"POINT"`
	OriginPoint string `db:"ORIGIN_POINT"`
	Pname       string `db:"PNAME"`
	Region      string `db:"REGION"`
	Unit        string `db:"UNIT"`
	Gases       string `db:"GASES"`
	Gas         string `db:"GAS"`
	Site        string `db:"SITE"`
	Pipeline    string `db:"PIPELINE"`
	Uptype      string `db:"UPTYPE"`
	Ptype       string `db:"PTYPE"`
	Calc        string `db:"CALC"`
}
