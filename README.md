# eksempler på Go net/http

## Opdatering december 2023

- opdateret go til 1.21
- opdateret gin til 1.9.1
- kræver installation af MariaDB på hosten hvor cmd udføres


## http.client

mappen cmd/client indeholder følgende eksempler

Eksempel fra bogen Learning Go Programming
    - gutenberg

leonora.test med slim4-dbal hvor backend kaldes

**customers**
    - Eksempel med http.client og Client.get

**authors**
    - Eksempel med http.NewRequest 
    - request header Accept application/json
    - request udføres med client.Do

**database** kræver installation af MariaDB på hosten
    - Importerer brundtoe/bookstore/connection
    - go.mod er opdateret med require og replace statements
