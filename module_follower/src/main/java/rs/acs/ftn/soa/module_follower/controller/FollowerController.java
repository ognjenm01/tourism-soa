package rs.acs.ftn.soa.module_follower.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import rs.acs.ftn.soa.module_follower.bean.Person;
import rs.acs.ftn.soa.module_follower.dto.FollowRequestDto;
import rs.acs.ftn.soa.module_follower.dto.PersonDto;
import rs.acs.ftn.soa.module_follower.service.FollowerService;

import java.util.List;

@RestController
public class FollowerController {
    FollowerService followerService;

    @Autowired
    public FollowerController(FollowerService followerService) {
        this.followerService = followerService;
    }

    @PostMapping("/api/reset")
    public HttpEntity<Void> resetDB(){
        return followerService.resetDB();
    }

    @GetMapping("/api/person")
    public HttpEntity<List<Person>> getPersons(){
        return followerService.findByUsers();
    }

    @PostMapping("/api/person")
    public HttpEntity<Person> createPerson(@RequestBody PersonDto personDto){
        return followerService.createPerson(personDto);
    }

    @GetMapping("/api/person/{id}")
    public HttpEntity<Person> getPerson(@PathVariable Long id){
        return followerService.findByUserID(id);
    }

    @GetMapping("/api/person/suggest/{id}")
    public HttpEntity<List<Person>> getSuggestion(@PathVariable Long id){
        return followerService.findSuggested(id);
    }

    @PostMapping("/api/person/follow")
    public HttpEntity<Person> followPerson(@RequestBody FollowRequestDto requestDto){
        return followerService.follow(requestDto);
    }

    @GetMapping("/api/person/follow")
    public HttpEntity<Boolean> checkAlreadyFollowed(@RequestBody FollowRequestDto requestDto){
        return followerService.checkIfFollowed(requestDto.getPersonId(), requestDto.getFollowerId());
    }

    @PostMapping("/api/person/unfollow")
    public HttpEntity<Person> unfollowPerson(@RequestBody FollowRequestDto requestDto) {
        return followerService.unfollow(requestDto);
    }

    @GetMapping("/api/person/followers/{id}")
    public ResponseEntity<List<Person>> findFollowers(@PathVariable long id){
        return followerService.findFollowers(id);
    }

    @GetMapping("/api/person/following/{id}")
    public ResponseEntity<List<Person>> findFollowing(@PathVariable long id){
        return followerService.findFollowing(id);
    }

    @GetMapping("/api/person/unfollowed/{id}")
    public ResponseEntity<List<Person>> getUnfollowed(@PathVariable Long id){
        return followerService.findUnfollowed(id);
    }
}
