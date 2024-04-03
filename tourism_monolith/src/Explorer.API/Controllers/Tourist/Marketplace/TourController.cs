using System.Text;
using System.Text.Json;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist.Marketplace;

[Authorize(Policy = "personPolicy")]
[AllowAnonymous]
[Route("api/marketplace/tours/")]
public class TourController : BaseApiController
{
    private readonly ITourService _tourService;
    private static readonly string _tourAppPort = Environment.GetEnvironmentVariable("TOURS_APP_PORT") ?? "8080";
    private static HttpClient httpToursClient = new()
    {
        BaseAddress = new Uri("http://tours-module:" + _tourAppPort + "/api/tours/"),
    };

    public TourController(ITourService tourService)
    {
        _tourService = tourService;
    }

    [HttpGet]
    public async Task<string> GetPublished([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _tourService.GetPublishedPaged(page, pageSize);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(new string[]{"ARCHIVED", "PUBLISHED"}), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpToursClient.PostAsync("bystatus",  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpGet("arhived-published")]
    [AllowAnonymous]
    public ActionResult<PagedResult<TourDto>> GetArchivedAndPublishedPaged([FromQuery] int page,
        [FromQuery] int pageSize)
    {
        var result = _tourService.GetArchivedAndPublishedPaged(page, pageSize);
        return CreateResponse(result);
    }
    
    [HttpGet("custom")]
    public ActionResult<PagedResult<TourDto>> GetCustomByUser([FromQuery] int page, [FromQuery] int pageSize)
    {
        var touristId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.GetCustomByUserPaged(touristId, page, pageSize);
        return CreateResponse(result);
    }

    [HttpGet("campaign")]
    public ActionResult<PagedResult<TourDto>> GetCampaignByUser([FromQuery] int page, [FromQuery] int pageSize)
    {
        var touristId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.GetCampaignByUserPaged(touristId, page, pageSize);
        return CreateResponse(result);
    }

    [HttpGet("author-published")]
    public ActionResult<PagedResult<TourDto>> GetPublishedByAuthor([FromQuery] int page, [FromQuery] int pageSize)
    {
        var authorId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.GetPublishedByAuthor(authorId, page, pageSize);
        return CreateResponse(result);
    }

    [HttpGet("allToursForAuthor/{id:int}")]
    [Authorize(Roles = "author")]
    public ActionResult<PagedResult<TourDto>> GetByAuthor(int page, int pageSize, int id)
    {
        var result = _tourService.GetByAuthor(page, pageSize, id);
        return CreateResponse(result);
    }
}