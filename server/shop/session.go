// Session is special, as the majority of it's functionality will be placed directly in the reverse-forwarders

package shop

import "github.com/satori/go.uuid"

// We use an interface incase we want to use different types of uses
// in the future - see user object.
type SessionObj interface {
	Create(*SessionObj) error
	Read(*SessionObj) error
	Update(*SessionObj, func(obj SessionObj)) error
	Delete(*SessionObj) error

}


// The session interface. We recommend any future object that extends this interface
// to also satisfy the SessionObj as well
type Session interface {
	// Gets a session OR creates a new one
	// RETURNS: the session id
	// NOTE: this session WILL time out
	//TODO: Add time out info to the endpoint description
	//FIXME: Rename - Only creates a new session - middleware handles session check/fetch
	CreateSession()(sessionId string, err error)

	// Wraps up a session by creating a tmp order and returning a uri representing this order
	// RETURNS: a callback URI representing the session
	Logout(session SessionObj)(callback string, err error)

	// Creates a new sessionObj(User)
	CreateUser(session *SessionObj)(userId string, err error)

	// Fetches the session or user by id
	FetchSessionInstanceDetails(sessionObjId uuid.UUID)(foundSession SessionObj, err error)

	// Updates a user by id and User obj
	//QUESTION: Maybe, we should return the user ID again to the user?
	UpdateUser(sessionObjId uuid.UUID, sessionObj *SessionObj)(err error)

	// Deletes the given user at the given ID
	DeleteUser(sessionObjId uuid.UUID)(err error)

}
