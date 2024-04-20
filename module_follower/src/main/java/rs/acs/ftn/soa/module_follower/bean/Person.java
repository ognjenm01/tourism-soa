package rs.acs.ftn.soa.module_follower.bean;


import com.fasterxml.jackson.annotation.JsonIgnore;
import org.springframework.data.neo4j.core.schema.GeneratedValue;
import org.springframework.data.neo4j.core.schema.Id;
import org.springframework.data.neo4j.core.schema.Node;
import org.springframework.data.neo4j.core.schema.Relationship;

import java.util.HashSet;
import java.util.Set;

@Node
public class Person {
    private @Id @GeneratedValue long id;

    private long userId;
    private String name;
    private String surname;
    private String email;

    @Relationship(type = "FOLLOWS")
    @JsonIgnore
    public Set<Person> followers;

    public Person() {
    }

    public Person(long userId, String name, String surname, String email) {
        this.userId = userId;
        this.name = name;
        this.surname = surname;
        this.email = email;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public long getUserId() {
        return userId;
    }

    public void setUserId(long userId) {
        this.userId = userId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSurname() {
        return surname;
    }

    public void setSurname(String surname) {
        this.surname = surname;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public Set<Person> getFollowers() {
        return followers;
    }

    public void setFollowers(Set<Person> followers) {
        this.followers = followers;
    }

    public boolean alreadyFollows(Person person) {
        if(followers == null) {
            followers = new HashSet<>();
        }
        return followers.contains(person);
    }

    public void follow(Person person) {
        if(followers == null)
            followers = new HashSet<>();
        followers.add(person);
    }

    public void unfollow(Person person) {
        if(followers == null)
            followers = new HashSet<>();
        followers.remove(person);
    }
}
