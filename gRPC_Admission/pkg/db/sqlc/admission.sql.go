// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: admission.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const admission = `-- name: Admission :one
INSERT INTO admission (
  full_name,
  gender,
  dob,
  address,
  phone_no,
  previous_company,
  designation,
  skills
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING full_name, id, gender, dob, address, phone_no, previous_company, designation, skills
`

type AdmissionParams struct {
	FullName        string         `json:"full_name"`
	Gender          string         `json:"gender"`
	Dob             time.Time      `json:"dob"`
	Address         string         `json:"address"`
	PhoneNo         int32          `json:"phone_no"`
	PreviousCompany sql.NullString `json:"previous_company"`
	Designation     sql.NullString `json:"designation"`
	Skills          sql.NullString `json:"skills"`
}

func (q *Queries) Admission(ctx context.Context, arg AdmissionParams) (Admission, error) {
	row := q.db.QueryRowContext(ctx, admission,
		arg.FullName,
		arg.Gender,
		arg.Dob,
		arg.Address,
		arg.PhoneNo,
		arg.PreviousCompany,
		arg.Designation,
		arg.Skills,
	)
	var i Admission
	err := row.Scan(
		&i.FullName,
		&i.ID,
		&i.Gender,
		&i.Dob,
		&i.Address,
		&i.PhoneNo,
		&i.PreviousCompany,
		&i.Designation,
		&i.Skills,
	)
	return i, err
}

const getAdmission = `-- name: GetAdmission :one
SELECT full_name, id, gender, dob, address, phone_no, previous_company, designation, skills FROM admission
WHERE full_name = $1 LIMIT 1
`

func (q *Queries) GetAdmission(ctx context.Context, fullName string) (Admission, error) {
	row := q.db.QueryRowContext(ctx, getAdmission, fullName)
	var i Admission
	err := row.Scan(
		&i.FullName,
		&i.ID,
		&i.Gender,
		&i.Dob,
		&i.Address,
		&i.PhoneNo,
		&i.PreviousCompany,
		&i.Designation,
		&i.Skills,
	)
	return i, err
}

const updateAdmission = `-- name: UpdateAdmission :one
UPDATE admission
SET
  full_name = COALESCE($1, full_name),
  gender = COALESCE($2, gender),
  dob = COALESCE($3, dob),
  address = COALESCE($4, address),
  phone_no = COALESCE($5, phone_no),
  previous_company = COALESCE($6, previous_company),
  designation = COALESCE($7, designation),
  skills = COALESCE($8, skills)
WHERE
  full_name = $1
RETURNING full_name, id, gender, dob, address, phone_no, previous_company, designation, skills
`

type UpdateAdmissionParams struct {
	FullName        sql.NullString `json:"full_name"`
	Gender          sql.NullString `json:"gender"`
	Dob             sql.NullTime   `json:"dob"`
	Address         sql.NullString `json:"address"`
	PhoneNo         sql.NullInt32  `json:"phone_no"`
	PreviousCompany sql.NullString `json:"previous_company"`
	Designation     sql.NullString `json:"designation"`
	Skills          sql.NullString `json:"skills"`
}

func (q *Queries) UpdateAdmission(ctx context.Context, arg UpdateAdmissionParams) (Admission, error) {
	row := q.db.QueryRowContext(ctx, updateAdmission,
		arg.FullName,
		arg.Gender,
		arg.Dob,
		arg.Address,
		arg.PhoneNo,
		arg.PreviousCompany,
		arg.Designation,
		arg.Skills,
	)
	var i Admission
	err := row.Scan(
		&i.FullName,
		&i.ID,
		&i.Gender,
		&i.Dob,
		&i.Address,
		&i.PhoneNo,
		&i.PreviousCompany,
		&i.Designation,
		&i.Skills,
	)
	return i, err
}
