
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/apache/trafficcontrol/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Domains struct {
	Name                   string       `db:"name" json:"name"`
	DomainCatalog          null.String  `db:"domain_catalog" json:"domainCatalog"`
	DomainSchema           null.String  `db:"domain_schema" json:"domainSchema"`
	Dnssec                 bool         `db:"dnssec" json:"dnssec"`
	DomainName             null.String  `db:"domain_name" json:"domainName"`
	DataType               null.String  `db:"data_type" json:"dataType"`
	CreatedAt              time.Time    `db:"created_at" json:"createdAt"`
	CharacterMaximumLength null.Int     `db:"character_maximum_length" json:"characterMaximumLength"`
	CharacterOctetLength   null.Int     `db:"character_octet_length" json:"characterOctetLength"`
	CharacterSetCatalog    null.String  `db:"character_set_catalog" json:"characterSetCatalog"`
	CharacterSetSchema     null.String  `db:"character_set_schema" json:"characterSetSchema"`
	CharacterSetName       null.String  `db:"character_set_name" json:"characterSetName"`
	CollationCatalog       null.String  `db:"collation_catalog" json:"collationCatalog"`
	CollationSchema        null.String  `db:"collation_schema" json:"collationSchema"`
	CollationName          null.String  `db:"collation_name" json:"collationName"`
	NumericPrecision       null.Int     `db:"numeric_precision" json:"numericPrecision"`
	NumericPrecisionRadix  null.Int     `db:"numeric_precision_radix" json:"numericPrecisionRadix"`
	NumericScale           null.Int     `db:"numeric_scale" json:"numericScale"`
	DatetimePrecision      null.Int     `db:"datetime_precision" json:"datetimePrecision"`
	IntervalType           null.String  `db:"interval_type" json:"intervalType"`
	IntervalPrecision      null.Int     `db:"interval_precision" json:"intervalPrecision"`
	DomainDefault          null.String  `db:"domain_default" json:"domainDefault"`
	UdtCatalog             null.String  `db:"udt_catalog" json:"udtCatalog"`
	UdtSchema              null.String  `db:"udt_schema" json:"udtSchema"`
	UdtName                null.String  `db:"udt_name" json:"udtName"`
	ScopeCatalog           null.String  `db:"scope_catalog" json:"scopeCatalog"`
	ScopeSchema            null.String  `db:"scope_schema" json:"scopeSchema"`
	ScopeName              null.String  `db:"scope_name" json:"scopeName"`
	MaximumCardinality     null.Int     `db:"maximum_cardinality" json:"maximumCardinality"`
	DtdIdentifier          null.String  `db:"dtd_identifier" json:"dtdIdentifier"`
	Links                  DomainsLinks `json:"_links" db:-`
}

type DomainsLinks struct {
	Self     string   `db:"self" json:"_self"`
	CdnsLink CdnsLink `json:"cdns" db:-`
}

type DomainsLink struct {
	ID  string `db:"domain" json:"name"`
	Ref string `db:"domains_name_ref" json:"_ref"`
}

