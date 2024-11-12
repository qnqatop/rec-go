// Code generated by mfd-generator v0.4.5; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

var Columns = struct {
	City struct {
		ID, Name string
	}
	Direction struct {
		ID, UniversityID, CityID, Code, Name, Params, Cost string

		University, City string
	}
	DirectionsFeedback struct {
		UserID, DirectionID, Rating string

		User, Direction string
	}
	UniversityFeedback struct {
		Rating, UserID, UniversityID string

		User, University string
	}
	Universyty struct {
		ID, Name, CityID string

		City string
	}
	User struct {
		ID, CityID, UniversityID, Gender, Age, Params string

		City, University string
	}
}{
	City: struct {
		ID, Name string
	}{
		ID:   "cityID",
		Name: "name",
	},
	Direction: struct {
		ID, UniversityID, CityID, Code, Name, Params, Cost string

		University, City string
	}{
		ID:           "directionID",
		UniversityID: "universityID",
		CityID:       "cityID",
		Code:         "code",
		Name:         "name",
		Params:       "params",
		Cost:         "cost",

		University: "University",
		City:       "City",
	},
	DirectionsFeedback: struct {
		UserID, DirectionID, Rating string

		User, Direction string
	}{
		UserID:      "userID",
		DirectionID: "directionID",
		Rating:      "rating",

		User:      "User",
		Direction: "Direction",
	},
	UniversityFeedback: struct {
		Rating, UserID, UniversityID string

		User, University string
	}{
		Rating:       "rating",
		UserID:       "userID",
		UniversityID: "universityID",

		User:       "User",
		University: "University",
	},
	Universyty: struct {
		ID, Name, CityID string

		City string
	}{
		ID:     "universityID",
		Name:   "Name",
		CityID: "cityID",

		City: "City",
	},
	User: struct {
		ID, CityID, UniversityID, Gender, Age, Params string

		City, University string
	}{
		ID:           "userID",
		CityID:       "cityID",
		UniversityID: "universityID",
		Gender:       "gender",
		Age:          "age",
		Params:       "params",

		City:       "City",
		University: "University",
	},
}

var Tables = struct {
	City struct {
		Name, Alias string
	}
	Direction struct {
		Name, Alias string
	}
	DirectionsFeedback struct {
		Name, Alias string
	}
	UniversityFeedback struct {
		Name, Alias string
	}
	Universyty struct {
		Name, Alias string
	}
	User struct {
		Name, Alias string
	}
}{
	City: struct {
		Name, Alias string
	}{
		Name:  "cities",
		Alias: "t",
	},
	Direction: struct {
		Name, Alias string
	}{
		Name:  "directions",
		Alias: "t",
	},
	DirectionsFeedback: struct {
		Name, Alias string
	}{
		Name:  "directionsFeedbacks",
		Alias: "t",
	},
	UniversityFeedback: struct {
		Name, Alias string
	}{
		Name:  "universityFeedbacks",
		Alias: "t",
	},
	Universyty: struct {
		Name, Alias string
	}{
		Name:  "universyties",
		Alias: "t",
	},
	User: struct {
		Name, Alias string
	}{
		Name:  "users",
		Alias: "t",
	},
}

type City struct {
	tableName struct{} `pg:"cities,alias:t,discard_unknown_columns"`

	ID   int    `pg:"cityID,pk"`
	Name string `pg:"name,use_zero"`
}

type Direction struct {
	tableName struct{} `pg:"directions,alias:t,discard_unknown_columns"`

	ID           int              `pg:"directionID,pk"`
	UniversityID int              `pg:"universityID,use_zero"`
	CityID       int              `pg:"cityID,use_zero"`
	Code         string           `pg:"code,use_zero"`
	Name         string           `pg:"name,use_zero"`
	Params       *DirectionParams `pg:"params"`
	Cost         *int             `pg:"cost"`

	University *Universyty `pg:"fk:universityID,rel:has-one"`
	City       *City       `pg:"fk:cityID,rel:has-one"`
}

type DirectionsFeedback struct {
	tableName struct{} `pg:"directionsFeedbacks,alias:t,discard_unknown_columns"`

	UserID      int     `pg:"userID,pk"`
	DirectionID int     `pg:"directionID,pk"`
	Rating      float32 `pg:"rating,use_zero"`

	User      *User      `pg:"fk:userID,rel:has-one"`
	Direction *Direction `pg:"fk:directionID,rel:has-one"`
}

type UniversityFeedback struct {
	tableName struct{} `pg:"universityFeedbacks,alias:t,discard_unknown_columns"`

	Rating       *int `pg:"rating"`
	UserID       int  `pg:"userID,use_zero"`
	UniversityID int  `pg:"universityID,use_zero"`

	User       *User       `pg:"fk:userID,rel:has-one"`
	University *Universyty `pg:"fk:universityID,rel:has-one"`
}

type Universyty struct {
	tableName struct{} `pg:"universyties,alias:t,discard_unknown_columns"`

	ID     int    `pg:"universityID,pk"`
	Name   string `pg:"Name,use_zero"`
	CityID int    `pg:"cityID,use_zero"`

	City *City `pg:"fk:cityID,rel:has-one"`
}

type User struct {
	tableName struct{} `pg:"users,alias:t,discard_unknown_columns"`

	ID           int         `pg:"userID,pk"`
	CityID       int         `pg:"cityID,use_zero"`
	UniversityID int         `pg:"universityID,use_zero"`
	Gender       *bool       `pg:"gender"`
	Age          *int        `pg:"age"`
	Params       *UserParams `pg:"params"`

	City       *City       `pg:"fk:cityID,rel:has-one"`
	University *Universyty `pg:"fk:universityID,rel:has-one"`
}