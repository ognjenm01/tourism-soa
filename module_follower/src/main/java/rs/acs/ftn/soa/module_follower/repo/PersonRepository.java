package rs.acs.ftn.soa.module_follower.repo;

import java.util.List;
import org.springframework.data.neo4j.repository.Neo4jRepository;
import rs.acs.ftn.soa.module_follower.bean.Person;

public interface PersonRepository extends Neo4jRepository<Person, Long> {

    Person findByName(String name);
    List<Person> findByFollowersName(String name);
}