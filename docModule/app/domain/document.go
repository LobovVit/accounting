package domain

import (
	"github.com/google/uuid"
	"time"
)

//Карточка документа для учета
type documentCard struct {
	DocUuid            uuid.UUID //гуид документа
	Source             string    //Источник документа (подсистема)
	Num                string    //номер документа
	Date               time.Time //дата документа
	SenderTofk         string    //Тофк отправитель
	ReceiverTofk       string    //Тофк получатель
	SenderPa           string    //ЛС отправитель
	ReceiverPa         string    //ЛС получатель
	SenderOrg          string    //Ортганизация отправитель
	ReceiverOrg        string    //Организация получатель
	SenderAccount      string    //казначейский счет отправитель
	ReceiverAccount    string    //казначейский счет получатель
	SenderAttribute1   string    //Attribute1 отправитель
	ReceiverAttribute1 string    //Attribute1 получатель
	SenderAttribute2   string    //Attribute2 отправитель
	ReceiverAttribute2 string    //Attribute2 получатель
	SenderAttribute3   string    //Attribute3 отправитель
	ReceiverAttribute3 string    //Attribute3 получатель
	SenderAttribute4   string    //Attribute4 отправитель
	ReceiverAttribute4 string    //Attribute4 получатель
	SenderAttribute5   string    //Attribute5 отправитель
	ReceiverAttribute5 string    //Attribute5 получатель
	CreateDate         time.Time //дата время создания
	CreateUser         string    //логин создателя
	CreateIp           string    //ip создателя
	LastUpdateDate     time.Time //дата время последнего изменения
	LastUpdateUser     string    //логин последнего измененившего
	LastUpdateIp       string    //ip последнего изменившего
}

type headerAttribute struct {
	HStrAttr01  string
	HStrAttr02  string
	HStrAttr03  string
	HStrAttr04  string
	HStrAttr05  string
	HStrAttr06  string
	HStrAttr07  string
	HStrAttr08  string
	HStrAttr09  string
	HStrAttr10  string
	HNumAttr01  int64
	HNumAttr02  int64
	HNumAttr03  int64
	HNumAttr04  int64
	HNumAttr05  int64
	HNumAttr06  int64
	HNumAttr07  int64
	HNumAttr08  int64
	HNumAttr09  int64
	HNumAttr10  int64
	HDateAttr01 time.Time
	HDateAttr02 time.Time
	HDateAttr03 time.Time
	HDateAttr04 time.Time
	HDateAttr05 time.Time
}

type line struct {
	LBaseNum         string    //номер документа основания
	LBaseDate        time.Time //дата документа основания
	LAccountedAmount int64
	LEnteredAmount   int64
	LCurrencyCode    string
	//LCurrencyRate int64
	LAccountedAmount1 int64
	LEnteredAmount1   int64
	LAccountedAmount2 int64
	LEnteredAmount2   int64
	LAccountedAmount3 int64
	LEnteredAmount3   int64
	LAccountedAmount4 int64
	LEnteredAmount4   int64
	LAccountedAmount5 int64
	LEnteredAmount5   int64
}

type lineAttribute struct {
	LStrAttr01  string
	LStrAttr02  string
	LStrAttr03  string
	LStrAttr04  string
	LStrAttr05  string
	LStrAttr06  string
	LStrAttr07  string
	LStrAttr08  string
	LStrAttr09  string
	LStrAttr10  string
	LNumAttr01  int64
	LNumAttr02  int64
	LNumAttr03  int64
	LNumAttr04  int64
	LNumAttr05  int64
	LNumAttr06  int64
	LNumAttr07  int64
	LNumAttr08  int64
	LNumAttr09  int64
	LNumAttr10  int64
	LDateAttr01 time.Time
	LDateAttr02 time.Time
	LDateAttr03 time.Time
	LDateAttr04 time.Time
	LDateAttr05 time.Time
}
