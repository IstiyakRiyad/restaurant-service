-- User
CREATE TABLE IF NOT EXISTS users (
	id				serial PRIMARY KEY, 
	name			varchar(500) NOT NULL, 
	cash_balance	decimal(10, 2) DEFAULT 0.0
);

-- Restaurant
CREATE TABLE IF NOT EXISTS restaurants (
	id				serial PRIMARY KEY, 
	name			varchar(500) NOT NULL, 
	cash_balance	decimal(10, 2) DEFAULT 0.0
);

-- Menu table
CREATE TABLE IF NOT EXISTS menus (
	id				serial PRIMARY KEY, 
	name			varchar(500) NOT NULL, 
	price			decimal(10, 2) DEFAULT 0.0,
	restaurant_id	INT NOT NULL,

	CONSTRAINT fk_menus_restaurant FOREIGN KEY(restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Opening Hours 
CREATE TABLE IF NOT EXISTS opening_hours (
	id				serial PRIMARY KEY, 
	day				varchar(500) NOT NULL, 
	start_time		timestamp(3) NOT NULL, 
	end_time		timestamp(3) NOT NULL, 
	restaurant_id	INT NOT NULL,

	CONSTRAINT fk_opening_hours_restaurant FOREIGN KEY(restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Purchase 
CREATE TABLE IF NOT EXISTS purchases (
	id				serial PRIMARY KEY, 
	amount			decimal(10, 2) DEFAULT 0.0,
	purchase_time	timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	restaurant_id	INT NOT NULL,
	user_id			INT NOT NULL,
	menu_id			INT NOT NULL,

	CONSTRAINT fk_purchage_restaurant	FOREIGN KEY(restaurant_id) REFERENCES restaurants(id) ON DELETE SET NULL ON UPDATE SET NULL,
	CONSTRAINT fk_purchage_user			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL ON UPDATE SET NULL,
	CONSTRAINT fk_purchage_menus		FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE SET NULL ON UPDATE SET NULL
);

-- Index for restaurant name
CREATE INDEX IF NOT EXISTS restaurant_name_idx ON restaurants(name);

-- Index for restaurant name
CREATE INDEX IF NOT EXISTS menu_name_idx ON menus(name);





