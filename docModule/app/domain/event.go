package domain

import "github.com/google/uuid"

type events struct {
	EventUuid      uuid.UUID
	DocUuid        uuid.UUID
	EType          string    //тип события
	ECode          string    //код события
	EUser          string    //Пользователь
	EIp            string    //IP пользователя вызвавшего событие
	ENum           int64     //порядковый номер события
	EPrevEventUuid uuid.UUID //Гуид предшествующего события
	EResult        []eventResult
	ETransaction   []transaction
}

type eventResult struct {
	ERType        string
	ERCode        string
	ERDescription string
}
