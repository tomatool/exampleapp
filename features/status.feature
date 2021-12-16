 Feature: status endpoint

    Scenario: Check base endpoint
        Given "mainapp-client" send request to "GET /"
        Then "mainapp-client" response code should be 200
        Then "mainapp-client" response body should equal
            """
            {
                "name":"mainapp",
                "version":"v0.1.0"
            }
            """