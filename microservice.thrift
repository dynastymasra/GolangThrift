struct Person {
  1: required i32 ID,
  2: required string Firstname,
  3: optional string Lastname,
  4: optional string Email,
  5: required i16 Age,
  6: required bool Active
}

service PersonService {
  Person create(1:Person person),
  Person read(1:i32 id),
  Person update(1:Person person),
  void destroy(1:i32 id),
  list<Person> getAll()
}
