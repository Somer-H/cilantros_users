package adapters

import (
	"fmt"
	"log"
	"users_api/src/core"
	"users_api/src/users/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type MySql struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySql, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySql{conn: conn}, nil
}


func (m *MySql) RegisterUser(user entities.User) (*entities.User, error){
	sqlStatement := `INSERT INTO user(username, password, role, gmail) VALUES (?,?, ?, ?)`
    result, err := m.conn.DB.Exec(sqlStatement, user.Username, user.Password, user.Role, user.Gmail);

    if err != nil {
        return nil, err
    }
    if result!= nil {
       rowsAffected, _ := result.RowsAffected(); 
	   if rowsAffected == 1{
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			user.IdUser = int(lastInsertID)
	   }else {
		log.Printf("[MySQL] - No se ha insertado ningún usuario.")
	   }
	}else{
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
    return &user, nil
}

func (m *MySql) FindUserByUsername(username string) (*entities.User, error) {
    sqlStatement := `SELECT * FROM user WHERE username = ?`
    
    result, err := m.conn.FetchRows(sqlStatement, username)
    if err != nil {
        return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
    }

    // Validar si no hay resultados
    if !result.Next() {
        return nil, fmt.Errorf("usuario no encontrado")
    }
    var user entities.User
    err = result.Scan(&user.IdUser, &user.Username, &user.Password, &user.Role, &user.Gmail )
    if err != nil {
        return nil, fmt.Errorf("error al mapear los resultados: %v", err)
    }

    return &user, nil
}

func (m *MySql) UpdateUser(idUser int, user entities.UserToUpdate) (*entities.User, error) {
	// Consulta para verificar si el usuario existe
	sql := "SELECT * FROM user WHERE idUser = ?"
	resultGet, err1 := m.conn.FetchRows(sql, idUser)
	if err1 != nil {
		return nil, fmt.Errorf("error al obtener el usuario: %v", err1)
	}

	var userGet entities.User
	if resultGet.Next() {
		err := resultGet.Scan(&userGet.IdUser, &userGet.Username, &userGet.Password, &userGet.Role, &userGet.Gmail)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los datos del usuario: %v", err)
		}
	} else {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	if(user.Gmail != ""){
		userGet.Gmail = user.Gmail
	}
	if(user.Password != ""){
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost);
        if err != nil {
            return nil, fmt.Errorf("error al generar la contraseña hash: %v", err)
        }
        userGet.Password = string(hashedPassword)
    }
	sqlStatement := `UPDATE user SET username = ?, password = ?, role = ?, gmail = ? WHERE idUser = ?`
	result, err := m.conn.DB.Exec(sqlStatement, userGet.Username, userGet.Password, userGet.Role, userGet.Gmail, idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de actualización: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error al obtener el número de filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		return &userGet, nil
	} else {
		return nil, fmt.Errorf("usuario no encontrado o no se actualizó")
	}
}