Feature: Validate API responses
  PERSONAS_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /persona_estado_civil
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      

    Examples: 
    |method |route                    |bodyreq                 |codres       |
    |GET    |/v1/persona_estado_civil |./files/req/Vacio.json  |200 OK       |
    |GET    |/v1/persona_estado_civi  |./files/req/Vacio.json  |404 Not Found|
    |POST   |/v1/persona_estado_civi  |./files/req/Vacio.json  |404 Not Found|
    |PUT    |/v1/persona_estado_civi  |./files/req/Vacio.json  |404 Not Found|
    |DELETE |/v1/persona_estado_civi  |./files/req/Vacio.json  |404 Not Found|

Scenario Outline: To probe response route /persona_estado_civil       
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method |route                     |bodyreq                |codres     |bodyres                         |
    |POST   |/v1/persona_estado_civil  |./files/req/Vacio.json |200 OK     |./files/res0/Ierr6.json         |
    
    