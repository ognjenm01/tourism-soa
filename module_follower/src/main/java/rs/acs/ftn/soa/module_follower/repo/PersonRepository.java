package rs.acs.ftn.soa.module_follower.repo;

import java.util.List;
import java.util.Optional;

import org.springframework.data.neo4j.repository.Neo4jRepository;
import org.springframework.data.neo4j.repository.query.Query;
import rs.acs.ftn.soa.module_follower.bean.Person;

public interface PersonRepository extends Neo4jRepository<Person, Long> {

    Optional<Person> findByName(String name);
    Optional<Person> findByUserId(Long userId);
    @Query("MATCH (a:Person {userId: $personId})-[:FOLLOWS]->(b:Person)-[:FOLLOWS]->(suggested:Person) " +
            "WHERE NOT (a)-[:FOLLOWS]->(suggested) AND suggested <> a " +
            "RETURN DISTINCT suggested LIMIT 3")
    List<Person> findSuggested(Long personId);
}