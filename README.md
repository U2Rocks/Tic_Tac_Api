# Tic Tac Toe Api

## Summary

This is an api that allows people to play tic tac toe games asynchronously

## Build

This backend uses **Fiber** to handle setup and manage api routes and **Gorm** to manage interactions with the sqlite 3 database

## Routes

- User Routes: **/ticgame/users/all**[GET] -> **/ticgame/users/add**[POST] -> **/ticgame/users/user**[POST] -> **/ticgame/users/update**[POST] -> **/ticgame/users/delete**[POST]
- Game Routes: **/ticgame/games/new**[POST] -> **/ticgame/delete**[POST] -> **/ticgame/games/turn**[POST]
- Board Routes: **/ticgame/boards/status**[POST]

## Final Comments and Notes

- The api returns all text messages to the client in a json object containing an statuscode/errorcode and a message
- The index html page is currently not being used
- BUG: The api currently does not know how to handle a tie(all spaces filled without a winner)
- Things to maybe add: route to see all game boards? route to list all active games?
