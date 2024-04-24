package rs.acs.ftn.soa.module_follower.mapper;

import org.mapstruct.Mapper;
import rs.acs.ftn.soa.module_follower.bean.Person;
import rs.acs.ftn.soa.module_follower.dto.PersonDto;

@Mapper(componentModel = "spring")
public interface PersonMapper {
    PersonDto sourceToDto(Person source);
    Person sourceToEntity(PersonDto source);
}
