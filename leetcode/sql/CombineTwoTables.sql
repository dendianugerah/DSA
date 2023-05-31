SELECT Person.firstName, Person.lastName, Address.city, Address.state 
FROM Address 
RIGHT JOIN Person 
ON Address.personId = Person.personId