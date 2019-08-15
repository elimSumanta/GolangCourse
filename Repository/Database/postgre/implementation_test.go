package database

import (
	mod "github.com/elim/GoCourses/Model"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDBResource_SelectCarByIDGerage(t *testing.T) {
	//setup SQL MOCK
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("DB Mock Error %s", err)
	}

	rowsNull := mock.NewRows([]string{"ownerName", "carName", "idCar"})
	rows := mock.NewRows([]string{"ownerName", "carName", "idCar"}).AddRow("j", "j", "1")

	//mock.ExpectBegin()
	mock.ExpectPrepare(`SELECT (.+) FROM (.+)`)
	mock.ExpectPrepare(`SELECT (.+) FROM (.+)`)

	//setup statement
	stmt1, err := db.Prepare(SelectCarByIDGerage)
	if err != nil {
		t.Fatalf("DB Statement fail %s", err)
	}
	stmt2, err := db.Prepare(SelectCarByIDGerage)
	if err != nil {
		t.Fatalf("DB Statement fail %s", err)
	}

	type fields struct {
		conn                 *sql.DB
		stmtGetCarByIDGerage *sql.Stmt
		stmtGetCarByIDCar    *sql.Stmt
	}
	type args struct {
		idGerage string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		WillReturnRows *sqlmock.Rows
		want           []mod.GerageStatus
	}{
		{
			name: "Case Test -> null Result",
			fields: fields{
				conn:                 db,
				stmtGetCarByIDCar:    stmt2,
				stmtGetCarByIDGerage: stmt1,
			},
			args: args{
				idGerage: "",
			},
			WillReturnRows: rowsNull,
			want:           []mod.GerageStatus{},
		},
		{
			name: "Case 1 -> success",
			fields: fields{
				conn:                 db,
				stmtGetCarByIDCar:    stmt2,
				stmtGetCarByIDGerage: stmt1,
			},
			args: args{
				idGerage: "1",
			},
			WillReturnRows: rows,
			want: []mod.GerageStatus{
				{
					OwnerName: "j",
					CarName:   "j",
					IDCar:     "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			DBRes := &DBResource{
				conn:                 tt.fields.conn,
				stmtGetCarByIDGerage: tt.fields.stmtGetCarByIDGerage,
				stmtGetCarByIDCar:    tt.fields.stmtGetCarByIDCar,
			}

			mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(tt.WillReturnRows)

			if got := DBRes.SelectCarByIDGerage(tt.args.idGerage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBResource.SelectCarByIDGerage() = %v, want %v", got, tt.want)
			}
		})
	}
}
