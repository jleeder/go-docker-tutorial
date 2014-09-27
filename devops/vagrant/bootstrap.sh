#!/bin/bash

export DEBIAN_FRONTEND=noninteractive

apt-get update

#install docker
apt-get -y install docker.io

ln -sf /usr/bin/docker.io /usr/local/bin/docker
sed -i '$acomplete -F _docker docker' /etc/bash_completion.d/docker.io

#run docker at boot
update-rc.d docker.io defaults

#install supervisor
apt-get -y install supervisor
service supervisor restart

echo "IP: $(ifconfig eth0 | grep -o 'inet addr:[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}' | cut -d: -f2)"

ln -sf /vagrant/devops/vagrant/start_application.sh /usr/local/bin/runserver
chmod 744 /vagrant/devops/vagrant/start_application.sh