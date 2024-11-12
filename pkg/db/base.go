package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type BaseRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewBaseRepo returns new repository
func NewBaseRepo(db orm.DB) BaseRepo {
	return BaseRepo{
		db:      db,
		filters: map[string][]Filter{},
		sort: map[string][]SortField{
			Tables.City.Name:               {{Column: Columns.City.ID, Direction: SortDesc}},
			Tables.Direction.Name:          {{Column: Columns.Direction.ID, Direction: SortDesc}},
			Tables.DirectionsFeedback.Name: {{Column: Columns.DirectionsFeedback.UserID, Direction: SortDesc}},
			Tables.Universyty.Name:         {{Column: Columns.Universyty.ID, Direction: SortDesc}},
			Tables.User.Name:               {{Column: Columns.User.ID, Direction: SortDesc}},
		},
		join: map[string][]string{
			Tables.City.Name:               {TableColumns},
			Tables.Direction.Name:          {TableColumns, Columns.Direction.University, Columns.Direction.City},
			Tables.DirectionsFeedback.Name: {TableColumns, Columns.DirectionsFeedback.User, Columns.DirectionsFeedback.Direction},
			Tables.UniversityFeedback.Name: {TableColumns, Columns.UniversityFeedback.User, Columns.UniversityFeedback.University},
			Tables.Universyty.Name:         {TableColumns, Columns.Universyty.City},
			Tables.User.Name:               {TableColumns, Columns.User.City, Columns.User.University},
		},
	}
}

