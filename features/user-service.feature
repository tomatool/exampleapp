 Feature: external dependency  check

    Scenario: Check if /users/me endpoint return correct stuff
        Given set "user-service" with method "GET" and path "/example" response code to 200 and response body
            """
          {
                "user_id":"word"
          }
            """
        Given "mainapp-client" send request to "GET /users/me"
        Then "mainapp-client" response code should be 200
        Then "mainapp-client" response body should equal
        """
        {
            "username":"word"
        }
        """