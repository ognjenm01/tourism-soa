package rs.acs.ftn.soa.module_follower.repo;

import java.util.List;
import java.util.Optional;
import java.util.Set;

import org.springframework.data.neo4j.repository.Neo4jRepository;
import org.springframework.data.neo4j.repository.query.Query;
import rs.acs.ftn.soa.module_follower.bean.Person;

public interface PersonRepository extends Neo4jRepository<Person, Long> {

    Optional<Person> findByName(String name);

    Optional<Person> findByUserId(Long userId);

    @Query("MATCH (a:Person {userId: $personId})-[:FOLLOWS]->(b:Person)-[:FOLLOWS]->(suggested:Person) " +
            "WHERE NOT (a)-[:FOLLOWS]->(suggested) AND suggested <> a " +
            "RETURN DISTINCT suggested")
    List<Person> findSuggested(Long personId);

    @Query("MATCH (u:Person)-[:FOLLOWS]->(follower:Person) " +
            "            WHERE follower.userId = $userId " +
            "            RETURN u")
    List<Person> findFollowers(Long userId);

    @Query("MATCH (p:Person) " +
            "WHERE p.userId <> $userId " +
            "AND NOT EXISTS { " +
            "    MATCH (u:Person {userId: $userId})-[:FOLLOWS]->(p) " +
            "} " +
            "RETURN p ")
    List<Person> findUnfollowed(Long userId);

}