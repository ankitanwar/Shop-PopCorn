## User Service 

This Service Is Responsible To Create The User Account And Add User Info To The Database

<b>User Service</b>

| Feature  | Description  |
|----------|:-------------|
| CreateAccount | Ability To Create A New Account |
| DeleteAccount | Ability To Delete Account From The System |
| GetDetails | Ability To View The Details Saved In The system |
| Update Details | Ability To Update The Saved Details In The System |
| Add Address | Ability To Add Address |
| Get Address | Ability To See All The  Addresses |

<br></br>

<b>End Points</b>

| Request  | Description  | Url |
|----------|:-------------|:-------------|
| Post | To Register New User | host:8081/users |
| Get | To Get The Details Of The User |host:8081/users |
| Patch | To Update User Details |host:8081/store/users |
| Delete | To Delete The User Account |host:8081/users |
| Get | To Get The User Addresses |host:8081/user/address |
| Post | To Add The User Address |host:8081/user/address |
| Post | To Verify The User Credentials |host:8081/user/verify |
| Get | To Get The User Specific Address |host:8081/user/specificaddress/:addressID |
| Delete | To Delete The User Address |host:8081/user/address/:addressID|

<br></br>