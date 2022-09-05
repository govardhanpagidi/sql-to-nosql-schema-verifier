package models

type MigratorSchema struct {
	Id   string `json:"id"`
	Full struct {
		Databases struct {
			PerformanceSchema struct {
				Schemas struct {
					PerformanceSchema struct {
						Tables struct {
						} `json:"tables"`
					} `json:"performance_schema"`
				} `json:"schemas"`
			} `json:"performance_schema"`
			InformationSchema struct {
				Schemas struct {
					InformationSchema struct {
						Tables struct {
						} `json:"tables"`
					} `json:"information_schema"`
				} `json:"schemas"`
			} `json:"information_schema"`
			Sakila struct {
				Schemas struct {
					Sakila struct {
						Tables struct {
							FilmCategory struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									CategoryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"category_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									FilmId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"film_id"`
								} `json:"columns"`
							} `json:"film_category"`
							Country struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									Country struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"country"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									CountryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"country_id"`
								} `json:"columns"`
							} `json:"country"`
							Address struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									Address struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"address"`
									Address2 struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"address2"`
									Phone struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"phone"`
									District struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"district"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									AddressId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"address_id"`
									Location struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"location"`
									PostalCode struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"postal_code"`
									CityId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"city_id"`
								} `json:"columns"`
							} `json:"address"`
							City struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									City struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"city"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									CountryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"country_id"`
									CityId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"city_id"`
								} `json:"columns"`
							} `json:"city"`
							FilmActor struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									ActorId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"actor_id"`
									FilmId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"film_id"`
								} `json:"columns"`
							} `json:"film_actor"`
							Language struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									Name struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"name"`
									LanguageId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"language_id"`
								} `json:"columns"`
							} `json:"language"`
							Staff struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									StoreId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"store_id"`
									Password struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"password"`
									StaffId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"staff_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									AddressId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"address_id"`
									LastName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_name"`
									Active struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"active"`
									FirstName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"first_name"`
									Picture struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"picture"`
									Email struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"email"`
									Username struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"username"`
								} `json:"columns"`
							} `json:"staff"`
							Film struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									SpecialFeatures struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"special_features"`
									RentalDuration struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"rental_duration"`
									RentalRate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"rental_rate"`
									ReleaseYear struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"release_year"`
									Length struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"length"`
									ReplacementCost struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"replacement_cost"`
									Rating struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"rating"`
									Description struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"description"`
									LanguageId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"language_id"`
									Title struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"title"`
									OriginalLanguageId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"original_language_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									FilmId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"film_id"`
								} `json:"columns"`
							} `json:"film"`
							Store struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									StoreId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"store_id"`
									ManagerStaffId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"manager_staff_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									AddressId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"address_id"`
								} `json:"columns"`
							} `json:"store"`
							Inventory struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									StoreId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"store_id"`
									InventoryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"inventory_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									FilmId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"film_id"`
								} `json:"columns"`
							} `json:"inventory"`
							Rental struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									InventoryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"inventory_id"`
									StaffId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"staff_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									RentalDate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"rental_date"`
									CustomerId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"customer_id"`
									RentalId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"rental_id"`
									ReturnDate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"return_date"`
								} `json:"columns"`
							} `json:"rental"`
							Actor struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									LastName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_name"`
									ActorId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"actor_id"`
									FirstName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"first_name"`
								} `json:"columns"`
							} `json:"actor"`
							FilmText struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									Description struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"description"`
									FilmId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"film_id"`
									Title struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"title"`
								} `json:"columns"`
							} `json:"film_text"`
							Payment struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									Amount struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"amount"`
									PaymentId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"payment_id"`
									StaffId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"staff_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									CustomerId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"customer_id"`
									PaymentDate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"payment_date"`
									RentalId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"rental_id"`
								} `json:"columns"`
							} `json:"payment"`
							Category struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									CategoryId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"category_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									Name struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"name"`
								} `json:"columns"`
							} `json:"category"`
							Customer struct {
								Type    string `json:"type"`
								Size    int    `json:"size"`
								Columns struct {
									StoreId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"store_id"`
									LastUpdate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_update"`
									AddressId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int         `json:"ordinalPosition"`
										PrimaryKey      interface{} `json:"primaryKey"`
										ForeignKey      struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"address_id"`
									LastName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"last_name"`
									Active struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"active"`
									CustomerId struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition int `json:"ordinalPosition"`
										PrimaryKey      struct {
											Name        string      `json:"name"`
											Schema      interface{} `json:"schema"`
											Table       interface{} `json:"table"`
											Column      interface{} `json:"column"`
											Cardinality interface{} `json:"cardinality"`
										} `json:"primaryKey"`
										ForeignKey         interface{} `json:"foreignKey"`
										ReferencingColumns []struct {
											Name        string `json:"name"`
											Schema      string `json:"schema"`
											Table       string `json:"table"`
											Column      string `json:"column"`
											Cardinality string `json:"cardinality"`
										} `json:"referencingColumns"`
										Nullable        bool `json:"nullable"`
										AutoIncremented bool `json:"autoIncremented"`
										Generated       bool `json:"generated"`
										Hidden          bool `json:"hidden"`
										Indexed         bool `json:"indexed"`
										UniqueIndexed   bool `json:"uniqueIndexed"`
									} `json:"customer_id"`
									CreateDate struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"create_date"`
									FirstName struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"first_name"`
									Email struct {
										Type struct {
											DatabaseSpecificType string `json:"databaseSpecificType"`
											JavaSqlType          string `json:"javaSqlType"`
										} `json:"type"`
										OrdinalPosition    int           `json:"ordinalPosition"`
										PrimaryKey         interface{}   `json:"primaryKey"`
										ForeignKey         interface{}   `json:"foreignKey"`
										ReferencingColumns []interface{} `json:"referencingColumns"`
										Nullable           bool          `json:"nullable"`
										AutoIncremented    bool          `json:"autoIncremented"`
										Generated          bool          `json:"generated"`
										Hidden             bool          `json:"hidden"`
										Indexed            bool          `json:"indexed"`
										UniqueIndexed      bool          `json:"uniqueIndexed"`
									} `json:"email"`
								} `json:"columns"`
							} `json:"customer"`
						} `json:"tables"`
					} `json:"sakila"`
				} `json:"schemas"`
			} `json:"sakila"`
		} `json:"databases"`
	} `json:"full"`
	Imported struct {
		Sakila struct {
			Sakila struct {
				FilmCategory struct {
				} `json:"film_category"`
				Country struct {
				} `json:"country"`
				Address struct {
				} `json:"address"`
				City struct {
				} `json:"city"`
				FilmActor struct {
				} `json:"film_actor"`
				Language struct {
				} `json:"language"`
				Staff struct {
				} `json:"staff"`
				Film struct {
				} `json:"film"`
				Store struct {
				} `json:"store"`
				Inventory struct {
				} `json:"inventory"`
				Rental struct {
				} `json:"rental"`
				Actor struct {
				} `json:"actor"`
				FilmText struct {
				} `json:"film_text"`
				Payment struct {
				} `json:"payment"`
				Category struct {
				} `json:"category"`
				Customer struct {
				} `json:"customer"`
			} `json:"sakila"`
		} `json:"sakila"`
	} `json:"imported"`
}
