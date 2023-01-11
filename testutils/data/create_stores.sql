CREATE TABLE IF NOT EXISTS stores
(
    id          SERIAL PRIMARY KEY,
    store_cd    INT NOT NULL,
    company_cd  INT NOT NULL,
    UNIQUE(store_cd,company_cd),
    store_name  VARCHAR(50) NOT NULL,
    address     VARCHAR(255),
    latitude    DECIMAL(7,4),
    longitude   DECIMAL(7,4),
    created_at  TIMESTAMP DEFAULT NOW()
);
