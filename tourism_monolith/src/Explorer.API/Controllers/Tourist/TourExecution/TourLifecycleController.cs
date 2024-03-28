using System.Text;
using System.Text.Json;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourExecution;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist.TourExecution;

[Authorize(Policy = "touristPolicy")]
[Route("api/tourexecution/")]
public class TourLifecycleController : BaseApiController
{
    private readonly ITourLifecycleService _tourLifecycleService;
    private static readonly string _tourAppPort = Environment.GetEnvironmentVariable("TOURS_APP_PORT") ?? "8080";
    private static HttpClient httpPositionClient = new()
    {
        BaseAddress = new Uri("http://localhost:" + _tourAppPort + "/api/touristposition"),
    };

    public TourLifecycleController(ITourLifecycleService tourLifecycleService)
    {
        _tourLifecycleService = tourLifecycleService;
    }

    [HttpGet("activeTour")]
    public ActionResult<TourProgressDto> GetActiveByUser()
    {
        var result = _tourLifecycleService.GetActiveByUser(User.PersonId());
        return CreateResponse(result);
    }

    [HttpPost("start/{tourId:int}")]
    public ActionResult<TourProgressDto> StartTour([FromRoute] int tourId)
    {
        var result = _tourLifecycleService.StartTour(tourId, User.PersonId());
        return CreateResponse(result);
    }
    
    [HttpGet("position/{id:int}")]
    public async Task<string> GetPosition(int id)
    {
        using HttpResponseMessage response = await httpPositionClient.GetAsync("http://localhost:" + _tourAppPort + "/api/touristposition/byuser/"+id);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }
    
    [HttpPost("position")]
    public async Task<string> CreatePosition([FromBody] TouristPositionDto position)
    {
        position.UpdatedAt = DateTime.Now;
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(position), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpPositionClient.PostAsync(httpPositionClient.BaseAddress,  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }
    
    [HttpPut("position/{id:int}")]
    public async Task<string> UpdatePosition([FromBody] TouristPositionDto position, [FromRoute] int id)
    {
        position.UpdatedAt = DateTime.Now;
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(position), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpPositionClient.PutAsync(httpPositionClient.BaseAddress+"/"+id,  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    
    
    [HttpPut("abandonActive")]
    public ActionResult<TourProgressDto> AbandonActiveTour()
    {
        var result = _tourLifecycleService.AbandonActiveTour(User.PersonId());
        return CreateResponse(result);
    }

    [HttpPut("updateActive")]
    public ActionResult<TourProgressDto> UpdateActiveTour([FromBody] bool areRequiredEncountersDone)
    {
        var result = _tourLifecycleService.UpdateActiveTour(User.PersonId(), areRequiredEncountersDone);
        return CreateResponse(result);
    }
}