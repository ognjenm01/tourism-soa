using System.Text;
using System.Text.Json;
using Explorer.Blog.API.Dtos;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using Explorer.Stakeholders.Infrastructure.Authentication;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers;

[Authorize(Policy = "userPolicy")]
[Route("api/profile")]
public class ProfileController : BaseApiController
{
    private readonly IProfileService _profileService;
    private static readonly string _followerAppPort = Environment.GetEnvironmentVariable("FOLLOWERS_APP_PORT") ?? "4201";
    private static HttpClient httpFollowerClient = new()
    {
        BaseAddress = new Uri("http://followers-module:" + _followerAppPort),
    };

    public ProfileController(IProfileService profileService)
    {
        _profileService = profileService;
    }

    [HttpGet("{userId:int}")]
    public ActionResult<AccountRegistrationDto> GetStakeholderProfile(long userId)
    {
        var result = _profileService.GetProfile(userId);
        return CreateResponse(result);
    }
    
    [HttpGet("all")]
    public ActionResult<AccountRegistrationDto> GetAll()
    {
        var result = _profileService.GetAll();
        return CreateResponse(result);
    }
    
    [HttpGet("zelimdaumrem/{userId:int}")]
    public ActionResult<AccountRegistrationDto> GetPersonDto(long userId)
    {
        var result = _profileService.GetPersonDto(userId);
        return CreateResponse(result);
    }

    [AllowAnonymous]
    [HttpGet("userinfo/{id:int}")]
    public GoUserInfoDto GetUserInfoDto(int id)
    {
        var result = _profileService.GetUserInfoById(id);
        return result;
    }

    [HttpGet("not-followed")]
    public async Task<ActionResult<PersonDto[]>> GetNonFollowedProfiles([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _profileService.GetUserNonFollowedProfiles(page, pageSize, User.PersonId());
        //return CreateResponse(result);
        //http://localhost:4201/api/person/unfollowed/${id}
        long id = ClaimsPrincipalExtensions.PersonId(User);
        using HttpResponseMessage response = await httpFollowerClient.GetAsync("/api/person/unfollowed/"+id);
        var jsonResponse = await response.Content.ReadAsStringAsync();
        if (response.IsSuccessStatusCode)
            return CreateResponse(Result.Ok(jsonResponse));
        else
            return CreateResponse(Result.Fail(response.Content.ToString()));
    }

    [HttpGet("followers")]
    public async Task<ActionResult<PersonDto[]>> GetFollowers()
    {
        //var result = _profileService.GetFollowers(User.PersonId());
        //return CreateResponse(result);
        //http://localhost:4201/api/person/following/${id}
        long id = ClaimsPrincipalExtensions.PersonId(User);
        using HttpResponseMessage response = await httpFollowerClient.GetAsync("/api/person/following/"+id);
        var jsonResponse = await response.Content.ReadAsStringAsync();
        if (response.IsSuccessStatusCode)
            return CreateResponse(Result.Ok(jsonResponse));
        else
            return CreateResponse(Result.Fail(response.Content.ToString()));
    }

    [HttpGet("following")]
    public async Task<ActionResult<PersonDto[]>> GetFollowing()
    {
        //var result = _profileService.GetFollowing(User.PersonId());
        //return CreateResponse(result);
        long id = ClaimsPrincipalExtensions.PersonId(User);
        using HttpResponseMessage response = await httpFollowerClient.GetAsync("/api/person/followers/"+id);
        var jsonResponse = await response.Content.ReadAsStringAsync();
        if (response.IsSuccessStatusCode)
            return CreateResponse(Result.Ok(jsonResponse));
        else
            return CreateResponse(Result.Fail(response.Content.ToString()));
    }
    
    [HttpGet("suggestion")]
    public async Task<ActionResult<PersonDto[]>> GetSuggestion()
    {
        //var result = _profileService.GetFollowing(User.PersonId());
        //return CreateResponse(result);
        long id = ClaimsPrincipalExtensions.PersonId(User);
        using HttpResponseMessage response = await httpFollowerClient.GetAsync("/api/person/suggest/"+id);
        var jsonResponse = await response.Content.ReadAsStringAsync();
        if (response.IsSuccessStatusCode)
            return CreateResponse(Result.Ok(jsonResponse));
        else
        {
            Console.WriteLine(jsonResponse);
            return CreateResponse(Result.Fail(response.Content.ToString()));
        }
    }

    [HttpPut("{id:int}")]
    public ActionResult<PersonDto> Update(int id, [FromBody] PersonDto updatedPerson)
    {
        updatedPerson.Id = id;

        var result = _profileService.UpdateProfile(updatedPerson);
        return CreateResponse(result);
    }

    [HttpPost("follow")]
    public async Task<ActionResult<PagedResult<PersonDto>>> Follow([FromBody] FollowRequestDto request)
    {
        try
        {
            //var result = _profileService.Follow(User.PersonId(), followedId);
            //return CreateResponse(result);
            using StringContent jsonContent = new(
                JsonSerializer.Serialize(request), Encoding.UTF8, "application/json");
            using HttpResponseMessage response = await httpFollowerClient.PostAsync("/api/person/follow",  jsonContent);
            
            var jsonResponse = response.Content.ReadAsStringAsync();
            if (response.IsSuccessStatusCode)
                return CreateResponse(Result.Ok(jsonResponse));
            else
            {
                //Console.WriteLine(response.Content.);
                Console.WriteLine(jsonResponse);
                Console.WriteLine(response.Content.ToString());
                return CreateResponse(Result.Fail(response.Content.ToString()));
            }
        }
        catch (ArgumentException e)
        {
            return CreateResponse(Result.Fail(FailureCode.InvalidArgument).WithError(e.Message));
        }
    }

    [HttpPost("unfollow")]
    public async Task<ActionResult<PagedResult<PersonDto>>> Unfollow([FromBody] FollowRequestDto request)
    {
        try
        {
            //var result = _profileService.Unfollow(User.PersonId(), unfollowedId);
            //return CreateResponse(result);
            using StringContent jsonContent = new(
                JsonSerializer.Serialize(request), Encoding.UTF8, "application/json");
            using HttpResponseMessage response = await httpFollowerClient.PostAsync("/api/person/unfollow",  jsonContent);
            
            var jsonResponse = response.Content.ReadAsStringAsync();
            if (response.IsSuccessStatusCode)
                return CreateResponse(Result.Ok(jsonResponse));
            else
            {
                Console.WriteLine(jsonResponse);
                return CreateResponse(Result.Fail(response.Content.ToString()));
            }
        }
        catch (ArgumentException e)
        {
            return CreateResponse(Result.Fail(FailureCode.InvalidArgument).WithError(e.Message));
        }
    }

    [HttpGet("canCreateEncounters")]
    public ActionResult<bool> CanTouristCreateEncounters()
    {
        var result = _profileService.CanTouristCreateEncounters(ClaimsPrincipalExtensions.PersonId(User));
        return CreateResponse(result);
    }

}