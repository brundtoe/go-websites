# Eksempler på Go net/http

## Opdatering december 2024

- Opdateret go til 1.23.4
- øvrige er ikke opdateret

## Opdatering december 2023

- Opdateret go til 1.21
- Opdateret gin til 1.9.1
- Kræver installation af MariaDB på hosten hvor cmd udføres
- Filen $HOME/.config/brundtoe/bookstore.env skal være opdateret med DB_HOST

## http.client

mappen cmd/client indeholder følgende eksempler

Eksempel fra bogen Learning Go Programming
    - gutenberg

valdemar.test med slim4-dbal hvor backend kaldes

**customers**
    - Eksempel med http.client og Client.get

**authors**
    - Eksempel med http.NewRequest 
    - request header Accept application/json
    - request udføres med client.Do

**database** kræver installation af MariaDB på hosten
    - Importerer brundtoe/bookstore/connection
    - go.mod er opdateret med require og replace statements
