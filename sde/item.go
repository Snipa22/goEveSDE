package sde

import (
	"context"
	"fmt"
)

type Item struct {
	ItemID        int
	Name          string
	Description   string
	MetaLevel     int8
	IconID        *int
	GraphicID     int
	GroupID       int
	MarketGroupID int
	Mass          float64
	Volume        float64
	Capacity      float64
	Produced      int
}

func GetItemByID(itemID int) (*Item, error) {
	item, ok := getItemFromCache(fmt.Sprint(itemID))
	if ok {
		retItem := item.(Item)
		return &retItem, nil
	}
	c := mustGetPoolConn()
	defer c.Release()
	rows, err := c.Query(context.Background(), "select groupID, typeName, description, mass, volume, capacity, portionSize, marketGroupID, iconID, graphicID from invTypes where typeID = ?", itemID)
	if err != nil {
		return nil, err
	}
	retItem := Item{}
	if rows.Next() {
		if err = rows.Scan(&retItem.GroupID, &retItem.Name, &retItem.Description, &retItem.Mass, &retItem.Volume, &retItem.Capacity, &retItem.Produced, &retItem.MarketGroupID, &retItem.IconID, &retItem.GraphicID); err != nil {
			return nil, err
		}
	}
	rows.Close()
	if retItem.ItemID == 0 {
		return nil, ItemNotFound
	}
	setItemInCache(fmt.Sprint(itemID), retItem)
	return &retItem, nil
}

func GetItemByName(itemName string) (*Item, error) {
	item, ok := getItemFromCache(itemName)
	if ok {
		retItem := item.(Item)
		return &retItem, nil
	}
	c := mustGetPoolConn()
	defer c.Release()
	rows, err := c.Query(context.Background(), "select groupID, typeName, description, mass, volume, capacity, portionSize, marketGroupID, iconID, graphicID from invTypes where typeName ilike ? limit 1 order by id desc", fmt.Sprintf("%%%v%%", itemName))
	if err != nil {
		return nil, err
	}
	retItem := Item{}
	if rows.Next() {
		if err = rows.Scan(&retItem.GroupID, &retItem.Name, &retItem.Description, &retItem.Mass, &retItem.Volume, &retItem.Capacity, &retItem.Produced, &retItem.MarketGroupID, &retItem.IconID, &retItem.GraphicID); err != nil {
			return nil, err
		}
	}
	rows.Close()
	if retItem.ItemID == 0 {
		return nil, ItemNotFound
	}
	setItemInCache(itemName, retItem)
	return &retItem, nil
}
