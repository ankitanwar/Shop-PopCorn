## Authentication Service

This Service is Responsible For Verifying The Request 

<b>Auth Services</b>

| Feature  | Description  |
|----------|:-------------|
| login | To Get The AccessToken And Access Private Services |
| logout | To Delete the Access Token |
| verify | To Verify Whether The User Has Valid Access Token Or Not  |


<br></br>

<b>End Points</b>

| Request  | Description  | Url |
|----------|:-------------|:-------------|
| Post | To Request For Access Token - Login | host:8090/login |
| Get | To Get Verify The Access Token |host:8090/validate |
| Delete | To Delete The Access Token - Logout |host:8090/logout |

<br></br>