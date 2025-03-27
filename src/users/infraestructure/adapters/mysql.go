package adapters

import (
	"fmt"
	"log"
	"users_api/src/core"
	"users_api/src/users/domain/entities"
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
	sqlStatement := `INSERT INTO user(username, password) VALUES (?,?)`
    result, err := m.conn.DB.Exec(sqlStatement, user.Username, user.Password);

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
     sqlStatement := `SELECT * FROM user where idUser = ?`; 
	 result, err := m.conn.FetchRows(sqlStatement, username); 
	 if err != nil {
        return nil, err
    }
	var user entities.User
	for result.Next() {
		err := result.Scan(&user.IdUser, &user.Username, &user.Password);
        if err != nil {
            log.Fatal(err)
        }
    }
return &user, nil
}