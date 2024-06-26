﻿using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.Administration;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Administrator.Administration;

[Authorize(Policy = "userPolicy")]
[Route("api/administration/equipment")]
public class EquipmentController : BaseApiController
{
    private readonly IEquipmentService _equipmentService;
    private static readonly string _tourAppPort = Environment.GetEnvironmentVariable("TOURS_APP_PORT") ?? "8080";
    private static HttpClient httpEquipmentClient = new()
    {
        BaseAddress = new Uri("http://tours-module:" + _tourAppPort + "/api/equipment/"),
    };

    public EquipmentController(IEquipmentService equipmentService)
    {
        _equipmentService = equipmentService;
    }

    [HttpGet]
    [Authorize(Roles = "administrator,author,tourist")]
    public async Task<string> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _equipmentService.GetPaged(page, pageSize);
        //return CreateResponse(result);
        using HttpResponseMessage response = await httpEquipmentClient.GetAsync(httpEquipmentClient.BaseAddress);
        
        var jsonResponse = await response.Content.ReadAsStringAsync();
        return jsonResponse;
    }

    [HttpPost]
    [Authorize(Roles = "administrator")]
    public ActionResult<EquipmentDto> Create([FromBody] EquipmentDto equipment)
    {
        var result = _equipmentService.Create(equipment);
        return CreateResponse(result);
    }

    [HttpPut("{id:int}")]
    [Authorize(Roles = "administrator")]
    public ActionResult<EquipmentDto> Update([FromBody] EquipmentDto equipment)
    {
        var result = _equipmentService.Update(equipment);
        return CreateResponse(result);
    }

    [HttpDelete("{id:int}")]
    [Authorize(Roles = "administrator")]
    public ActionResult Delete(int id)
    {
        var result = _equipmentService.Delete(id);
        return CreateResponse(result);
    }
}