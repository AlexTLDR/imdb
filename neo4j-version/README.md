In order to start the docker db, run the below command:

docker run \
    --detach \
    --name Neo4j \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data \
    --volume=$HOME/neo4j/logs:/logs \
    --env NEO4J_dbms_memory_pagecache_size=4G \
    neo4j:latest

You can try out your Neo4j container by opening http://localhost:7474/ (the Neo4j’s Browser interface) in a web browser. 

After the docker container is created, navigate to the Built in guides -> :guide movie-graph to create the movie data base


mutation CreateActor{
  createActor(
    input: {name: "Alex99", email: "alex99@gmail.com", phone: "012345699"}
  ){
    _id
    name
    email
    phone
  }
}

mutation CreateMovie{
  createMovie(
    input: {
      actorId: "641da3004cf15a4e1f12879f",
      name: "Matrix8", 
      description: "SF", 
      status: IN_PROGRESS}
  ){
    actors{name, phone}
    name
    description
    status
  }
}

query GetAllActors{
  actors{
    name
    phone
    email
    _id
  }
}