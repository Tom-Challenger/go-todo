/*
  Создадим таблицу, где:
  id - генерируется порядковые номера, обязательно заполнен и значения уникальные
  name - массив символов до 255 в длину, обязательтно заполяется
  username - массив символов до 255 в длину, обязательтно заполяется и значения уникальные в рамках таблицы
  password_hash - массив символов до 255 в длину, обязательтно заполяется
*/
CREATE TABLE users
(
    id  serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

/*
  Создадим таблицу, где:
  id - первичный ключ записи
  user_id - первичный ключ записи о пользователе
  list_id - первичный ключ записи о списке который принадлежит пользователю

  При этом, значением будет ссылка на целое число (int)
  и при удалении записи должна быть выполнена очистка от дочерних записей в таблице lists и items
*/
CREATE TABLE users_lists
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

/*
  Создадим таблицу для зранения записей о задачах, где:
  done - содержит boolean, не пустое и при создании по умолчанию будет false
*/
CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);