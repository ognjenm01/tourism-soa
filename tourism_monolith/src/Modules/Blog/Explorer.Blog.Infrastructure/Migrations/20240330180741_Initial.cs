using System;
using System.Collections.Generic;
using Explorer.Blog.Core.Domain;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Explorer.Blog.Infrastructure.Migrations
{
    /// <inheritdoc />
    public partial class Initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.EnsureSchema(
                name: "blog");

            migrationBuilder.CreateTable(
                name: "BlogComments",
                schema: "blog",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    BlogId = table.Column<int>(type: "integer", nullable: false),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    Comment = table.Column<string>(type: "text", nullable: false),
                    PostTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    LastEditTime = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    IsDeleted = table.Column<bool>(type: "boolean", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_BlogComments", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "BlogRatings",
                schema: "blog",
                columns: table => new
                {
                    Rating = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                });

            migrationBuilder.CreateTable(
                name: "Blogs",
                schema: "blog",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    CreatorId = table.Column<int>(type: "integer", nullable: false),
                    Title = table.Column<string>(type: "text", nullable: false),
                    Description = table.Column<string>(type: "text", nullable: false),
                    SystemStatus = table.Column<int>(type: "integer", nullable: false),
                    CreationDate = table.Column<DateOnly>(type: "date", nullable: false),
                    ImageLinks = table.Column<List<string>>(type: "text[]", nullable: true),
                    BlogRatings = table.Column<List<BlogRating>>(type: "jsonb", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Blogs", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "BlogStatuses",
                schema: "blog",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    BlogId = table.Column<long>(type: "bigint", nullable: false),
                    Name = table.Column<string>(type: "text", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_BlogStatuses", x => x.Id);
                    table.ForeignKey(
                        name: "FK_BlogStatuses_Blogs_BlogId",
                        column: x => x.BlogId,
                        principalSchema: "blog",
                        principalTable: "Blogs",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_BlogStatuses_BlogId",
                schema: "blog",
                table: "BlogStatuses",
                column: "BlogId");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "BlogComments",
                schema: "blog");

            migrationBuilder.DropTable(
                name: "BlogRatings",
                schema: "blog");

            migrationBuilder.DropTable(
                name: "BlogStatuses",
                schema: "blog");

            migrationBuilder.DropTable(
                name: "Blogs",
                schema: "blog");
        }
    }
}
