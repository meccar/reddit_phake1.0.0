run Makefile to config system

how to run?


1- COPY this code then run it in terminal:

  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
  
  sudo apt-get update
  
  sudo apt-get install -y migrate
  

2- Run "make postgres" in terminal

3- Run "make network" in terminal

4- Run "docker ps -a" in terminal
   Copy "PORTS" of "postgres:16.2" 

5- Run "docker start THE_COPIED_PORTS"

6- Run "make createdb"

7- Run "make migrateup"

To run the server: "make server"
