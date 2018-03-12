CREATE TYPE status AS ENUM (
  'placed', 'approved', 'delivered', 'cancelled'
);

CREATE TABLE Categories (
  id    UUID PRIMARY KEY,
  name  VARCHAR(64) NOT NULL
);

CREATE TABLE Inventory (
  id    UUID PRIMARY KEY,
  name  VARCHAR(64) NOT NULL,
  cost  DECIMAL(11, 2) NOT NULL,
  url   VARCHAR(512) NULL,
  photoUrls text DEFAULT '{}',
  category UUID NULL REFERENCES Categories
);


CREATE TABLE Orders (
  id UUID PRIMARY KEY,
  item UUID NOT NULL REFERENCES Inventory,
  quantity NUMERIC NOT NULL
);


CREATE TABLE Vend (
  id     UUID PRIMARY KEY ,
  total  DECIMAL(11, 2) NOT NULL,
  status status DEFAULT 'placed'
);

CREATE TABLE VendOrder (
  id UUID PRIMARY KEY,
  vend_ref UUID REFERENCES Vend,
  order_ref UUID REFERENCES Orders
);

CREATE TABLE Users (
  id        UUID PRIMARY KEY,
  username  VARCHAR(64) NULL,
  firstname VARCHAR(64) NULL,
  lastname  VARCHAR(64) NULL,
  email     text UNIQUE NOT NULL,
  password  text UNIQUE NOT NULL,
  phone     integer UNIQUE NULL
);