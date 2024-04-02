-- delete purchases table before restaurants, menus & users table
DROP TABLE IF EXISTS purchases;

-- delete purchases & menus table before restaurants
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS opening_hours;
DROP TABLE IF EXISTS restaurants;

-- delete the users table
DROP TABLE IF EXISTS users;


-- Delete the index
DROP INDEX IF EXISTS restaurant_name_idx;
DROP INDEX IF EXISTS menu_name_idx;

