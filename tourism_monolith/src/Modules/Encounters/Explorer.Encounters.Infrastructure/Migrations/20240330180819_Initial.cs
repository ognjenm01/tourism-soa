using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Explorer.Encounters.Infrastructure.Migrations
{
    /// <inheritdoc />
    public partial class Initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.EnsureSchema(
                name: "encounters");

            migrationBuilder.CreateTable(
                name: "Encounters",
                schema: "encounters",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserId = table.Column<int>(type: "integer", nullable: false),
                    Name = table.Column<string>(type: "text", nullable: false),
                    Description = table.Column<string>(type: "text", nullable: false),
                    Latitude = table.Column<double>(type: "double precision", nullable: false),
                    Longitude = table.Column<double>(type: "double precision", nullable: false),
                    Xp = table.Column<int>(type: "integer", nullable: false),
                    Status = table.Column<int>(type: "integer", nullable: false),
                    Type = table.Column<int>(type: "integer", nullable: false),
                    Range = table.Column<double>(type: "double precision", nullable: false),
                    Image = table.Column<string>(type: "text", nullable: true),
                    ImageLatitude = table.Column<double>(type: "double precision", nullable: true),
                    ImageLongitude = table.Column<double>(type: "double precision", nullable: true),
                    PeopleCount = table.Column<int>(type: "integer", nullable: true),
                    ApprovalStatus = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Encounters", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "EncounterCompletions",
                schema: "encounters",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    LastUpdatedAt = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    EncounterId = table.Column<long>(type: "bigint", nullable: false),
                    Xp = table.Column<int>(type: "integer", nullable: false),
                    Status = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_EncounterCompletions", x => x.Id);
                    table.ForeignKey(
                        name: "FK_EncounterCompletions_Encounters_EncounterId",
                        column: x => x.EncounterId,
                        principalSchema: "encounters",
                        principalTable: "Encounters",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "KeypointEncounters",
                schema: "encounters",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    EncounterId = table.Column<long>(type: "bigint", nullable: false),
                    KeyPointId = table.Column<long>(type: "bigint", nullable: false),
                    IsRequired = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_KeypointEncounters", x => x.Id);
                    table.ForeignKey(
                        name: "FK_KeypointEncounters_Encounters_EncounterId",
                        column: x => x.EncounterId,
                        principalSchema: "encounters",
                        principalTable: "Encounters",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_EncounterCompletions_EncounterId",
                schema: "encounters",
                table: "EncounterCompletions",
                column: "EncounterId");

            migrationBuilder.CreateIndex(
                name: "IX_KeypointEncounters_EncounterId",
                schema: "encounters",
                table: "KeypointEncounters",
                column: "EncounterId",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "EncounterCompletions",
                schema: "encounters");

            migrationBuilder.DropTable(
                name: "KeypointEncounters",
                schema: "encounters");

            migrationBuilder.DropTable(
                name: "Encounters",
                schema: "encounters");
        }
    }
}
