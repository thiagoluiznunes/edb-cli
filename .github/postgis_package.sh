#!/bin/bash

echo "Installing build-essential packages ..."
sudo apt install build-essential binutils lintian debhelper dh-make devscripts

echo "Check whether the package already exist ..."
apt-cache search package-name

echo "Preparing the package ..."
cd /home
mkdir postgisedb-0.1
cd postgisedb-0.1

echo "Download postgis ..."
wget http://postgis.net/stuff/postgis-3.3.4dev.tar.gz
tar -xvzf postgis-3.3.4dev.tar.gz
cd postgis-3.3.4dev

export DEBEMAIL="thiagoluiz.dev@gmail.com"
export DEBFULLNAME="Thiago Nunes"

dh_make --indep --createorig
