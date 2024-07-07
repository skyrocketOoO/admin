package Session

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	SessionSvc struct {
		Data             map[SessionID]*Session
		AccountToSession map[uint]SessionID
		sync.RWMutex
	}

	SessionID uuid.UUID

	Session struct {
		AccountID uint
		CreatedAt time.Time
	}
)

func NewSessionSvc() (svc *SessionSvc) {
	svc = &SessionSvc{
		Data:             map[SessionID]*Session{},
		AccountToSession: map[uint]SessionID{},
		RWMutex:          sync.RWMutex{},
	}

	svc.AutoClean()
	return
}

func (s *SessionSvc) GetSession(accountID uint) SessionID {
	s.Lock()
	defer s.Unlock()

	if sessionID, ok := s.AccountToSession[accountID]; ok {
		session := s.Data[sessionID]
		session.CreatedAt = time.Now()
		return sessionID
	} else {
		sessionID = SessionID(uuid.New())
		s.Data[sessionID] = &Session{
			AccountID: accountID,
			CreatedAt: time.Now(),
		}
		s.AccountToSession[accountID] = sessionID
		return sessionID
	}
}

// 0 means no session
func (s *SessionSvc) GetAccount(sessionID SessionID) (accountID uint) {
	s.RLock()
	defer s.RUnlock()

	if session, ok := s.Data[sessionID]; ok {
		return session.AccountID
	}
	return 0
}

func (s *SessionSvc) DeleteSession(sessionID SessionID) {
	s.Lock()
	defer s.Unlock()

	if session, ok := s.Data[sessionID]; ok {
		delete(s.AccountToSession, session.AccountID)
		delete(s.Data, sessionID)
	}
}

func (s *SessionSvc) AutoClean() {
	go func() {
		for {
			keyToDel := []SessionID{}
			s.Lock()
			for id, session := range s.Data {
				if session.CreatedAt.Before(time.Now().Add(-10 * time.Minute)) {
					keyToDel = append(keyToDel, id)
					delete(s.AccountToSession, session.AccountID)
				}
			}

			for _, id := range keyToDel {
				delete(s.Data, id)
			}
			s.Unlock()
			time.Sleep(10 * time.Minute)
		}
	}()
}
