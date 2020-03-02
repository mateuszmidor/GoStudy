# Architectural drivers

## Functions
* Allow fetching new comic books from online service
* Allow persisting the fetched comic books for further use offline
* Allow browsing previously fetched comic books offline
* Allow simple user-system interoperation thru clean command line interface

## Quality
* Should be ready for adding new comic book online sources in the future
* Should be ready for adding new offline storage methods in the future 
* Should be ready for adding new user interface types in the future
* Should be built in maximum of 3 seconds
* Should be deployed from code to running instance in maximum of 10 seconds
* Should only need to serve single user

## Constraints
* Should go through numerous changes as this is greenfield
* Should get MVP ready within 16 hours work
* Should be done by 1 junior developer

## Conventions
* Should be done in golang as this is what we use here

## Project goals
* This is a prototype - should disregard security, resilience, logging, error handling