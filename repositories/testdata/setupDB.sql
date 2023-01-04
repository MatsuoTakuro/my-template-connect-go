CREATE TABLE IF NOT EXISTS stores (
  id          INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  store_cd    INTEGER NOT NULL,
  company_cd  INTEGER NOT NULL, UNIQUE(store_cd, company_cd),
  store_name  VARCHAR(100) NOT NULL,
  address     VARCHAR(255),
  latitude    DECIMAL(7,4),
  longitude   DECIMAL(7,4),
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)engine=innodb charset=utf8mb4;

INSERT INTO stores (store_cd, company_cd, store_name, address, latitude, longitude) VALUES
  (4,1,'スーパーマーケット 田村店','福岡県福岡市早良区田村1-15-5',33.54463620,130.32588870)
  ,(5,1,'スーパーセンター 北九州空港バイパス店','福岡県北九州市小倉南区葛原東4-6-10',33.83731030,130.93260370)
  ,(6,1,'スーパーセンター 宇美店','糟屋郡宇美町ゆりが丘1-5-1',33.55540520,130.53104970)
  ,(7,1,'スーパーセンタータイヨー 本渡店','熊本県天草市丸尾町6-35',32.46904760,130.18486300)
  ,(8,1,'スーパーセンター 上津役店','福岡県北九州市八幡西区中の原2-19-1',33.82974650,130.75156990)
  ,(9,1,'スーパーセンター 豊前店','福岡県豊前市大字四郎丸1102-6',33.61825930,131.09891200)
  ,(10,1,'スーパーセンター 石田店','福岡県北九州市小倉南区八重洲町5-15',33.84003930,130.89101360);


