package models

import "time"

const NEW_DOCUMENT = "NEW_DOCUMENT"
const EMBEDDED_DOCUMENT_ARRAY = "EMBEDDED_DOCUMENT_ARRAY"
const EMBEDDED_DOCUMENT = "EMBEDDED_DOCUMENT"

type MigrationSchema struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	LastModified time.Time `json:"lastModified"`
	SchemasId    string    `json:"schemasId"`
	Content      struct {
		Settings struct {
			Casing      interface{} `json:"casing"`
			KeyHandling string      `json:"keyHandling"`
		} `json:"settings"`
		Collections   map[string]Collection `json:"collections"`
		Mappings      map[string]Mapping    `json:"mappings"`
		Relationships struct {
			Tables      map[string]Mappings `json:"tables"`
			Collections map[string]Mappings `json:"collections"`
			Mappings    map[string]Children `json:"mappings"`
		} `json:"relationships"`
		Diagrams struct {
			Tabs []struct {
				Relational struct {
					Nodes []struct {
						Id       string `json:"id"`
						Type     string `json:"type"`
						Position struct {
							X float64 `json:"x"`
							Y float64 `json:"y"`
						} `json:"position"`
						TargetPosition interface{} `json:"targetPosition"`
						SourcePosition interface{} `json:"sourcePosition"`
						Width          int         `json:"width"`
						Height         int         `json:"height"`
					} `json:"nodes"`
					Edges []struct {
						Id          string `json:"id"`
						Source      string `json:"source"`
						Target      string `json:"target"`
						MarkerStart string `json:"markerStart"`
						MarkerEnd   string `json:"markerEnd"`
					} `json:"edges"`
				} `json:"relational"`
				Collection struct {
					Nodes []struct {
						Id       string `json:"id"`
						Type     string `json:"type"`
						Position struct {
							X float64 `json:"x"`
							Y float64 `json:"y"`
						} `json:"position"`
						TargetPosition interface{} `json:"targetPosition"`
						SourcePosition interface{} `json:"sourcePosition"`
						Width          int         `json:"width"`
						Height         int         `json:"height"`
					} `json:"nodes"`
					Edges []struct {
						Id          string `json:"id"`
						Source      string `json:"source"`
						Target      string `json:"target"`
						MarkerStart string `json:"markerStart"`
						MarkerEnd   string `json:"markerEnd"`
					} `json:"edges"`
				} `json:"collection"`
			} `json:"tabs"`
		} `json:"diagrams"`
	} `json:"content"`
	ConnectionDetails Config `json:"connectionDetails"`
}

type Config struct {
	Jdbc struct {
		Type     string `json:"type"`
		Url      string `json:"url"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"jdbc"`
	Mongodb struct {
		ConnectionString string `json:"connectionString"`
	} `json:"mongodb"`
}

type Target struct {
	Name              string      `json:"name"`
	Included          bool        `json:"included"`
	Type              string      `json:"type"`
	MappedParentField interface{} `json:"mappedParentField"`
}
type Source struct {
	Name                 string `json:"name"`
	JavaSqlType          string `json:"javaSqlType"`
	DatabaseSpecificType string `json:"databaseSpecificType"`
	IsPrimaryKey         bool   `json:"isPrimaryKey"`
}

type Settings struct {
	Type         string      `json:"type"`
	Notes        interface{} `json:"notes"`
	EmbeddedPath string      `json:"embeddedPath"`
}

type Mapping struct {
	Settings     Settings         `json:"settings"`
	Fields       map[string]Filed `json:"fields"`
	CollectionId string           `json:"collectionId"`
	Table        string           `json:"table"`
}

type Mappings struct {
	Ids []string `json:"mappings"`
}

type Children struct {
	Children []string `json:"children"`
}

type Collection struct {
	Name string `json:"name"`
}

type RecordData struct {
	TableName      string
	RowCount       int64
	CollectionName string
	DocCount       int64
}

type Filed struct {
	Target struct {
		Name              string      `json:"name"`
		Included          bool        `json:"included"`
		Type              string      `json:"type"`
		MappedParentField interface{} `json:"mappedParentField"`
	} `json:"target"`
	Source struct {
		Name                 string `json:"name"`
		DatabaseSpecificType string `json:"databaseSpecificType"`
		IsPrimaryKey         bool   `json:"isPrimaryKey"`
	} `json:"source"`
}
