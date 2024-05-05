package rs.acs.ftn.soa.module_follower.dto;

public class FollowRequestDto {
    private long personId;
    private long followerId;

    public FollowRequestDto() {
    }

    public FollowRequestDto(long personId, long followerId) {
        this.personId = personId;
        this.followerId = followerId;
    }

    public long getPersonId() {
        return personId;
    }

    public void setPersonId(long personId) {
        this.personId = personId;
    }

    public long getFollowerId() {
        return followerId;
    }

    public void setFollowerId(long followerId) {
        this.followerId = followerId;
    }
}
