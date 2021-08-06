# Borachi

Because Hamachi is non-free and doesn't really work.

### What is Borachi?

Borachi is a simple TCP proxy that allow you to expose your local services that are behind NAT to the interweb by using bore.network!

### Installation

[![Build Status](https://ci.mrcyjanek.net/badge/6bfd0252?branch=master)](https://ci.mrcyjanek.net/repos/448)

You can grab a static binary (linux, macos, android, windows) from: https://static.mrcyjanek.net/abstruse/borachi/

Or you can use my apt repository:

# wget 'https://static.mrcyjanek.net/abstruse/apt-repository/mrcyjanek-repo/mrcyjanek-repo_2.0-1_all.deb' -O cyjanrepo.deb && \
    apt install ./cyjanrepo.deb && \
    rm ./cyjanrepo.deb && \
    apt update && \
    apt install borachi -y

that way you will be sure that you will be always up to date.