// WithTransaction is a function that wraps BaseRepo with pg.Tx transaction.
func (br BaseRepo) WithTransaction(tx *pg.Tx) BaseRepo {
	br.db = tx
	return br
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (br BaseRepo) WithEnabledOnly() BaseRepo {
	f := make(map[string][]Filter, len(br.filters))
	for i := range br.filters {
		f[i] = make([]Filter, len(br.filters[i]))
		copy(f[i], br.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	br.filters = f

	return br
}

/*** City ***/

// FullCity returns full joins with all columns
func (br BaseRepo) FullCity() OpFunc {
	return WithColumns(br.join[Tables.City.Name]...)
}

// DefaultCitySort returns default sort.
func (br BaseRepo) DefaultCitySort() OpFunc {
	return WithSort(br.sort[Tables.City.Name]...)
}

// CityByID is a function that returns City by ID(s) or nil.
func (br BaseRepo) CityByID(ctx context.Context, id int, ops ...OpFunc) (*City, error) {
	return br.OneCity(ctx, &CitySearch{ID: &id}, ops...)
}

// OneCity is a function that returns one City by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneCity(ctx context.Context, search *CitySearch, ops ...OpFunc) (*City, error) {
	obj := &City{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.City.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// CitiesByFilters returns City list.
func (br BaseRepo) CitiesByFilters(ctx context.Context, search *CitySearch, pager Pager, ops ...OpFunc) (cities []City, err error) {
	err = buildQuery(ctx, br.db, &cities, search, br.filters[Tables.City.Name], pager, ops...).Select()
	return
}

// CountCities returns count
func (br BaseRepo) CountCities(ctx context.Context, search *CitySearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &City{}, search, br.filters[Tables.City.Name], PagerOne, ops...).Count()
}

// AddCity adds City to DB.
func (br BaseRepo) AddCity(ctx context.Context, city *City, ops ...OpFunc) (*City, error) {
	q := br.db.ModelContext(ctx, city)
	applyOps(q, ops...)
	_, err := q.Insert()

	return city, err
}

// UpdateCity updates City in DB.
func (br BaseRepo) UpdateCity(ctx context.Context, city *City, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, city).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.City.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteCity deletes City from DB.
func (br BaseRepo) DeleteCity(ctx context.Context, id int) (deleted bool, err error) {
	city := &City{ID: id}

	res, err := br.db.ModelContext(ctx, city).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** Direction ***/

// FullDirection returns full joins with all columns
func (br BaseRepo) FullDirection() OpFunc {
	return WithColumns(br.join[Tables.Direction.Name]...)
}

// DefaultDirectionSort returns default sort.
func (br BaseRepo) DefaultDirectionSort() OpFunc {
	return WithSort(br.sort[Tables.Direction.Name]...)
}

// DirectionByID is a function that returns Direction by ID(s) or nil.
func (br BaseRepo) DirectionByID(ctx context.Context, id int, ops ...OpFunc) (*Direction, error) {
	return br.OneDirection(ctx, &DirectionSearch{ID: &id}, ops...)
}

// OneDirection is a function that returns one Direction by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneDirection(ctx context.Context, search *DirectionSearch, ops ...OpFunc) (*Direction, error) {
	obj := &Direction{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.Direction.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// DirectionsByFilters returns Direction list.
func (br BaseRepo) DirectionsByFilters(ctx context.Context, search *DirectionSearch, pager Pager, ops ...OpFunc) (directions []Direction, err error) {
	err = buildQuery(ctx, br.db, &directions, search, br.filters[Tables.Direction.Name], pager, ops...).Select()
	return
}

// CountDirections returns count
func (br BaseRepo) CountDirections(ctx context.Context, search *DirectionSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &Direction{}, search, br.filters[Tables.Direction.Name], PagerOne, ops...).Count()
}

// AddDirection adds Direction to DB.
func (br BaseRepo) AddDirection(ctx context.Context, direction *Direction, ops ...OpFunc) (*Direction, error) {
	q := br.db.ModelContext(ctx, direction)
	applyOps(q, ops...)
	_, err := q.Insert()

	return direction, err
}

// UpdateDirection updates Direction in DB.
func (br BaseRepo) UpdateDirection(ctx context.Context, direction *Direction, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, direction).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.Direction.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteDirection deletes Direction from DB.
func (br BaseRepo) DeleteDirection(ctx context.Context, id int) (deleted bool, err error) {
	direction := &Direction{ID: id}

	res, err := br.db.ModelContext(ctx, direction).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** DirectionsFeedback ***/

// FullDirectionsFeedback returns full joins with all columns
func (br BaseRepo) FullDirectionsFeedback() OpFunc {
	return WithColumns(br.join[Tables.DirectionsFeedback.Name]...)
}

// DefaultDirectionsFeedbackSort returns default sort.
func (br BaseRepo) DefaultDirectionsFeedbackSort() OpFunc {
	return WithSort(br.sort[Tables.DirectionsFeedback.Name]...)
}

// DirectionsFeedbackByID is a function that returns DirectionsFeedback by ID(s) or nil.
func (br BaseRepo) DirectionsFeedbackByID(ctx context.Context, userID int, directionID int, ops ...OpFunc) (*DirectionsFeedback, error) {
	return br.OneDirectionsFeedback(ctx, &DirectionsFeedbackSearch{UserID: &userID, DirectionID: &directionID}, ops...)
}

// OneDirectionsFeedback is a function that returns one DirectionsFeedback by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneDirectionsFeedback(ctx context.Context, search *DirectionsFeedbackSearch, ops ...OpFunc) (*DirectionsFeedback, error) {
	obj := &DirectionsFeedback{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.DirectionsFeedback.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// DirectionsFeedbacksByFilters returns DirectionsFeedback list.
func (br BaseRepo) DirectionsFeedbacksByFilters(ctx context.Context, search *DirectionsFeedbackSearch, pager Pager, ops ...OpFunc) (directionsFeedbacks []DirectionsFeedback, err error) {
	err = buildQuery(ctx, br.db, &directionsFeedbacks, search, br.filters[Tables.DirectionsFeedback.Name], pager, ops...).Select()
	return
}

// CountDirectionsFeedbacks returns count
func (br BaseRepo) CountDirectionsFeedbacks(ctx context.Context, search *DirectionsFeedbackSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &DirectionsFeedback{}, search, br.filters[Tables.DirectionsFeedback.Name], PagerOne, ops...).Count()
}

// AddDirectionsFeedback adds DirectionsFeedback to DB.
func (br BaseRepo) AddDirectionsFeedback(ctx context.Context, directionsFeedback *DirectionsFeedback, ops ...OpFunc) (*DirectionsFeedback, error) {
	q := br.db.ModelContext(ctx, directionsFeedback)
	applyOps(q, ops...)
	_, err := q.Insert()

	return directionsFeedback, err
}

// UpdateDirectionsFeedback updates DirectionsFeedback in DB.
func (br BaseRepo) UpdateDirectionsFeedback(ctx context.Context, directionsFeedback *DirectionsFeedback, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, directionsFeedback).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.DirectionsFeedback.UserID, Columns.DirectionsFeedback.DirectionID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteDirectionsFeedback deletes DirectionsFeedback from DB.
func (br BaseRepo) DeleteDirectionsFeedback(ctx context.Context, userID int, directionID int) (deleted bool, err error) {
	directionsFeedback := &DirectionsFeedback{UserID: userID, DirectionID: directionID}

	res, err := br.db.ModelContext(ctx, directionsFeedback).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** UniversityFeedback ***/

// FullUniversityFeedback returns full joins with all columns
func (br BaseRepo) FullUniversityFeedback() OpFunc {
	return WithColumns(br.join[Tables.UniversityFeedback.Name]...)
}

// DefaultUniversityFeedbackSort returns default sort.
func (br BaseRepo) DefaultUniversityFeedbackSort() OpFunc {
	return WithSort(br.sort[Tables.UniversityFeedback.Name]...)
}

// OneUniversityFeedback is a function that returns one UniversityFeedback by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneUniversityFeedback(ctx context.Context, search *UniversityFeedbackSearch, ops ...OpFunc) (*UniversityFeedback, error) {
	obj := &UniversityFeedback{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.UniversityFeedback.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// UniversityFeedbacksByFilters returns UniversityFeedback list.
func (br BaseRepo) UniversityFeedbacksByFilters(ctx context.Context, search *UniversityFeedbackSearch, pager Pager, ops ...OpFunc) (universityFeedbacks []UniversityFeedback, err error) {
	err = buildQuery(ctx, br.db, &universityFeedbacks, search, br.filters[Tables.UniversityFeedback.Name], pager, ops...).Select()
	return
}

// CountUniversityFeedbacks returns count
func (br BaseRepo) CountUniversityFeedbacks(ctx context.Context, search *UniversityFeedbackSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &UniversityFeedback{}, search, br.filters[Tables.UniversityFeedback.Name], PagerOne, ops...).Count()
}

// AddUniversityFeedback adds UniversityFeedback to DB.
func (br BaseRepo) AddUniversityFeedback(ctx context.Context, universityFeedback *UniversityFeedback, ops ...OpFunc) (*UniversityFeedback, error) {
	q := br.db.ModelContext(ctx, universityFeedback)
	applyOps(q, ops...)
	_, err := q.Insert()

	return universityFeedback, err
}

// UpdateUniversityFeedback updates UniversityFeedback in DB.
func (br BaseRepo) UpdateUniversityFeedback(ctx context.Context, universityFeedback *UniversityFeedback, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, universityFeedback).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** Universyty ***/

// FullUniversyty returns full joins with all columns
func (br BaseRepo) FullUniversyty() OpFunc {
	return WithColumns(br.join[Tables.Universyty.Name]...)
}

// DefaultUniversytySort returns default sort.
func (br BaseRepo) DefaultUniversytySort() OpFunc {
	return WithSort(br.sort[Tables.Universyty.Name]...)
}

// UniversytyByID is a function that returns Universyty by ID(s) or nil.
func (br BaseRepo) UniversytyByID(ctx context.Context, id int, ops ...OpFunc) (*Universyty, error) {
	return br.OneUniversyty(ctx, &UniversytySearch{ID: &id}, ops...)
}

// OneUniversyty is a function that returns one Universyty by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneUniversyty(ctx context.Context, search *UniversytySearch, ops ...OpFunc) (*Universyty, error) {
	obj := &Universyty{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.Universyty.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// UniversytiesByFilters returns Universyty list.
func (br BaseRepo) UniversytiesByFilters(ctx context.Context, search *UniversytySearch, pager Pager, ops ...OpFunc) (universyties []Universyty, err error) {
	err = buildQuery(ctx, br.db, &universyties, search, br.filters[Tables.Universyty.Name], pager, ops...).Select()
	return
}

// CountUniversyties returns count
func (br BaseRepo) CountUniversyties(ctx context.Context, search *UniversytySearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &Universyty{}, search, br.filters[Tables.Universyty.Name], PagerOne, ops...).Count()
}

// AddUniversyty adds Universyty to DB.
func (br BaseRepo) AddUniversyty(ctx context.Context, universyty *Universyty, ops ...OpFunc) (*Universyty, error) {
	q := br.db.ModelContext(ctx, universyty)
	applyOps(q, ops...)
	_, err := q.Insert()

	return universyty, err
}

// UpdateUniversyty updates Universyty in DB.
func (br BaseRepo) UpdateUniversyty(ctx context.Context, universyty *Universyty, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, universyty).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.Universyty.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteUniversyty deletes Universyty from DB.
func (br BaseRepo) DeleteUniversyty(ctx context.Context, id int) (deleted bool, err error) {
	universyty := &Universyty{ID: id}

	res, err := br.db.ModelContext(ctx, universyty).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** User ***/

// FullUser returns full joins with all columns
func (br BaseRepo) FullUser() OpFunc {
	return WithColumns(br.join[Tables.User.Name]...)
}

// DefaultUserSort returns default sort.
func (br BaseRepo) DefaultUserSort() OpFunc {
	return WithSort(br.sort[Tables.User.Name]...)
}

// UserByID is a function that returns User by ID(s) or nil.
func (br BaseRepo) UserByID(ctx context.Context, id int, ops ...OpFunc) (*User, error) {
	return br.OneUser(ctx, &UserSearch{ID: &id}, ops...)
}

// OneUser is a function that returns one User by filters. It could return pg.ErrMultiRows.
func (br BaseRepo) OneUser(ctx context.Context, search *UserSearch, ops ...OpFunc) (*User, error) {
	obj := &User{}
	err := buildQuery(ctx, br.db, obj, search, br.filters[Tables.User.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// UsersByFilters returns User list.
func (br BaseRepo) UsersByFilters(ctx context.Context, search *UserSearch, pager Pager, ops ...OpFunc) (users []User, err error) {
	err = buildQuery(ctx, br.db, &users, search, br.filters[Tables.User.Name], pager, ops...).Select()
	return
}

// CountUsers returns count
func (br BaseRepo) CountUsers(ctx context.Context, search *UserSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, br.db, &User{}, search, br.filters[Tables.User.Name], PagerOne, ops...).Count()
}

// AddUser adds User to DB.
func (br BaseRepo) AddUser(ctx context.Context, user *User, ops ...OpFunc) (*User, error) {
	q := br.db.ModelContext(ctx, user)
	applyOps(q, ops...)
	_, err := q.Insert()

	return user, err
}

// UpdateUser updates User in DB.
func (br BaseRepo) UpdateUser(ctx context.Context, user *User, ops ...OpFunc) (bool, error) {
	q := br.db.ModelContext(ctx, user).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.User.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteUser deletes User from DB.
func (br BaseRepo) DeleteUser(ctx context.Context, id int) (deleted bool, err error) {
	user := &User{ID: id}

	res, err := br.db.ModelContext(ctx, user).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}
