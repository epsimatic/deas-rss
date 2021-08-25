package models

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func (a *Article) String() string {
	return a.Guid
}

func (a *Article) ValidateAndUpsertByGUID(tx *pop.Connection) error {
	vErrs, err := a.Validate(tx)
	if err != nil {
		return errors.Wrapf(err, "validate '%s'", a.Guid)
	}
	if vErrs.Count() > 0 {
		return errors.Errorf("vaidation issues for '%s': %s", a.Guid, vErrs)
	}

	return errors.Wrapf(a.UpsertByGUID(tx), "upsert '%s'", a.Guid)
}

func (a *Article) UpsertByGUID(tx *pop.Connection) error {
	fmt.Printf("Looking for id for article '%s'\n", a.Guid)
	var tempArticle Article
	err := tx.Where("guid = ?", a.Guid).Select("id").First(&tempArticle)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrapf(err, "look for article '%s' in DB", a.Guid)
	}
	if err == nil {
		fmt.Printf("Found id=%v for article '%s'\n", tempArticle.ID, a.Guid)
		a.ID = tempArticle.ID
	}

	return errors.Wrap(tx.Save(a), "save")
}
