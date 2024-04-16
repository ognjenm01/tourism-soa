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

	@Bean
	CommandLineRunner demo(PersonRepository personRepository) {
		return args -> {

			personRepository.deleteAll();

			Person greg = new Person(1, "Greg", "Gregovic", "greggy@gmail.com");
			Person roy = new Person(2, "Roy", "Royovic", "royke@gmail.com");
			Person craig = new Person(3, "Craig", "Cregovic", "creggy@gmail.com");

			personRepository.save(greg);
			personRepository.save(roy);
			personRepository.save(craig);

			greg = personRepository.findByName(greg.getName());
			if(greg == null) {
				greg = new Person(1, "Greg", "Gregovic", "greggy@gmail.com");
			}
			greg.follow(roy);
			greg.follow(craig);
			personRepository.save(greg);

			roy = personRepository.findByName(roy.getName());
			if(roy == null) {
				roy = new Person(2, "Roy", "Royovic", "royke@gmail.com");
			}
			roy.follow(craig);
			personRepository.save(roy);

			// We already know craig works with roy and greg

			List<Person> followers = personRepository.findByFollowersName(craig.getName());
			System.out.println("The following people follow Craig...");
			followers.stream().forEach(person -> System.out.println("\t" + person.getName()));
		};
	}
}
