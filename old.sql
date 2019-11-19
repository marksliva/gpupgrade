--
-- Greenplum Database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE gpadmin;
ALTER ROLE gpadmin WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN NOREPLICATION PASSWORD 'md5e438385a827b74c540ae46b8f1781fab' CREATEEXTTABLE (protocol='gpfdist', type='readable') CREATEEXTTABLE (protocol='gpfdist', type='writable') CREATEEXTTABLE (protocol='http');








--
-- Database creation
--

SET allow_system_table_mods = true;
RESET allow_system_table_mods;
SET allow_system_table_mods = true;
RESET allow_system_table_mods;
REVOKE ALL ON DATABASE template1 FROM PUBLIC;
REVOKE ALL ON DATABASE template1 FROM gpadmin;
GRANT ALL ON DATABASE template1 TO gpadmin;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;
SET allow_system_table_mods = true;
CREATE DATABASE test WITH TEMPLATE = template0 OWNER = gpadmin;
RESET allow_system_table_mods;


\connect postgres

SET default_transaction_read_only = off;

--
-- Greenplum Database database dump
--

SET gp_default_storage_options = '';
SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;

SET default_with_oids = false;

--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: gpadmin
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM gpadmin;
GRANT ALL ON SCHEMA public TO gpadmin;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Greenplum Database database dump complete
--

\connect template1

SET default_transaction_read_only = off;

--
-- Greenplum Database database dump
--

SET gp_default_storage_options = '';
SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;

SET default_with_oids = false;

--
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: gpadmin
--

COMMENT ON DATABASE template1 IS 'default template database';


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: gpadmin
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM gpadmin;
GRANT ALL ON SCHEMA public TO gpadmin;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Greenplum Database database dump complete
--

\connect test

SET default_transaction_read_only = off;

--
-- Greenplum Database database dump
--

SET gp_default_storage_options = '';
SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: test_table; Type: TABLE; Schema: public; Owner: gpadmin; Tablespace: 
--

CREATE TABLE public.test_table (
    a integer,
    b character varying
) DISTRIBUTED BY (a);


ALTER TABLE public.test_table OWNER TO gpadmin;

--
-- Data for Name: test_table; Type: TABLE DATA; Schema: public; Owner: gpadmin
--

COPY public.test_table (a, b) FROM stdin;
1	a
4	d
2	b
3	c
\.


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: gpadmin
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM gpadmin;
GRANT ALL ON SCHEMA public TO gpadmin;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Greenplum Database database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

