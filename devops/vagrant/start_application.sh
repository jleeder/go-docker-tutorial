#!/bin/bash

cd /vagrant/boardy
sudo docker build -t boardy_image:latest .
sudo docker stop boardy
sudo docker rm boardy
sudo docker run -d --name boardy -p 8080:8080 boardy_image:latest