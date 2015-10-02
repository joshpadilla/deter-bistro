CREATE DATABASE bistrofridge;
GRANT ALL ON bistrofridge.* to root@'0.0.0.0';

CREATE TABLE ingredient_packs ( count int );
CREATE TABLE doughballs ( count int );

INSERT INTO ingredient_packs VALUES (150);
INSERT INTO doughballs VALUES (150);
