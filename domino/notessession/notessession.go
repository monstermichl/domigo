/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSESSION_CLASS.html */
package notessession

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesdatabase"
	"errors"
)

type NotesSession struct {
	domino.NotesStruct
}

func New() (NotesSession, error) {
	var session NotesSession
	com, err := com.CreateObject("Lotus.NotesSession")

	/* Make sure session is released if something happened (https://stackoverflow.com/a/20507179). */
	defer func() {
		if err != nil {
			session.Release()
		}
	}()

	if err != nil {
		return session, errors.New("session could not be initialized")
	}
	session.Com = com
	_, err = session.Com.CallMethod("Initialize")

	if err != nil {
		return session, errors.New("session could not be initialized")
	}
	return session, err
}

func (s NotesSession) Release() {
	s.Com.Release()
}

func (s NotesSession) GetDatabase(server, dbPath string) (notesdatabase.NotesDatabase, error) {
	var db notesdatabase.NotesDatabase
	dispatchPtr, err := s.Com.CallObjectMethod("GetDatabase", server, dbPath, false)

	if err != nil {
		return db, errors.New("database could not be retrieved")
	}
	return notesdatabase.New(dispatchPtr), err
}
