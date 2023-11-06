package handlers

import (
	"api/models"
	"api/utility"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Register(c echo.Context) error {
	// Get Request Body
	new_user := new(models.User)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := new_user.Is_Populated()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	hashed_pass, err := utility.Hash_Password(new_user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	new_user.Password = hashed_pass

	// Convert User Struct To Map
	user_map, err := new_user.To_Map()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// Connect To Database
	ctx, driver, session, err := utility.Database_Connect_Write()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer driver.Close(ctx)
	defer session.Close(ctx)

	user_id, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"CREATE (u:User) SET u.email = $email, u.name = $name, u.password = $password, u.username = $username RETURN ID(u)", user_map)
		if err != nil {
			return nil, err
		}
		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user_id)
}

func Login(c echo.Context) error {
	// Fetch JSON Body
	user_login := new(models.Login)
	if err := c.Bind(user_login); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Verify JSON
	err := user_login.Is_Populated()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Connect To Database
	ctx, driver, session, err := utility.Database_Connect_Write()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer driver.Close(ctx)
	defer session.Close(ctx)

	login_map, err := user_login.To_Map()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// Validate Credentials
	stored_pass, err := session.ExecuteRead(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (u:User) WHERE u.username = $username RETURN u.password", login_map)
		if err != nil {
			return nil, err
		}
		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = utility.Check_Password(user_login.Password, stored_pass.(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Create New Session
	new_session := new(models.Session)
	new_session.Username = user_login.Username
	new_session.Session_ID = uuid.New().String()

	// Add Session To Database
	session_map, err := new_session.To_Map()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	session_id, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"MATCH (u:User) WHERE u.username = $username CREATE (s:Session {username: $username, session_id: $session_id})-[r:Authenicates]->(u) RETURN ID(r);", session_map)
		if err != nil {
			return nil, err
		}
		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Generate And Set Cookie
	cookie := new(http.Cookie)
	cookie.Name = "session_id"
	cookie.Value = new_session.Session_ID
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, session_id)
}

// func Fetch_User(c echo.Context) error {

// }

// func Update_User(c echo.Context) error {

// }

// func Delete_User(c echo.Context) error {

// }
