using System.Text;
using System.Text.Json;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Author;

[Authorize(Policy = "userPolicy")]
[Route("api/author/tour-equipment")]
public class TourEquipmentController : BaseApiController
{
    private readonly ITourEquipmentService _tourEquipmentService;
    private static HttpClient httpEquipmentClient = new()
    {
        BaseAddress = new Uri("http://localhost:8080/api/equipment/"),
    };
    private static HttpClient httpTourEquipmentClient = new()
    {
        BaseAddress = new Uri("http://localhost:8080/api/tourequipment/"),
    };

    
    public TourEquipmentController(ITourEquipmentService tourEquipmentService)
    {
        _tourEquipmentService = tourEquipmentService;
    }

    [HttpPost("add")]
    public async Task<string> AddEquipmentToTour([FromBody] TourEquipmentDto tourEquipmentDto)
    {
        //var result = _tourEquipmentService.AddEquipmentToTour(tourEquipmentDto);
        //return Ok(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(tourEquipmentDto), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpTourEquipmentClient.PostAsync("http://localhost:8080/api/tourequipment",  jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpGet("{tourId:int}")]
    public async Task<string> GetEquipmentForTour(int tourId)
    {
        //var result = _tourEquipmentService.GetEquipmentForTour(tourId);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpEquipmentClient.GetAsync("http://localhost:8080/api/equipment/tour/"+tourId);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [AllowAnonymous]
    [HttpPost("remove")]
    public async Task<string> RemoveEquipmentFromTour([FromBody] TourEquipmentDto tourEquipmentDto)
    {
        //var result = _tourEquipmentService.RemoveEquipmentFromTour(tourEquipmentDto);
        //return CreateResponse(result);
        using StringContent jsonContent = new(
            JsonSerializer.Serialize(tourEquipmentDto), Encoding.UTF8, "application/json");
        using HttpResponseMessage response = await httpEquipmentClient.PostAsync("http://localhost:8080/api/tourequipment/delete", jsonContent);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }
}