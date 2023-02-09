
SELECT Person.firstname, Person.lastName,Address.city, Address.state
FROM Person
Left JOIN Address
on Person.personId=Address.personId