// @Title getDomainsById
// @Description retrieves the domains information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Domains
// @Resource /api/2.0
// @Router /api/2.0/domains/{id} [get]
func getDomain(name string, db *sqlx.DB) (interface{}, error) {
	ret := []Domains{}
	arg := Domains{}
	arg.Name = name
	queryStr := "select *, concat('" + API_PATH + "domains/', name) as self"
	queryStr += ", concat('" + API_PATH + "cdns/', cdn) as cdns_name_ref"
	queryStr += " from domains WHERE name=:name"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getDomainss
// @Description retrieves the domains
// @Accept  application/json
// @Success 200 {array}    Domains
// @Resource /api/2.0
// @Router /api/2.0/domains [get]
func getDomains(db *sqlx.DB) (interface{}, error) {
	ret := []Domains{}
	queryStr := "select *, concat('" + API_PATH + "domains/', name) as self"
	queryStr += ", concat('" + API_PATH + "cdns/', cdn) as cdns_name_ref"
	queryStr += " from domains"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postDomains
// @Description enter a new domains
// @Accept  application/json
// @Param                 Body body     Domains   true "Domains object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/domains [post]
func postDomain(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Domains
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO domains("
	sqlString += "name"
	sqlString += ",domain_catalog"
	sqlString += ",domain_schema"
	sqlString += ",cdn"
	sqlString += ",dnssec"
	sqlString += ",domain_name"
	sqlString += ",data_type"
	sqlString += ",created_at"
	sqlString += ",character_maximum_length"
	sqlString += ",character_octet_length"
	sqlString += ",character_set_catalog"
	sqlString += ",character_set_schema"
	sqlString += ",character_set_name"
	sqlString += ",collation_catalog"
	sqlString += ",collation_schema"
	sqlString += ",collation_name"
	sqlString += ",numeric_precision"
	sqlString += ",numeric_precision_radix"
	sqlString += ",numeric_scale"
	sqlString += ",datetime_precision"
	sqlString += ",interval_type"
	sqlString += ",interval_precision"
	sqlString += ",domain_default"
	sqlString += ",udt_catalog"
	sqlString += ",udt_schema"
	sqlString += ",udt_name"
	sqlString += ",scope_catalog"
	sqlString += ",scope_schema"
	sqlString += ",scope_name"
	sqlString += ",maximum_cardinality"
	sqlString += ",dtd_identifier"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:domain_catalog"
	sqlString += ",:domain_schema"
	sqlString += ",:cdn"
	sqlString += ",:dnssec"
	sqlString += ",:domain_name"
	sqlString += ",:data_type"
	sqlString += ",:created_at"
	sqlString += ",:character_maximum_length"
	sqlString += ",:character_octet_length"
	sqlString += ",:character_set_catalog"
	sqlString += ",:character_set_schema"
	sqlString += ",:character_set_name"
	sqlString += ",:collation_catalog"
	sqlString += ",:collation_schema"
	sqlString += ",:collation_name"
	sqlString += ",:numeric_precision"
	sqlString += ",:numeric_precision_radix"
	sqlString += ",:numeric_scale"
	sqlString += ",:datetime_precision"
	sqlString += ",:interval_type"
	sqlString += ",:interval_precision"
	sqlString += ",:domain_default"
	sqlString += ",:udt_catalog"
	sqlString += ",:udt_schema"
	sqlString += ",:udt_name"
	sqlString += ",:scope_catalog"
	sqlString += ",:scope_schema"
	sqlString += ",:scope_name"
	sqlString += ",:maximum_cardinality"
	sqlString += ",:dtd_identifier"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putDomains
// @Description modify an existing domainsentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Domains   true "Domains object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/domains/{id}  [put]
func putDomain(name string, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg Domains
	err := json.Unmarshal(payload, &arg)
	arg.Name = name
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE domains SET "
	sqlString += "name = :name"
	sqlString += ",domain_catalog = :domain_catalog"
	sqlString += ",domain_schema = :domain_schema"
	sqlString += ",cdn = :cdn"
	sqlString += ",dnssec = :dnssec"
	sqlString += ",domain_name = :domain_name"
	sqlString += ",data_type = :data_type"
	sqlString += ",created_at = :created_at"
	sqlString += ",character_maximum_length = :character_maximum_length"
	sqlString += ",character_octet_length = :character_octet_length"
	sqlString += ",character_set_catalog = :character_set_catalog"
	sqlString += ",character_set_schema = :character_set_schema"
	sqlString += ",character_set_name = :character_set_name"
	sqlString += ",collation_catalog = :collation_catalog"
	sqlString += ",collation_schema = :collation_schema"
	sqlString += ",collation_name = :collation_name"
	sqlString += ",numeric_precision = :numeric_precision"
	sqlString += ",numeric_precision_radix = :numeric_precision_radix"
	sqlString += ",numeric_scale = :numeric_scale"
	sqlString += ",datetime_precision = :datetime_precision"
	sqlString += ",interval_type = :interval_type"
	sqlString += ",interval_precision = :interval_precision"
	sqlString += ",domain_default = :domain_default"
	sqlString += ",udt_catalog = :udt_catalog"
	sqlString += ",udt_schema = :udt_schema"
	sqlString += ",udt_name = :udt_name"
	sqlString += ",scope_catalog = :scope_catalog"
	sqlString += ",scope_schema = :scope_schema"
	sqlString += ",scope_name = :scope_name"
	sqlString += ",maximum_cardinality = :maximum_cardinality"
	sqlString += ",dtd_identifier = :dtd_identifier"
	sqlString += " WHERE name=:name"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delDomainsById
// @Description deletes domains information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Domains
// @Resource /api/2.0
// @Router /api/2.0/domains/{id} [delete]
func delDomain(name string, db *sqlx.DB) (interface{}, error) {
	arg := Domains{}
	arg.Name = name
	result, err := db.NamedExec("DELETE FROM domains WHERE name=:name", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
