// Copyright 2024 The Gaea Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

type DBType uint8

const (
	DBTypeMySQL      DBType = 0x01
	DBTypeOceanBase  DBType = 0x02
	DBTypePostgreSQL DBType = 0x03
)

func (d DBType) String() string {
	switch d {
	case DBTypeMySQL:
		return "mysql"
	case DBTypeOceanBase:
		return "oceanbase"
	case DBTypePostgreSQL:
		return "postgresql"
	default:
		return "unknown"
	}
}

func ParseDBType(s string) (DBType, bool) {
	switch s {
	case "mysql", "MySQL":
		return DBTypeMySQL, true
	case "oceanbase", "OceanBase", "oceanbase/mysql":
		return DBTypeOceanBase, true
	case "postgresql", "PostgreSQL", "postgres":
		return DBTypePostgreSQL, true
	default:
		return DBTypeMySQL, false
	}
}
