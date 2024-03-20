package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// "github.com/jackc/pgx/v5/pgtype"
// "time"

type CreateSessionTxParams struct {
	createSessionParams
}

type SessionTxResult struct {
	Session *Session
}

func (h *Handlers) CreateSession(ctx context.Context, username, role string) (Session, error) {
	var result SessionTxResult

	ranID, err := uuid.NewRandom()

	params := createSessionParams{
		ID:       ranID,
		Username: username,
		Role:     role,
	}

	session, err := h.Queries.createSession(ctx, params)

	if err != nil {
		return Session{}, err
	}
	result.Session = &session
	return session, nil
}

func (h *Handlers) DeleteSession(ctx context.Context, token string) error {
	err := h.Queries.deleteSession(ctx, token)
	if err != nil {
		return fmt.Errorf("error deleting expired sessions: %w", err)
	}
	return nil
}

// SessionCleaner periodically checks for expired sessions and deletes them.
// func (h *Handlers) SessionCleaner(store sessions.Store, interval time.Duration) {
// 	ticker := time.NewTicker(interval)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:

// 			err := h.deleteExpiredSessions(store)
// 			if err != nil {
// 				fmt.Println("Error cleaning expired sessions: ", err)
// 			}
// 		}
// 	}
// }

// // deleteExpiredSessions deletes expired sessions from the session store.
// func (h *Handlers) deleteExpiredSessions(store sessions.Store) error {
// 	sessionIDs, _ := h.queries.getAllSessionID(context.Background())

// 	for _, sessionID := range sessionIDs {
// 		session, _ := store.Get(nil, sessionID)
// 		if isExpired(session) {
// 			session.Options.MaxAge = -1
// 			err := session.Save(nil, nil)
// 			if err != nil {
// 				return fmt.Errorf("error deleting session with ID %s: %w", sessionID, err)
// 			}
// 		}
// 	}
// 	return nil
// }

// // isExpired checks if a session is expired.
// func isExpired(session *sessions.Session) bool {
// 	expireTime, ok := session.Values["expires_at"].(time.Time)
// 	if !ok {
// 		fmt.Println("Warning: expires_at value not found or has invalid type in session")
// 		return false
// 	}
// 	return time.Now().After(expireTime)
// }
