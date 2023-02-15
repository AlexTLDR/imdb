1. Install docker and configure a Dockerfile:

    create a directory called mysql-docker in the project's directory
    copy the sql-scripts directory in the mysql-docker folder
    copy the Dockerfile in the mysql-docker folder

# Or use the shared image that I already created and skip to step 3. In the terminal, type -> docker pull alextldr/imdb-sql:latest not yet implemented

2. Create an image from the docker file:

    open a terminal from the mysql-docker folder
    run this command to build the image -> docker build -t imdb-sql .

3. Start the container from the image:

    in the same terminal from the mysql-docker folder, run -> docker run -d -p 3306:3306 --name imdb-sql -e MYSQL_ROOT_PASSWORD=123456 imdb-sql

    In order to restart the container, just run -> docker start imdb-sql
