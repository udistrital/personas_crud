-- inserts persona.perfil_profesional
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1,'Bachiller','Perfil Bachiller','BACH', true,1, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (2,'Técnico','Perfil Técnico','TEC', true,2, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (3,'Tecnólogo','Perfil Tecnólogo','TECN', true,3, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (4,'Licenciado','Perfil Licenciado','LIC', true,4, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (5,'Profesional','Perfil Profesional','PROF', true,5, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (6,'Universitario','Perfil Universitario','UNI', true,6, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (7,'Especializado','Perfil Especializado','ESP', true,7, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (8,'Certificado','Perfil Certificado','CER', true,8, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (9,'Máster/Magister','Perfil Máster/Magister','MAS', true,9, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (10,'Doctorado','Perfil Doctorado','DOC', true,10, now(), now());
INSERT INTO persona.perfil_profesional(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (11,'Postdoctorado','Perfil Postdoctorado','PDOC', true,11, now(), now());

-- inserts persona.grupo_etnico
INSERT INTO persona.grupo_etnico(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1,'Pueblo indígena','Pueblo indígena','INDIGENA', true,1, now(), now());
INSERT INTO persona.grupo_etnico(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (2,'Población negra o afrocolombiana','Población negra o afrocolombiana','AFRO', true,2, now(), now());
INSERT INTO persona.grupo_etnico(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (3,'Población raizal','Población raizal','RAIZAL', true,3, now(), now());
INSERT INTO persona.grupo_etnico(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (4,'Pueblo rom (gitano)','Pueblo rom','ROM', true,4, now(), now());
INSERT INTO persona.grupo_etnico(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (5,'Pueblo palenquero','Pueblo palenquero','PALENQUE', true,5, now(), now());

-- inserts persona.genero
INSERT INTO persona.genero(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (1,'Masculino','Corresponde al género masculino','M', true,1, now(), now());
INSERT INTO persona.genero(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (2,'Femenino','Corresponde al género femenino','F', true,2, now(), now());
INSERT INTO persona.genero(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (3,'Otro','Otro genero','Otro', true,3, now(), now());

-- inserts persona.estado_civil
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1,'Soltero(a)','Soltero(a)','SOLTERO', true,1, now(), now());
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (2,'Casado(a)','Casado(a)','CASADO', true,2, now(), now());
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (3,'Divorciado(a)','Divorciado(a)','DIVORCIADO', true,3, now(), now());
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (4,'Viudo(a)','Viudo(a)','VIUDO', true,4, now(), now());
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (5,'Unión libre','Unión libre','UL', true,5, now(), now());
INSERT INTO persona.estado_civil(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (6,'Separado(a)','Separado(a)','SEPARADO', true,6, now(), now());

-- inserts persona.tipo_discapacidad
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (1,'Sordera profunda','Sordera profunda','SP', true,1, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (2,'Hipoacusia o baja audición','Hipoacusia o baja audición','HBA', true,2, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (3,'Autismo','Autismo','A', true,3, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (4,'Baja visión diagnóstica','Baja visión diagnóstica','BVD', true,4, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (5,'Ceguera','Ceguera','C', true,5, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (6,'Dificultad sensorial en gusto, olfato y/o tacto','Dificultad sensorial en gusto, olfato y/o tacto','DSGOT', true,6, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (7,'Discapacidad física','Discapacidad física','DF', true,7, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (8,'Lesión neuromuscular','Lesión neuromuscular','LNM', true,8, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (9,'Deficiencia de movilidad','Deficiencia de movilidad','DM', true,9, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (10,'Deficiencia de voz y/o habla','Deficiencia de voz y/o habla','DVH', true,10, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (11,'Parálisis cerebral','Parálisis cerebral','PC', true,11, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (12,'Deficiencia cognitiva','Deficiencia cognitiva','DC', true,12, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (13,'Deficiencia mental y/o psicosocial','Deficiencia mental y/o psicosocial','DPS', true,13, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (14,'Síndrome de Down','Síndrome de Down','SD', true,14, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (15,'Enfermedades crónicas','Enfermedades crónicas','EC', true,15, now(), now());
INSERT INTO persona.tipo_discapacidad(id,nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES (16,'Deficiencia sistémica','Deficiencia sistémica','DS', true,16, now(), now());

-- inserts persona.tipo_relacion_personas
INSERT INTO persona.tipo_relacion_personas(nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES ('Hijo(a)','Padre/Madre - Hijo(a)','HI', true,1, now(), now());
INSERT INTO persona.tipo_relacion_personas(nombre, descripcion, codigo_abreviacion, activo, numero_orden,fecha_creacion, fecha_modificacion) VALUES ('Conyuge','Conyuge','CY', true,2, now(), now());