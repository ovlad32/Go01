-- Table: nsi.currency

-- DROP TABLE nsi.currency;

CREATE TABLE nsi.currency
(
  isocode character varying(3),
  name character varying(100),
  full_name character varying(100),
  country character varying(3),
  currency_type character varying(1),
  is_metal boolean
)
WITH (
  OIDS=FALSE
);
ALTER TABLE nsi.currency
  OWNER TO postgres;
GRANT ALL ON TABLE nsi.currency TO postgres;
GRANT SELECT ON TABLE nsi.currency TO gouser;

