Feature: Validate API responses
  PERSONAS_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_relacion_personas
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                     |bodyreq               |codres       |
  |GET   |/v1/tipo_relacion_personas|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_relacion_persona |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_relacion_persona |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_relacion_persona |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_relacion_persona |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tipo_relacion_personas       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route                     |bodyreq               |codres         |bodyres                 |
  |GET   |/v1/tipo_relacion_personas|./files/req/Vacio.json|200 OK         |./files/res14/Vok1.json |
  |POST  |/v1/tipo_relacion_personas|./files/req/Yt1.json  |201 Created    |./files/res14/Vok2.json |
  |POST  |/v1/tipo_relacion_personas|./files/req/Vacio.json|400 Bad Request|./files/res14/Ierr1.json|
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt1.json  |400 Bad Request|./files/res14/Ierr2.json|
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt2.json  |400 Bad Request|./files/res14/Ierr3.json|
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt3.json  |400 Bad Request|./files/res14/Ierr4.json|
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt4.json  |400 Bad Request|./files/res14/Ierr5.json|
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt5.json  |400 Bad Request|./files/res14/Ierr6.json| 
  |POST  |/v1/tipo_relacion_personas|./files/req/Nt6.json  |400 Bad Request|./files/res14/Ierr7.json| 
  |POST  |/v1/tipo_relacion_personas|./files/req/Yt2.json  |400 Bad Request|./files/res14/Ierr8.json| 
  |PUT   |/v1/tipo_relacion_personas|./files/req/Yt2.json  |200 OK         |./files/res14/Vok2.json |
  |GETID |/v1/tipo_relacion_personas|./files/req/Vacio.json|200 OK         |./files/res14/Vok2.json |
  |DELETE|/v1/tipo_relacion_personas|./files/req/Vacio.json|200 OK         |./files/res14/Ino.json  |
  |DELETE|/v1/tipo_relacion_personas|./files/req/Vacio.json|404 Not Found  |./files/res14/Ierr9.json|
