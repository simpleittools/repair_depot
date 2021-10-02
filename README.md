#Repair Depot
This application is focused on building a central place to store all repair info needed for an in-depot shop.

To accomplis this goal I will need to set some basic rules.
1. How will I manage this project
2. Define our application goals
3. Define our backend
4. Define our frontend
5. Build our API communication
6. Create our Database Schema

##How will I manage this project
To keep myself on task, I will utilize the projects section of Github.com
For the sake of simplicity, I will generate 3 sections
1. Must Have
2. Need
3. Would like to have

Everything starts in the "Would like to have", then moves to the "Need", from there gets set to the "Must Have." Since I am
working on this project alone, I will never have more than 3 in the "Must Have", or I will lose track.

##Application Goals
The hope of a repair depot is to be quite busy. The problem in a repair depot is you get quite busy.
The result of all this busy activity is, you can lose track of where every individual "project" is. The goal will be to 
create a central reporting system to collect every possible part of the process.
1. Client report
2. Client dropoff
3. Defined Client Expectations
4. Items delivered by client
5. tasks and activities needed to complete the repair
6. parts
   1. identifying
   2. quoting
   3. approval
   4. purchasing
   5. delivery
   6. installation
7. Updating the client
8. Finalizing return to the client
9. Historical tracking for future data retrieval
10. Knowledge Base (this may become its own project as this can become quite involved)

#Define our backend
The backed will be coded in Go with GoFiber. For now, Go/GoFiber will also manage the frontend and templating. This will
ultimately be moved into a Single Page Application, but a more basic HTML/CSS template for user input/interactivity will be sufficient to start.

The usage of a SQL database seems to make the most sense due to the intricate relationships that will be needed. Though an ORM will be utilized
(Gorm) this system will initially be built for PostgreSQL. More DBMS integrations will be available at a future time.

