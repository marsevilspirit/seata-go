/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testdata

import "seata.apache.org/seata-go/pkg/datasource/sql/types"

func MockWantTypesMeta(tableName string) types.TableMeta {
	return types.TableMeta{
		TableName: tableName,
		Columns: map[string]types.ColumnMeta{
			"id": {
				ColumnName: "id",
			},
			"name": {
				ColumnName: "name",
			},
		},
		Indexs: map[string]types.IndexMeta{
			"": {
				ColumnName: "id",
				IType:      types.IndexTypePrimaryKey,
				Columns: []types.ColumnMeta{
					{
						ColumnName:   "id",
						DatabaseType: types.GetSqlDataType("BIGINT"),
					},
					{
						ColumnName: "id",
					},
				},
			},
		},
		ColumnNames: []string{
			"id",
			"name",
		},
	}
}
