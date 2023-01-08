CREATE TABLE IF NOT EXISTS stores (
  id          SERIAL PRIMARY KEY,
  store_cd    INTEGER NOT NULL,
  company_cd  INTEGER NOT NULL, UNIQUE(store_cd, company_cd),
  store_name  VARCHAR(100) NOT NULL,
  address     VARCHAR(255),
  latitude    DECIMAL(7,4),
  longitude   DECIMAL(7,4),
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
