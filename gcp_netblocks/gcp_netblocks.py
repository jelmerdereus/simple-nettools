#!/usr/bin/env python3

#################################################################
# Google Cloud Netblock resolver - prints all subnets
#
# dependencies:
# pip3 install dnspython
#################################################################

import dns.resolver

def main():

    netblock_response = dns.resolver.query('_cloud-netblocks.googleusercontent.com', 'TXT').rrset

    netblock_names = [rec[8:] for rec in str(netblock_response).split(' ')[5:-1]]

    all_subnets = []

    for name in netblock_names:
        netblock_response = dns.resolver.query(name, 'TXT').rrset
        subnets = [net[4:] for net in str(netblock_response).split(' ')[5:-1]]
        all_subnets = all_subnets + subnets

    print(all_subnets)

if __name__ == '__main__':
    main()
