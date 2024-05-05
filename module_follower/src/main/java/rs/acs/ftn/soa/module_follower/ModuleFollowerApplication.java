package rs.acs.ftn.soa.module_follower;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.data.neo4j.repository.config.EnableNeo4jRepositories;
import rs.acs.ftn.soa.module_follower.bean.Person;
import rs.acs.ftn.soa.module_follower.repo.PersonRepository;

import java.util.List;

@SpringBootApplication(exclude = {DataSourceAutoConfiguration.class})
@EnableNeo4jRepositories
public class ModuleFollowerApplication {

	public static void main(String[] args) {
		SpringApplication.run(ModuleFollowerApplication.class, args);
	}
}
