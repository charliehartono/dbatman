package proxy

/*
import (
	. "github.com/bytedance/dbatman/mysql"
	"testing"
)

func TestConn_Handshake(t *testing.T) {
	c := newTestDBConn(t)

	if err := c.Ping(); err != nil {
		t.Fatal(err)
	}

	c.Close()
}

func TestConn_DeleteTable(t *testing.T) {
	server := newTestServer(t)
	n := server.nodes["node1"]
	c, err := n.getMasterConn()
	if err != nil {
		t.Fatal(err)
	}
	c.UseDB("go_proxy")
	if _, err := c.Execute(`drop table if exists go_proxy_test_proxy_conn`); err != nil {
		t.Fatal(err)
	}
	c.Close()
}


func TestConn_Insert(t *testing.T) {
	s := `insert into go_proxy_test_proxy_conn (id, str, f, e, u, i) values(1, "abc", 3.14, "test1", 255, -127)`

	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 1 {
			t.Fatal(r.AffectedRows)
		}
	}
}

func TestConn_Select(t *testing.T) {
	s := `select str, f, e, u, i, ni from go_proxy_test_proxy_conn where id = 1`

	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 1 {
			t.Fatal(r.RowNumber())
		}

		if r.ColumnNumber() != 6 {
			t.Fatal(r.ColumnNumber())
		}

		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}

		if v, _ := r.GetFloat(0, 1); v != 3.14 {
			t.Fatal(v)
		}

		if v, _ := r.GetString(0, 2); v != `test1` {
			t.Fatal(v)
		}

		if v, _ := r.GetUint(0, 3); v != 255 {
			t.Fatal(v)
		}

		if v, _ := r.GetInt(0, 4); v != -127 {
			t.Fatal(v)
		}

		if v, _ := r.IsNull(0, 5); !v {
			t.Fatal("ni not null")
		}
	}
}

func TestConn_Update(t *testing.T) {
	s := `update go_proxy_test_proxy_conn set str = "123" where id = 1`

	c := newTestDBConn(t)
	defer c.Close()

	if _, err := c.Execute(s); err != nil {
		t.Fatal(err)
	}

	if r, err := c.Execute(`select str from go_proxy_test_proxy_conn where id = 1`); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != `123` {
			t.Fatal(v)
		}
	}
}

func TestConn_Replace(t *testing.T) {
	s := `replace into go_proxy_test_proxy_conn (id, str, f) values(1, "abc", 3.14159)`

	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 2 {
			t.Fatal(r.AffectedRows)
		}
	}

	s = `replace into go_proxy_test_proxy_conn (id, str) values(2, "abcb")`

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 1 {
			t.Fatal(r.AffectedRows)
		}
	}

	s = `select str, f from go_proxy_test_proxy_conn`

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}

		if v, _ := r.GetString(1, 0); v != `abcb` {
			t.Fatal(v)
		}

		if v, _ := r.GetFloat(0, 1); v != 3.14159 {
			t.Fatal(v)
		}

		if v, _ := r.IsNull(1, 1); !v {
			t.Fatal(v)
		}
	}
}

func TestConn_Delete(t *testing.T) {
	s := `delete from go_proxy_test_proxy_conn where id = 100000`

	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 0 {
			t.Fatal(r.AffectedRows)
		}
	}
}

func TestConn_SetAutoCommit(t *testing.T) {
	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute("set autocommit = 1"); err != nil {
		t.Fatal(err)
	} else {
		if !(r.Status&SERVER_STATUS_AUTOCOMMIT > 0) {
			t.Fatal(r.Status)
		}
	}

	if r, err := c.Execute("set autocommit = 0"); err != nil {
		t.Fatal(err)
	} else {
		if !(r.Status&SERVER_STATUS_AUTOCOMMIT == 0) {
			t.Fatal(r.Status)
		}
	}
}

func TestConn_Trans(t *testing.T) {
	c1 := newTestDBConn(t)
	defer c1.Close()

	c2 := newTestDBConn(t)
	defer c2.Close()

	var err error

	if err = c1.Begin(); err != nil {
		t.Fatal(err)
	}

	if err = c2.Begin(); err != nil {
		t.Fatal(err)
	}

	if _, err := c1.Execute(`insert into go_proxy_test_proxy_conn (id, str) values (111, "abc")`); err != nil {
		t.Fatal(err)
	}

	if r, err := c2.Execute(`select str from go_proxy_test_proxy_conn where id = 111`); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 0 {
			t.Fatal(r.RowNumber())
		}
	}

	if err := c1.Commit(); err != nil {
		t.Fatal(err)
	}

	if err := c2.Commit(); err != nil {
		t.Fatal(err)
	}

	if r, err := c1.Execute(`select str from go_proxy_test_proxy_conn where id = 111`); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 1 {
			t.Fatal(r.RowNumber())
		}

		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}
	}
}

func TestConn_SetNames(t *testing.T) {
	c := newTestDBConn(t)
	defer c.Close()

	if err := c.SetCharset("gb2312"); err != nil {
		t.Fatal(err)
	}
}

func TestConn_LastInsertId(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS go_proxy_test_conn_id (
          id BIGINT(64) UNSIGNED AUTO_INCREMENT NOT NULL,
          str VARCHAR(256),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	server := newTestServer(t)
	n := server.nodes["node1"]

	c1, err := n.getMasterConn()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c1.Execute(s); err != nil {
		t.Fatal(err)
	}

	c1.Close()

	c := newTestDBConn(t)
	defer c.Close()

	r, err := c.Execute(`insert into go_proxy_test_conn_id (str) values ("abc")`)
	if err != nil {
		t.Fatal(err)
	}

	lastId := r.InsertId
	if r, err := c.Execute(`select last_insert_id()`); err != nil {
		t.Fatal(err)
	} else {
		if r.ColumnNumber() != 1 {
			t.Fatal(r.ColumnNumber())
		}

		if v, _ := r.GetUint(0, 0); v != lastId {
			t.Fatal(v)
		}
	}

	if r, err := c.Execute(`select last_insert_id() as a`); err != nil {
		t.Fatal(err)
	} else {
		if string(r.Fields[0].Name) != "a" {
			t.Fatal(string(r.Fields[0].Name))
		}

		if v, _ := r.GetUint(0, 0); v != lastId {
			t.Fatal(v)
		}
	}

	c1, _ = n.getMasterConn()

	if _, err := c1.Execute(`drop table if exists go_proxy_test_conn_id`); err != nil {
		t.Fatal(err)
	}

	c1.Close()
}

func TestConn_RowCount(t *testing.T) {
	c := newTestDBConn(t)
	defer c.Close()

	r, err := c.Execute(`insert into go_proxy_test_proxy_conn (id, str) values (1002, "abc")`)
	if err != nil {
		t.Fatal(err)
	}

	row := r.AffectedRows

	if r, err := c.Execute("select row_count()"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetUint(0, 0); v != row {
			t.Fatal(v)
		}
	}

	if r, err := c.Execute("select row_count() as b"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetInt(0, 0); v != -1 {
			t.Fatal(v)
		}
	}
}

func TestConn_SelectVersion(t *testing.T) {
	c := newTestDBConn(t)
	defer c.Close()

	if r, err := c.Execute("select version()"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != ServerVersion {
			t.Fatal(v)
		}
	}
}
*/
