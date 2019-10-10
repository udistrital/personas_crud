-- object: persona | type: SCHEMA --
-- DROP SCHEMA IF EXISTS persona CASCADE;
CREATE SCHEMA persona;
-- ddl-end --

SET search_path TO pg_catalog,public,persona;
-- ddl-end --

-- object: persona.estado_civil | type: TABLE --
-- DROP TABLE IF EXISTS persona.estado_civil CASCADE;
CREATE TABLE persona.estado_civil (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_estado_civil PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.estado_civil IS 'Almacena los parametros de estado civil.';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.id IS 'identificador de la tabla estado_civil';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.estado_civil.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_civil ON persona.estado_civil  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: persona.genero | type: TABLE --
-- DROP TABLE IF EXISTS persona.genero CASCADE;
CREATE TABLE persona.genero (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_genero PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.genero IS 'Almacena los tipos de genero ';
-- ddl-end --
COMMENT ON COLUMN persona.genero.id IS 'identificador de la tabla genero';
-- ddl-end --
COMMENT ON COLUMN persona.genero.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro de tipo genero';
-- ddl-end --
COMMENT ON COLUMN persona.genero.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN persona.genero.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN persona.genero.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.genero.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_genero ON persona.genero  IS 'llave primaria de la tabla tipo genero';
-- ddl-end --

-- object: persona.grupo_etnico | type: TABLE --
-- DROP TABLE IF EXISTS persona.grupo_etnico CASCADE;
CREATE TABLE persona.grupo_etnico (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_grupo_etnico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.grupo_etnico IS 'Almacena los grupos etnicos en los que se catalogan las personas.';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.id IS 'identificador de la tabla grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_etnico.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_grupo_etnico ON persona.grupo_etnico  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: persona.grupo_sanguineo_persona | type: TABLE --
-- DROP TABLE IF EXISTS persona.grupo_sanguineo_persona CASCADE;
CREATE TABLE persona.grupo_sanguineo_persona (
	id serial NOT NULL,
	factor_rh character varying(1) NOT NULL,
	grupo_sanguineo character varying(2) NOT NULL,
	persona bigint NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	activo boolean NOT NULL,
	CONSTRAINT pk_grupo_sanguineo_persona PRIMARY KEY (id),
	CONSTRAINT uq_grupo_sanguineo_persona UNIQUE (persona)

);
-- ddl-end --
COMMENT ON TABLE persona.grupo_sanguineo_persona IS 'Tabla que almacena el RH de una persona';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_sanguineo_persona.id IS 'Identificador de la tabla grupo_sanguineo_persona';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_sanguineo_persona.factor_rh IS 'Factor correspondiente al RH de la persona (+/-)';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_sanguineo_persona.grupo_sanguineo IS 'Campo que almacena el nombre del grupo sanguíneo de la persona (A, B, AB, O)';
-- ddl-end --
COMMENT ON COLUMN persona.grupo_sanguineo_persona.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_grupo_sanguineo_persona ON persona.grupo_sanguineo_persona  IS 'Restricción de llave primaria de la tabla grupo_sanguineo_persona';
-- ddl-end --
COMMENT ON CONSTRAINT uq_grupo_sanguineo_persona ON persona.grupo_sanguineo_persona  IS 'Restricción para que una persona solo tenga un grupo sanguíneo';
-- ddl-end --

-- object: persona.perfil_profesional | type: TABLE --
-- DROP TABLE IF EXISTS persona.perfil_profesional CASCADE;
CREATE TABLE persona.perfil_profesional (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_perfil_profesional PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.perfil_profesional IS 'tabla en la que se registran los diferntes perfiles profesionales para las personas.';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.id IS 'identificador de la tabla perfil_profesional';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.nombre IS 'Campo que indica el nombre del perfil.';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.perfil_profesional.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_perfil_profesional ON persona.perfil_profesional  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: persona.persona | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona CASCADE;
CREATE TABLE persona.persona (
	id serial NOT NULL,
	primer_nombre character varying(50) NOT NULL,
	segundo_nombre character varying(50),
	primer_apellido character varying(50) NOT NULL,
	segundo_apellido character varying(50),
	fecha_nacimiento date,
	usuario character varying(50),
	ente integer,
	foto integer,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona PRIMARY KEY (id),
	CONSTRAINT uq_ente UNIQUE (ente)

);
-- ddl-end --
COMMENT ON TABLE persona.persona IS 'tabla que almacena la informacion basica geral relacionada a las personas';
-- ddl-end --
COMMENT ON COLUMN persona.persona.id IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona.primer_nombre IS 'Corresponde al primer nombre de la persona.';
-- ddl-end --
COMMENT ON COLUMN persona.persona.segundo_nombre IS 'Corresponde al segundo nombre de la persona.';
-- ddl-end --
COMMENT ON COLUMN persona.persona.primer_apellido IS 'Corresponde al primer apellidode la persona.';
-- ddl-end --
COMMENT ON COLUMN persona.persona.segundo_apellido IS 'Corresponde al segundo apellido de la persona.';
-- ddl-end --
COMMENT ON COLUMN persona.persona.fecha_nacimiento IS 'Fecha de nacimiento de la persona.';
-- ddl-end --
COMMENT ON COLUMN persona.persona.usuario IS 'identificador del usuario relacionado al sistema de autenticación';
-- ddl-end --
COMMENT ON COLUMN persona.persona.ente IS 'Identificador de la tabla ente';
-- ddl-end --
COMMENT ON COLUMN persona.persona.foto IS 'Corresponde al id de la tabla documento';
-- ddl-end --
COMMENT ON COLUMN persona.persona.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona ON persona.persona  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: persona.persona_estado_civil | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona_estado_civil CASCADE;
CREATE TABLE persona.persona_estado_civil (
	id serial NOT NULL,
	estado_civil integer NOT NULL,
	persona integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona_estado_civil PRIMARY KEY (id),
	CONSTRAINT uq_estado_civil_persona UNIQUE (estado_civil,persona)

);
-- ddl-end --
COMMENT ON TABLE persona.persona_estado_civil IS 'Tabla que relaciona la persona con el estado civil.';
-- ddl-end --
COMMENT ON COLUMN persona.persona_estado_civil.id IS 'Identificador de la tabla persona_estado_civil';
-- ddl-end --
COMMENT ON COLUMN persona.persona_estado_civil.estado_civil IS 'Identificador foraneo a la tabla parametrica estado_civil.';
-- ddl-end --
COMMENT ON COLUMN persona.persona_estado_civil.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona_estado_civil.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_estado_civil ON persona.persona_estado_civil IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: persona.persona_genero | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona_genero CASCADE;
CREATE TABLE persona.persona_genero (
	id serial NOT NULL,
	genero integer NOT NULL,
	persona integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona_genero PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.persona_genero IS 'Tabla que relaciona la persona con el correspondiente genero';
-- ddl-end --
COMMENT ON COLUMN persona.persona_genero.id IS 'Identificador de la tabla persona_genero';
-- ddl-end --
COMMENT ON COLUMN persona.persona_genero.genero IS 'Identificador de la tabla genero.';
-- ddl-end --
COMMENT ON COLUMN persona.persona_genero.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona_genero.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_genero ON persona.persona_genero  IS 'Restriccion de llave primaria de la tabla persona genero.';
-- ddl-end --

-- object: persona.persona_grupo_etnico | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona_grupo_etnico CASCADE;
CREATE TABLE persona.persona_grupo_etnico (
	id serial NOT NULL,
	grupo_etnico integer NOT NULL,
	persona integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona_grupo_etnico PRIMARY KEY (id),
	CONSTRAINT uq_grupo_etnico_persona UNIQUE (persona)

);
-- ddl-end --
COMMENT ON TABLE persona.persona_grupo_etnico IS 'Tabla que relaciona la persona con el grupo étnico.';
-- ddl-end --
COMMENT ON COLUMN persona.persona_grupo_etnico.id IS 'Identificador de la tabla persona_grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN persona.persona_grupo_etnico.grupo_etnico IS 'Identificador de la tbala grupo_etnico';
-- ddl-end --
COMMENT ON COLUMN persona.persona_grupo_etnico.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona_grupo_etnico.activo IS 'Estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_grupo_etnico ON persona.persona_grupo_etnico  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --
COMMENT ON CONSTRAINT uq_grupo_etnico_persona ON persona.persona_grupo_etnico  IS 'Restricción para que una persona tenga máximo un grupo étnico';
-- ddl-end --

-- object: persona.persona_perfil_profesional | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona_perfil_profesional CASCADE;
CREATE TABLE persona.persona_perfil_profesional (
	id serial NOT NULL,
	perfil_profesional integer NOT NULL,
	persona integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona_perfil_profesional PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.persona_perfil_profesional IS 'Tabla que relaciona la persona con el perfil.';
-- ddl-end --
COMMENT ON COLUMN persona.persona_perfil_profesional.id IS 'Identificador primario de la tabla';
-- ddl-end --
COMMENT ON COLUMN persona.persona_perfil_profesional.perfil_profesional IS 'Identificador de la tabla perfil_profesional';
-- ddl-end --
COMMENT ON COLUMN persona.persona_perfil_profesional.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona_perfil_profesional.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_perfil_profesional ON persona.persona_perfil_profesional  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: persona.persona_tipo_discapacidad | type: TABLE --
-- DROP TABLE IF EXISTS persona.persona_tipo_discapacidad CASCADE;
CREATE TABLE persona.persona_tipo_discapacidad (
	id serial NOT NULL,
	tipo_discapacidad integer NOT NULL,
	persona integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_persona_tipo_discapacidad PRIMARY KEY (id),
	CONSTRAINT uq_discapacidad_persona UNIQUE (tipo_discapacidad,persona)

);
-- ddl-end --
COMMENT ON TABLE persona.persona_tipo_discapacidad IS 'Tabla que relaciona la persona con el tipo de discapacidad';
-- ddl-end --
COMMENT ON COLUMN persona.persona_tipo_discapacidad.id IS 'Identificador  de la tabla persona_tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN persona.persona_tipo_discapacidad.tipo_discapacidad IS 'Identificador de la tabla tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN persona.persona_tipo_discapacidad.persona IS 'Identificador de la tabla persona';
-- ddl-end --
COMMENT ON COLUMN persona.persona_tipo_discapacidad.activo IS 'Campo para inactivar una discapacidad cuando la persona lo requiera';
-- ddl-end --
COMMENT ON CONSTRAINT pk_persona_tipo_discapacidad ON persona.persona_tipo_discapacidad  IS 'Restriccion de llave primaria de la tabla.';
-- ddl-end --

-- object: persona.relacion_personas | type: TABLE --
-- DROP TABLE IF EXISTS persona.relacion_personas CASCADE;
CREATE TABLE persona.relacion_personas (
	id serial NOT NULL,
	persona_principal integer NOT NULL,
	persona_relacionada integer NOT NULL,
	tipo_relacion_personas integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_relacion_personas PRIMARY KEY (id),
	CONSTRAINT uq_relacion_personas UNIQUE (persona_principal,persona_relacionada,tipo_relacion_personas)

);
-- ddl-end --
COMMENT ON TABLE persona.relacion_personas IS 'Describe la relación de parentesco que tiene una persona principal con una persona relacionada';
-- ddl-end --
COMMENT ON COLUMN persona.relacion_personas.id IS 'identificador de la tabla relacion_personas';
-- ddl-end --
COMMENT ON COLUMN persona.relacion_personas.persona_principal IS 'identificador de la tabla persona, que indica la persona a la que se le indica la relacion';
-- ddl-end --
COMMENT ON COLUMN persona.relacion_personas.persona_relacionada IS 'identificador de la tabla persona, que indica la persona que es relacionada';
-- ddl-end --
COMMENT ON COLUMN persona.relacion_personas.activo IS 'Indica el estado del registro';
-- ddl-end --

-- object: persona.tipo_discapacidad | type: TABLE --
-- DROP TABLE IF EXISTS persona.tipo_discapacidad CASCADE;
CREATE TABLE persona.tipo_discapacidad (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_tipo_discapacidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.tipo_discapacidad IS 'Almacena los parametros de tipo discapacidad.';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.id IS 'identificador unico de la tabla tipo_discapacidad';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_discapacidad.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_discapacidad ON persona.tipo_discapacidad  IS 'llave primaria de la tabla';
-- ddl-end --

-- object: persona.tipo_relacion_personas | type: TABLE --
-- DROP TABLE IF EXISTS persona.tipo_relacion_personas CASCADE;
CREATE TABLE persona.tipo_relacion_personas (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_tipo_relacion_personas PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE persona.tipo_relacion_personas IS 'Tabla parametrica que indica el tipo de relacion que pueden tener las personas. Ejemplo: Madre, Padre, Conyugue';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.id IS 'identificador de la tabla tipo_relacion_personas';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.nombre IS 'Campo que indica el nombre del tipo de relación';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.descripcion IS 'Campo para agregar una descripcion acerca del tipo de relacion';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.codigo_abreviacion IS 'Codigo de abreviacion opcional para referirse al registro';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.activo IS 'Campo que indica si el tipo de relación esta activo';
-- ddl-end --
COMMENT ON COLUMN persona.tipo_relacion_personas.numero_orden IS 'Campo para organizar el orden en el que se presentan los registros';
-- ddl-end --

-- object: fk_grupo_sanguineo_persona_persona | type: CONSTRAINT --
-- ALTER TABLE persona.grupo_sanguineo_persona DROP CONSTRAINT IF EXISTS fk_grupo_sanguineo_persona_persona CASCADE;
ALTER TABLE persona.grupo_sanguineo_persona ADD CONSTRAINT fk_grupo_sanguineo_persona_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_estado_civil_persona | type: CONSTRAINT --
-- ALTER TABLE persona.persona_estado_civil DROP CONSTRAINT IF EXISTS fk_persona_estado_civil_persona CASCADE;
ALTER TABLE persona.persona_estado_civil ADD CONSTRAINT fk_persona_estado_civil_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_estado_civil_estado_civil | type: CONSTRAINT --
-- ALTER TABLE persona.persona_estado_civil DROP CONSTRAINT IF EXISTS fk_persona_estado_civil_estado_civil CASCADE;
ALTER TABLE persona.persona_estado_civil ADD CONSTRAINT fk_persona_estado_civil_estado_civil FOREIGN KEY (estado_civil)
REFERENCES persona.estado_civil (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_genero_persona | type: CONSTRAINT --
-- ALTER TABLE persona.persona_genero DROP CONSTRAINT IF EXISTS fk_persona_genero_persona CASCADE;
ALTER TABLE persona.persona_genero ADD CONSTRAINT fk_persona_genero_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_genero_genero | type: CONSTRAINT --
-- ALTER TABLE persona.persona_genero DROP CONSTRAINT IF EXISTS fk_persona_genero_genero CASCADE;
ALTER TABLE persona.persona_genero ADD CONSTRAINT fk_persona_genero_genero FOREIGN KEY (genero)
REFERENCES persona.genero (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_grupo_etnico_grupo_etnico | type: CONSTRAINT --
-- ALTER TABLE persona.persona_grupo_etnico DROP CONSTRAINT IF EXISTS fk_persona_grupo_etnico_grupo_etnico CASCADE;
ALTER TABLE persona.persona_grupo_etnico ADD CONSTRAINT fk_persona_grupo_etnico_grupo_etnico FOREIGN KEY (grupo_etnico)
REFERENCES persona.grupo_etnico (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_grupo_etnico_persona | type: CONSTRAINT --
-- ALTER TABLE persona.persona_grupo_etnico DROP CONSTRAINT IF EXISTS fk_persona_grupo_etnico_persona CASCADE;
ALTER TABLE persona.persona_grupo_etnico ADD CONSTRAINT fk_persona_grupo_etnico_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_perfil_profesional_perfil_profesional | type: CONSTRAINT --
-- ALTER TABLE persona.persona_perfil_profesional DROP CONSTRAINT IF EXISTS fk_persona_perfil_profesional_perfil_profesional CASCADE;
ALTER TABLE persona.persona_perfil_profesional ADD CONSTRAINT fk_persona_perfil_profesional_perfil_profesional FOREIGN KEY (perfil_profesional)
REFERENCES persona.perfil_profesional (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_persona_perfil_profesional_persona | type: CONSTRAINT --
-- ALTER TABLE persona.persona_perfil_profesional DROP CONSTRAINT IF EXISTS fk_persona_perfil_profesional_persona CASCADE;
ALTER TABLE persona.persona_perfil_profesional ADD CONSTRAINT fk_persona_perfil_profesional_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_tipo_discapacidad_persona | type: CONSTRAINT --
-- ALTER TABLE persona.persona_tipo_discapacidad DROP CONSTRAINT IF EXISTS fk_persona_tipo_discapacidad_persona CASCADE;
ALTER TABLE persona.persona_tipo_discapacidad ADD CONSTRAINT fk_persona_tipo_discapacidad_persona FOREIGN KEY (persona)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_persona_tipo_discapacidad_tipo_discapacidad | type: CONSTRAINT --
-- ALTER TABLE persona.persona_tipo_discapacidad DROP CONSTRAINT IF EXISTS fk_persona_tipo_discapacidad_tipo_discapacidad CASCADE;
ALTER TABLE persona.persona_tipo_discapacidad ADD CONSTRAINT fk_persona_tipo_discapacidad_tipo_discapacidad FOREIGN KEY (tipo_discapacidad)
REFERENCES persona.tipo_discapacidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_relacion_personas_persona_principal | type: CONSTRAINT --
-- ALTER TABLE persona.relacion_personas DROP CONSTRAINT IF EXISTS fk_relacion_personas_persona_principal CASCADE;
ALTER TABLE persona.relacion_personas ADD CONSTRAINT fk_relacion_personas_persona_principal FOREIGN KEY (persona_principal)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_relacion_personas_persona_relacionada | type: CONSTRAINT --
-- ALTER TABLE persona.relacion_personas DROP CONSTRAINT IF EXISTS fk_relacion_personas_persona_relacionada CASCADE;
ALTER TABLE persona.relacion_personas ADD CONSTRAINT fk_relacion_personas_persona_relacionada FOREIGN KEY (persona_relacionada)
REFERENCES persona.persona (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_relacion_personas_tipo_relacion_personas | type: CONSTRAINT --
-- ALTER TABLE persona.relacion_personas DROP CONSTRAINT IF EXISTS fk_relacion_personas_tipo_relacion_personas CASCADE;
ALTER TABLE persona.relacion_personas ADD CONSTRAINT fk_relacion_personas_tipo_relacion_personas FOREIGN KEY (tipo_relacion_personas)
REFERENCES persona.tipo_relacion_personas (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --