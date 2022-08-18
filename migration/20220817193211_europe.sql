-- +goose Up
-- +goose StatementBegin
CREATE schema europe;
CREATE TABLE IF NOT EXISTS europe.general
(
    country   VARCHAR(20) UNIQUE NOT NULL ,
    capital    VARCHAR(20) NOT NULL,
    square INTEGER   NOT NULL  ,
    population INTEGER   NOT NULL 
 ) ;
CREATE TABLE IF NOT EXISTS europe.economics
(
    country VARCHAR(20) UNIQUE  NOT NULL REFERENCES europe.general(country),
    GDP     INTEGER   NOT NULL 
    
    );
INSERT INTO europe.general VALUES 
(
    'Albania', 'Tirana', 75000, 4000000

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS europe.general CASCADE;
DROP TABLE IF EXISTS europe.economics CASCADE;
Drop schema IF EXISTS europe;
-- +goose StatementEnd
