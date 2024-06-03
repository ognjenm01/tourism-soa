package rs.acs.ftn.soa.module_follower.controller;
/*import follower.*;
import io.grpc.stub.StreamObserver;
import org.springframework.beans.factory.annotation.Autowired;
import rs.acs.ftn.soa.module_follower.dto.PersonDto;
import rs.acs.ftn.soa.module_follower.service.FollowerService;

public class FollowerServiceController extends FollowerServiceGrpc.FollowerServiceImplBase {
    FollowerService followerService;

    //@Autowired
    FollowerServiceController(FollowerService followerService) {this.followerService = followerService;;}

    public PersonDto mapGrpcPersonToDto(Person person) {
        PersonDto personDto = new PersonDto();
        personDto.setId(person.getId());
        personDto.setName(person.getName());
        personDto.setSurname(person.getSurname());
        personDto.setEmail(person.getEmail());
        personDto.setUserId(person.getUserId());
        return personDto;
    }

    @Override
    public void create(Person request, StreamObserver<EmptyResponse> responseObserver) {
        followerService.createPerson(mapGrpcPersonToDto(request));
        responseObserver.onNext(EmptyResponse.newBuilder().build());
        responseObserver.onCompleted();
    }

    @Override
    public void getPerson(Id request, StreamObserver<PersonResponse> responseObserver) {
        responseObserver.onNext(PersonResponse.newBuilder().build());
        responseObserver.onCompleted();
    }
}*/

