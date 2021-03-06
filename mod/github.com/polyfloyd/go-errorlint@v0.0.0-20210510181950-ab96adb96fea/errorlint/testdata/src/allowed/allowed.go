package testdata

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

func CompareErrIndirect(r io.Reader) {
	var buf [4096]byte
	_, err := r.Read(buf[:])
	eof := io.EOF

	// Do not bother to check for comparing to aliased std errors, I have never
	// seen any code that does this in the wild. This makes the checker a bit simpler to write.
	// Supporting this use case is acceptable, patches welcome :)
	if err == eof { // want `comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error`
		fmt.Println(err)
	}
}

func CompareAssignIndirect(r io.Reader) {
	var buf [4096]byte
	_, err1 := r.Read(buf[:])
	err2 := err1
	err3 := err2
	if err3 == io.EOF {
		fmt.Println(err3)
	}
}

func CompareInline(db *sql.DB) {
	var i int
	row := db.QueryRow(`SELECT 1`)
	if row.Scan(&i) == sql.ErrNoRows {
		fmt.Println("no rows!")
	}
}

func IoReadEOF(r io.Reader) {
	var buf [4096]byte
	_, err := r.Read(buf[:])
	if err == io.EOF {
		fmt.Println(err)
	}
}

func OsFileReadEOF(fd *os.File) {
	var buf [4096]byte
	_, err := fd.Read(buf[:])
	if err == io.EOF {
		fmt.Println(err)
	}
}

func IoPipeWriterWrite(w *io.PipeWriter) {
	var buf [4096]byte
	_, err := w.Write(buf[:])
	if err == io.ErrClosedPipe {
		fmt.Println(err)
	}
}

func IoReadAtLeast(r io.Reader) {
	var buf [4096]byte
	_, err := io.ReadAtLeast(r, buf[:], 8192)
	if err == io.ErrShortBuffer {
		fmt.Println(err)
	}
	if err == io.ErrUnexpectedEOF {
		fmt.Println(err)
	}
}

func IoReadFull(r io.Reader) {
	var buf [4096]byte
	_, err := io.ReadFull(r, buf[:])
	if err == io.ErrUnexpectedEOF {
		fmt.Println(err)
	}
}

func SqlRowScan(db *sql.DB) {
	var i int
	row := db.QueryRow(`SELECT 1`)
	err := row.Scan(&i)
	if err == sql.ErrNoRows {
		fmt.Println("no rows!")
	}
}
