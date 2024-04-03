using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Explorer.Stakeholders.Infrastructure.Migrations
{
    /// <inheritdoc />
    public partial class Initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.EnsureSchema(
                name: "stakeholders");

            migrationBuilder.CreateTable(
                name: "Achievements",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Name = table.Column<string>(type: "text", nullable: false),
                    Description = table.Column<string>(type: "text", nullable: false),
                    Type = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Achievements", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "ApplicationRatings",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Rating = table.Column<int>(type: "integer", nullable: false),
                    Comment = table.Column<string>(type: "text", nullable: false),
                    UserId = table.Column<int>(type: "integer", nullable: false),
                    LastModified = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    IsRated = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ApplicationRatings", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "ClubInvitations",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    ClubId = table.Column<long>(type: "bigint", nullable: false),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    Status = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ClubInvitations", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "ClubJoinRequests",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    ClubId = table.Column<long>(type: "bigint", nullable: false),
                    Status = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ClubJoinRequests", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Notifications",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserId = table.Column<int>(type: "integer", nullable: false),
                    Type = table.Column<int>(type: "integer", nullable: false),
                    Content = table.Column<string>(type: "text", nullable: true),
                    ActionURL = table.Column<string>(type: "text", nullable: true),
                    CreationDateTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    IsRead = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Notifications", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Users",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Username = table.Column<string>(type: "text", nullable: false),
                    Password = table.Column<string>(type: "text", nullable: false),
                    Role = table.Column<int>(type: "integer", nullable: false),
                    IsActive = table.Column<bool>(type: "boolean", nullable: false),
                    Email = table.Column<string>(type: "text", nullable: false),
                    IsBlocked = table.Column<bool>(type: "boolean", nullable: false),
                    IsEnabled = table.Column<bool>(type: "boolean", nullable: false),
                    VerificationToken = table.Column<string>(type: "text", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Users", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "NewsletterPreferences",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserID = table.Column<long>(type: "bigint", nullable: false),
                    Frequency = table.Column<long>(type: "bigint", nullable: false),
                    LastSent = table.Column<DateTime>(type: "timestamp with time zone", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_NewsletterPreferences", x => x.Id);
                    table.ForeignKey(
                        name: "FK_NewsletterPreferences_Users_UserID",
                        column: x => x.UserID,
                        principalSchema: "stakeholders",
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "TourIssue",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Category = table.Column<string>(type: "text", nullable: false),
                    Priority = table.Column<long>(type: "bigint", nullable: false),
                    Description = table.Column<string>(type: "text", nullable: false),
                    CreationDateTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    ResolveDateTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: true),
                    IsResolved = table.Column<bool>(type: "boolean", nullable: false),
                    TourId = table.Column<long>(type: "bigint", nullable: false),
                    UserId = table.Column<long>(type: "bigint", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TourIssue", x => x.Id);
                    table.ForeignKey(
                        name: "FK_TourIssue_Users_UserId",
                        column: x => x.UserId,
                        principalSchema: "stakeholders",
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "TourIssueComments",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Comment = table.Column<string>(type: "text", nullable: false),
                    CreationDateTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    TourIssueId = table.Column<long>(type: "bigint", nullable: false),
                    UserId = table.Column<long>(type: "bigint", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TourIssueComments", x => x.Id);
                    table.ForeignKey(
                        name: "FK_TourIssueComments_TourIssue_TourIssueId",
                        column: x => x.TourIssueId,
                        principalSchema: "stakeholders",
                        principalTable: "TourIssue",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_TourIssueComments_Users_UserId",
                        column: x => x.UserId,
                        principalSchema: "stakeholders",
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "AchievementClub",
                schema: "stakeholders",
                columns: table => new
                {
                    AchievementsId = table.Column<long>(type: "bigint", nullable: false),
                    ClubId = table.Column<long>(type: "bigint", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_AchievementClub", x => new { x.AchievementsId, x.ClubId });
                    table.ForeignKey(
                        name: "FK_AchievementClub_Achievements_AchievementsId",
                        column: x => x.AchievementsId,
                        principalSchema: "stakeholders",
                        principalTable: "Achievements",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "ChatMessages",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    SenderId = table.Column<long>(type: "bigint", nullable: false),
                    ReceiverId = table.Column<long>(type: "bigint", nullable: false),
                    Content = table.Column<string>(type: "text", nullable: false),
                    CreationDateTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    IsRead = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ChatMessages", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "ClubChallengeRequests",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    ChallengerId = table.Column<long>(type: "bigint", nullable: false),
                    ChallengedId = table.Column<long>(type: "bigint", nullable: false),
                    Status = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ClubChallengeRequests", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "ClubFights",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    WinnerId = table.Column<int>(type: "integer", nullable: true),
                    StartOfFight = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    EndOfFight = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    Club1Id = table.Column<long>(type: "bigint", nullable: false),
                    Club2Id = table.Column<long>(type: "bigint", nullable: false),
                    IsInProgress = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ClubFights", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Clubs",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    Name = table.Column<string>(type: "text", nullable: false),
                    Description = table.Column<string>(type: "text", nullable: true),
                    Image = table.Column<string>(type: "text", nullable: true),
                    OwnerId = table.Column<long>(type: "bigint", nullable: false),
                    FightsWon = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Clubs", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "People",
                schema: "stakeholders",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    Name = table.Column<string>(type: "text", nullable: false),
                    Surname = table.Column<string>(type: "text", nullable: false),
                    Email = table.Column<string>(type: "text", nullable: false),
                    ProfileImage = table.Column<string>(type: "text", nullable: false),
                    Biography = table.Column<string>(type: "text", nullable: false),
                    Quote = table.Column<string>(type: "text", nullable: false),
                    XP = table.Column<int>(type: "integer", nullable: false),
                    Level = table.Column<int>(type: "integer", nullable: false),
                    ClubId = table.Column<long>(type: "bigint", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_People", x => x.Id);
                    table.ForeignKey(
                        name: "FK_People_Clubs_ClubId",
                        column: x => x.ClubId,
                        principalSchema: "stakeholders",
                        principalTable: "Clubs",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.SetNull);
                    table.ForeignKey(
                        name: "FK_People_Users_UserId",
                        column: x => x.UserId,
                        principalSchema: "stakeholders",
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "PersonPerson",
                schema: "stakeholders",
                columns: table => new
                {
                    FollowersId = table.Column<long>(type: "bigint", nullable: false),
                    FollowingId = table.Column<long>(type: "bigint", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_PersonPerson", x => new { x.FollowersId, x.FollowingId });
                    table.ForeignKey(
                        name: "FK_PersonPerson_People_FollowersId",
                        column: x => x.FollowersId,
                        principalSchema: "stakeholders",
                        principalTable: "People",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_PersonPerson_People_FollowingId",
                        column: x => x.FollowingId,
                        principalSchema: "stakeholders",
                        principalTable: "People",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_AchievementClub_ClubId",
                schema: "stakeholders",
                table: "AchievementClub",
                column: "ClubId");

            migrationBuilder.CreateIndex(
                name: "IX_ChatMessages_ReceiverId",
                schema: "stakeholders",
                table: "ChatMessages",
                column: "ReceiverId");

            migrationBuilder.CreateIndex(
                name: "IX_ChatMessages_SenderId",
                schema: "stakeholders",
                table: "ChatMessages",
                column: "SenderId");

            migrationBuilder.CreateIndex(
                name: "IX_ClubChallengeRequests_ChallengedId",
                schema: "stakeholders",
                table: "ClubChallengeRequests",
                column: "ChallengedId");

            migrationBuilder.CreateIndex(
                name: "IX_ClubChallengeRequests_ChallengerId",
                schema: "stakeholders",
                table: "ClubChallengeRequests",
                column: "ChallengerId");

            migrationBuilder.CreateIndex(
                name: "IX_ClubFights_Club1Id",
                schema: "stakeholders",
                table: "ClubFights",
                column: "Club1Id");

            migrationBuilder.CreateIndex(
                name: "IX_ClubFights_Club2Id",
                schema: "stakeholders",
                table: "ClubFights",
                column: "Club2Id");

            migrationBuilder.CreateIndex(
                name: "IX_Clubs_OwnerId",
                schema: "stakeholders",
                table: "Clubs",
                column: "OwnerId");

            migrationBuilder.CreateIndex(
                name: "IX_NewsletterPreferences_UserID",
                schema: "stakeholders",
                table: "NewsletterPreferences",
                column: "UserID",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_People_ClubId",
                schema: "stakeholders",
                table: "People",
                column: "ClubId");

            migrationBuilder.CreateIndex(
                name: "IX_People_UserId",
                schema: "stakeholders",
                table: "People",
                column: "UserId",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_PersonPerson_FollowingId",
                schema: "stakeholders",
                table: "PersonPerson",
                column: "FollowingId");

            migrationBuilder.CreateIndex(
                name: "IX_TourIssue_UserId",
                schema: "stakeholders",
                table: "TourIssue",
                column: "UserId");

            migrationBuilder.CreateIndex(
                name: "IX_TourIssueComments_TourIssueId",
                schema: "stakeholders",
                table: "TourIssueComments",
                column: "TourIssueId");

            migrationBuilder.CreateIndex(
                name: "IX_TourIssueComments_UserId",
                schema: "stakeholders",
                table: "TourIssueComments",
                column: "UserId");

            migrationBuilder.CreateIndex(
                name: "IX_Users_Username",
                schema: "stakeholders",
                table: "Users",
                column: "Username",
                unique: true);

            migrationBuilder.AddForeignKey(
                name: "FK_AchievementClub_Clubs_ClubId",
                schema: "stakeholders",
                table: "AchievementClub",
                column: "ClubId",
                principalSchema: "stakeholders",
                principalTable: "Clubs",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ChatMessages_People_ReceiverId",
                schema: "stakeholders",
                table: "ChatMessages",
                column: "ReceiverId",
                principalSchema: "stakeholders",
                principalTable: "People",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ChatMessages_People_SenderId",
                schema: "stakeholders",
                table: "ChatMessages",
                column: "SenderId",
                principalSchema: "stakeholders",
                principalTable: "People",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ClubChallengeRequests_Clubs_ChallengedId",
                schema: "stakeholders",
                table: "ClubChallengeRequests",
                column: "ChallengedId",
                principalSchema: "stakeholders",
                principalTable: "Clubs",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ClubChallengeRequests_Clubs_ChallengerId",
                schema: "stakeholders",
                table: "ClubChallengeRequests",
                column: "ChallengerId",
                principalSchema: "stakeholders",
                principalTable: "Clubs",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ClubFights_Clubs_Club1Id",
                schema: "stakeholders",
                table: "ClubFights",
                column: "Club1Id",
                principalSchema: "stakeholders",
                principalTable: "Clubs",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_ClubFights_Clubs_Club2Id",
                schema: "stakeholders",
                table: "ClubFights",
                column: "Club2Id",
                principalSchema: "stakeholders",
                principalTable: "Clubs",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_Clubs_People_OwnerId",
                schema: "stakeholders",
                table: "Clubs",
                column: "OwnerId",
                principalSchema: "stakeholders",
                principalTable: "People",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_People_Clubs_ClubId",
                schema: "stakeholders",
                table: "People");

            migrationBuilder.DropTable(
                name: "AchievementClub",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ApplicationRatings",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ChatMessages",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ClubChallengeRequests",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ClubFights",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ClubInvitations",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "ClubJoinRequests",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "NewsletterPreferences",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "Notifications",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "PersonPerson",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "TourIssueComments",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "Achievements",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "TourIssue",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "Clubs",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "People",
                schema: "stakeholders");

            migrationBuilder.DropTable(
                name: "Users",
                schema: "stakeholders");
        }
    }
}
