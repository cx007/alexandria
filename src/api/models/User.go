/*
 * Alexandria CMDB - Open source configuration management database
 * Copyright (C) 2014  Ryan Armstrong <ryan@cavaliercoder.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * package controllers
 */
package models

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	model `json:"-" bson:",inline"`

	TenantId  bson.ObjectId `json:"-"`
	ApiKey    string        `json:"apiKey"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Email     string        `json:"email" binding:"required"`
}

func (c *User) Init() {
	c.SetCreated()
	c.GenerateApiKey()
}

func (c *User) GenerateApiKey() string {
	// Create API key
	jsonHash, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	shaSum := sha1.Sum(jsonHash)
	c.ApiKey = fmt.Sprintf("%x", shaSum)

	return c.ApiKey
}
