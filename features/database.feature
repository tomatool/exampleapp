 Feature: database endpoint

    Scenario: Check database endpoint
        Given "mainapp-client" send request to "GET /users"
        Then "mainapp-client" response code should be 200
        Then "mainapp-client" response body should equal
            """
            {"error": "pq: relation \"users\" does not exist"}
            """