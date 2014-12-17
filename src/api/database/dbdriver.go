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
package database

import (
	"alexandria/api/configuration"
	"alexandria/api/database/mongodb"
	"fmt"
)

type DbDriver interface {
	IsBootStrapped() (bool, error)
	BootStrap(*configuration.Answers) error
}

var dbsession DbDriver

func Connect() (DbDriver, error) {
	if dbsession == nil {
		config, err := configuration.GetConfig()
		if err != nil {
			return nil, err
		}

		switch config.Database.Driver {
		case "mongodb":
			// Connect to database
			dbsession, err = mongodb.GetConnection()
			if err != nil {
				return nil, err
			}

		default:
			panic(fmt.Sprintf("Unsupported database driver: %s", config.Database.Driver))
		}
	}
	return dbsession, nil
}