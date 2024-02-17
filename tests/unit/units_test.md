You can also use cURL commands to test the endpoints from the command line. Here are example commands for each endpoint:

Sign Up:
curl -X POST -H "Content-Type: application/json" -d '{"name":"example", "email":"example@example.com", "password":"password"}' http://localhost:8080/signup

Sign In:
curl -X POST -H "Content-Type: application/json" -d '{"email":"example@example.com", "password":"password"}' http://localhost:8080/signin

Create Organization:
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"name":"Organization 1", "description":"Description of Organization 1"}' http://localhost:8080/organization

Invite User to Organization:
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"user_email":"user@example.com"}' http://localhost:8080/organization/{organization_id}/invite

Read Organization:
curl -X GET -H "Authorization: Bearer {token}" http://localhost:8080/organization/{organization_id}

Read All Organizations:
curl -X GET -H "Authorization: Bearer {token}" http://localhost:8080/organization

Update Organization:
curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"name":"Updated Name", "description":"Updated Description"}' http://localhost:8080/organization/{organization_id}

Delete Organization:
curl -X DELETE -H "Authorization: Bearer {token}" http://localhost:8080/organization/{organization_id}

Postman tests for the provided cURL commands, you can follow these steps:

    Sign Up:
        Create a new POST request in Postman.
        Set the request URL to http://localhost:8080/signup.
        Set the request body to {"name":"example", "email":"example@example.com", "password":"password"}.
        Add a test script to verify the response status code and body.

    Sign In:
        Create a new POST request in Postman.
        Set the request URL to http://localhost:8080/signin.
        Set the request body to {"email":"example@example.com", "password":"password"}.
        Add a test script to verify the response status code, body, and extract the access token.

    Create Organization:
        Create a new POST request in Postman.
        Set the request URL to http://localhost:8080/organization.
        Set the request headers with "Authorization: Bearer {token}", where {token} is the extracted access token from the Sign In response.
        Set the request body to {"name":"Organization 1", "description":"Description of Organization 1"}.
        Add a test script to verify the response status code and body.

    Invite User to Organization:
        Create a new POST request in Postman.
        Set the request URL to http://localhost:8080/organization/{organization_id}/invite.
        Set the request headers with "Authorization: Bearer {token}".
        Set the request body to {"user_email":"user@example.com"}.
        Add a test script to verify the response status code and body.

    Read Organization:
        Create a new GET request in Postman.
        Set the request URL to http://localhost:8080/organization/{organization_id}.
        Set the request headers with "Authorization: Bearer {token}".
        Add a test script to verify the response status code and body.

    Read All Organizations:
        Create a new GET request in Postman.
        Set the request URL to http://localhost:8080/organization.
        Set the request headers with "Authorization: Bearer {token}".
        Add a test script to verify the response status code and body.

    Update Organization:
        Create a new PUT request in Postman.
        Set the request URL to http://localhost:8080/organization/{organization_id}.
        Set the request headers with "Authorization: Bearer {token}".
        Set the request body to {"name":"Updated Name", "description":"Updated Description"}.
        Add a test script to verify the response status code and body.

    Delete Organization:
        Create a new DELETE request in Postman.
        Set the request URL to http://localhost:8080/organization/{organization_id}.
        Set the request headers with "Authorization: Bearer {token}".
        Add a test script to verify the response status code.

For each request, ensure to replace {token} and {organization_id} with the actual values obtained from previous requests or variables.