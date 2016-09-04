# Tic Tac Toe Squared Server

A server for the Tic Tac Toe Squared (AKA Meta Tic Tac Toe)

This is written in Go

### Tech Stack:
 *  MySQL database to store game and user information
 *  Go server to interface with database and respond to requests
 *  (Another) Go server to handle authentication
 *  NGINX server to reverse proxy (so separate apps are separate processes)
 *  RESTful API as well as Websockets (to update when game changes)

## API
    URL | Function
    --- | --------
    GET /users/{UserID}/games | lists all games
    GET /games/{GameID}/info | lists info on game with specified GameID
    GET /games/{GameID}/board | gives board in JSON
    GET /games/{GameID}/string | gives board in string format (use monospaced font)
    GET /games/{GameID}/ws | gives websocket that emits on game changes and chats
    POST /games?Player1={PID1}&Player2={PID2} | makes a new game with specified GameID's and returns the GameID of the game created
    POST /games/{GameID}/move?Player={PID}&Box={BID}&Square={SID} | makes a move and responds with an error if unsucessful; broadcasts on ws if succesful 

### Request Authorization 
    Header | Value
    ------ | --------
    HMAC | encoded HMAC with SHA 256
    Encoding | encoding format for HMAC (if not provided, defaults to hex) 
    Time-Sent | seconds since epoch (fails if more than 10 seconds away from time received)

The HMAC uses (seconds in epoch):(path including initial / and without the T9) as the message and the login secret (in base 64 parsed as a string) as a secret.

####Example:

I want to get all the games for user 54689, which is at https://www.marktai.com/T9/users/54689/games

* Secret = "G7lc+B73JZnRQAYL0h15b4MvMZADllCleJtdJApSBJrq9x9kTm94FdOhnd31LttH995jbYeH7h5qP1W8WO1zng=="
* Path = /users/54689/games
* Time-Sent = 1472274995

This results in "1472274995:/users/54689/games" as the message, and the secret is "G7lc+B73JZnRQAYL0h15b4MvMZADllCleJtdJApSBJrq9x9kTm94FdOhnd31LttH995jbYeH7h5qP1W8WO1zng=="

Therefore, the HMAC is "7a2dc3e1fff05531a2a853c08e10878382e346b4793649e01ea37373d564b3b9", verifiable at http://www.freeformatter.com/hmac-generator.html


## Class Organization

* Game
    - Board (all 81 squares)
        + Box (9 squares)
            * Square (stored as uint)
    - MoveHistory
* DbGame - Game converts to dbgame before accessing database.  dbgame also converts to Game

