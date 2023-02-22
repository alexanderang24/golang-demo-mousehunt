# golang-demo-mousehunt
 
MouseHunt: A Passive Hunting Game

MouseHunt is a passive game where players can hunt mice, collect golds rewards. Gold can be used to travel to another location where different breed of mice - and certainly more powerful- exists. To be able to catch those mice, you will need more powerful trap upgrades, that can be bought at the nearest Trapsmith in your location.
<br><br>

Base URL: `https://golang-demo-mousehunt-production.up.railway.app`

All APIs are authenticated by JWT, except login and register.

To use the API, first you need to register a user using User -> Register API (will be registered with role "player") or use the existing user as follows:

| **username** | **password** | **role** |
| --- | --- | --- |
| admin | admin | admin |
| john | password | player |

Then, you need to login using User -> Login API to get the JWT token that will be used by all APIs.

<br>
Collection's request folder are arranged as follows:

- Player API: all API that can be used by players once they are logged in.
- User, Trap, Location, Mouse: CRUD APIs that are used by admin to manage game data.


<img src="https://content.pstmn.io/fdc2c9e4-e24d-4f9a-a76e-b54967987af9/Zm9sZGVycy5wbmc=" alt="Collection's%20folder" width="164" height="218">
<br><br>

All requests are complemented by examples that cover all possible responses. You can choose example using dropdown menu on the top-right side of the right column.

<img src="https://content.pstmn.io/7e9b191d-5c36-426b-8848-58c613d2b79b/ZXhhbXBsZS5wbmc=" alt="Request%20example">