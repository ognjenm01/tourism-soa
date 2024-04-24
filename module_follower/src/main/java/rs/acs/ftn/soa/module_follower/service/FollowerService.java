package rs.acs.ftn.soa.module_follower.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import rs.acs.ftn.soa.module_follower.bean.Person;
import rs.acs.ftn.soa.module_follower.dto.FollowRequestDto;
import rs.acs.ftn.soa.module_follower.dto.PersonDto;
import rs.acs.ftn.soa.module_follower.mapper.PersonMapperImpl;
import rs.acs.ftn.soa.module_follower.repo.PersonRepository;

import java.util.List;
import java.util.Optional;

@Service
public class FollowerService {

    private final PersonMapperImpl personMapperImpl;
    PersonRepository personRepository;

    @Autowired
    public FollowerService(PersonRepository personRepository, PersonMapperImpl personMapperImpl) {
        this.personRepository = personRepository;
        this.personMapperImpl = personMapperImpl;
    }

    public ResponseEntity<Void> resetDB(){
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            personRepository.deleteAll();
            return new ResponseEntity<>(responseHeaders, HttpStatus.OK);
        }
        catch(Exception e){
            System.out.println(e.getMessage());
            return new ResponseEntity<>(responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<Person> findByUserID(long id) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> person = personRepository.findByUserId(id);
            return person.map(value -> new ResponseEntity<>(value, responseHeaders, HttpStatus.OK)).orElseGet(() -> new ResponseEntity<>(HttpStatus.NOT_FOUND));
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<Person>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<List<Person>> findByUsers() {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            List<Person> person = personRepository.findAll();
            return new ResponseEntity<>(person, responseHeaders, HttpStatus.OK);
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<List<Person>>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<Person> findByName(String name) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> person = personRepository.findByName(name);
            return person.map(value -> new ResponseEntity<>(value, responseHeaders, HttpStatus.OK)).orElseGet(() -> new ResponseEntity<>(HttpStatus.NOT_FOUND));
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<Person>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<List<Person>> findSuggested(Long id) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            List<Person> suggested = personRepository.findSuggested(id);
            if(suggested == null) return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            else{
                return new ResponseEntity<>(suggested, responseHeaders, HttpStatus.OK);
            }
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<List<Person>> findFollowing(Long id) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            List<Person> followers = personRepository.findFollowers(id);
            if(followers== null) return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            else{
                return new ResponseEntity<>(followers, responseHeaders, HttpStatus.OK);
            }
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<List<Person>> findUnfollowed(Long id) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            List<Person> followers = personRepository.findUnfollowed(id);
            if(followers== null) return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            else{
                return new ResponseEntity<>(followers, responseHeaders, HttpStatus.OK);
            }
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<List<Person>> findFollowers(Long id) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> personOptional = personRepository.findByUserId(id);
            if(personOptional.isEmpty()) return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            else{
                Person person = personOptional.get();
                return new ResponseEntity<>(person.getFollowers().stream().toList(), responseHeaders, HttpStatus.OK);
            }
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<Boolean> checkIfFollowed(Long currentUserId, Long targetUserId) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> currentUserOpt = personRepository.findByUserId(currentUserId);
            Optional<Person> targetUserOpt = personRepository.findByUserId(targetUserId);
            if(currentUserOpt.isEmpty() || targetUserOpt.isEmpty()) return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            else{
                Person currentUser = currentUserOpt.get();
                Person targetUser = targetUserOpt.get();
                if(currentUser.alreadyFollowsByID(targetUser.getUserId())){
                    return new ResponseEntity<>(true, responseHeaders, HttpStatus.OK);
                }
                return new ResponseEntity<>(false, responseHeaders, HttpStatus.OK);
            }
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(null, responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    public ResponseEntity<Person> createPerson(PersonDto persondto) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            if(personRepository.findByUserId(persondto.getUserId()).isPresent()){
                return new ResponseEntity<>(responseHeaders, HttpStatus.CONFLICT);
            }
            Person person = personMapperImpl.sourceToEntity(persondto);
            personRepository.save(person);
            return new ResponseEntity<>(person, responseHeaders, HttpStatus.CREATED);
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }


    public ResponseEntity<Person> follow(FollowRequestDto requestDto) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> optionalPerson = personRepository.findByUserId(requestDto.getPersonId());
            Optional<Person> optionalFollower = personRepository.findByUserId(requestDto.getFollowerId());
            if(optionalPerson.isPresent() && optionalFollower.isPresent()){
                Person person = optionalPerson.get();
                Person follower = optionalFollower.get();
                if(person.alreadyFollows(follower))
                    return new ResponseEntity<>(responseHeaders, HttpStatus.CONFLICT);
                person.follow(follower);
                personRepository.save(person);
                return new ResponseEntity<>(person, responseHeaders, HttpStatus.OK);
            }
            return new ResponseEntity<>(responseHeaders, HttpStatus.NOT_FOUND);
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }


    public ResponseEntity<Person> unfollow(FollowRequestDto requestDto) {
        HttpHeaders responseHeaders = new HttpHeaders();
        try {
            Optional<Person> optionalPerson = personRepository.findByUserId(requestDto.getPersonId());
            Optional<Person> optionalFollower = personRepository.findByUserId(requestDto.getFollowerId());
            if(optionalPerson.isPresent() && optionalFollower.isPresent()){
                Person person = optionalPerson.get();
                Person follower = optionalFollower.get();
                if(!person.alreadyFollows(follower))
                    return new ResponseEntity<>(responseHeaders, HttpStatus.NOT_FOUND);
                person.unfollow(follower.getUserId());
                personRepository.save(person);
                return new ResponseEntity<>(person, responseHeaders, HttpStatus.OK);
            }
            return new ResponseEntity<>(responseHeaders, HttpStatus.NOT_FOUND);
        }
        catch (Exception e) {
            System.out.println(e.getMessage());
            return new ResponseEntity<>(responseHeaders, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

}