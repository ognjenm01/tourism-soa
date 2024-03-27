using System.Text;
using System.Text.Json;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Author;

[Authorize(Policy = "userPolicy")]
[Route("api/author/tours/")]
public class TourManagementController : BaseApiController
{
    private readonly ITourService _tourService;
    private static HttpClient httpTourClient = new()
    {
        BaseAddress = new Uri("http://localhost:8080/api/tours/"),
    };

    public TourManagementController(ITourService tourService)
    {
        _tourService = tourService;
    }

    [HttpGet]
    [Authorize(Roles = "author, tourist")]
    public async Task<string> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _tourService.GetPaged(page, pageSize);
        //return CreateResponse(result);
        //var result = _tourService.Get(tourId);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpTourClient.GetAsync("http://localhost:8080/api/tours");
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [AllowAnonymous]
    [HttpGet("{tourId:int}")]
    [Authorize(Roles = "author")]
    public async Task<string> GetById([FromRoute] int tourId)
    {
        //var result = _tourService.Get(tourId);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpTourClient.GetAsync(tourId.ToString());
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPost]
    [Authorize(Roles = "author")]
    public async Task<string> Create([FromBody] TourDto tour)
    {
        //tour.UserId = User.PersonId();
        //var result = _tourService.Create(tour);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(tour), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpTourClient.PostAsync("http://localhost:8080/api/tours",  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPut("{id:int}")]
    [Authorize(Roles = "author,tourist")]
    public async Task<string> Update([FromBody] TourDto tour, [FromRoute] int id)
    {
        //tour.UserId = User.PersonId();
        //var result = _tourService.Update(tour);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(tour), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpTourClient.PutAsync(httpTourClient.BaseAddress+"/"+id,  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }
    
    [HttpDelete("{id:int}")]
    [Authorize(Roles = "author")]
    public ActionResult Delete(int id)
    {
        var result = _tourService.Delete(id);
        return CreateResponse(result);
    }

    [HttpGet("byauthor/{id:int}")]
    [Authorize(Roles = "author")]
    public async Task<string> GetByAuthor([FromRoute] int id)
    {
        //var authorId = User.PersonId();
        //var result = _tourService.GetByAuthor(page, pageSize, authorId);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpTourClient.GetAsync("byauthor/"+id);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [AllowAnonymous]
    [HttpPut("disable/{id:int}")]
    public ActionResult<TourDto> Disable([FromBody] TourDto tour)
    {
        if (User.IsInRole("administrator"))
        {
            tour.UserId = User.PersonId();
            var result = _tourService.Update(tour);
            return CreateResponse(result);
        }
        return null;
    }
    
    [HttpPost("custom")]
    public ActionResult<TourDto> CreateCustomTour([FromBody] TourDto tourDto)
    {   
        tourDto.UserId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.CreateCustom(tourDto);
        return CreateResponse(result);
    }

    [HttpPost("campaign")]
    public ActionResult<TourDto> CreateCampaignTour([FromBody] TourDto tourDto) 
    {
        tourDto.UserId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.CreateCampaign(tourDto);
        return CreateResponse(result);
    }
}