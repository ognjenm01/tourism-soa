﻿// <auto-generated />
using System;
using Explorer.Encounters.Infrastructure.Database;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Migrations;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Explorer.Encounters.Infrastructure.Migrations
{
    [DbContext(typeof(EncountersContext))]
    [Migration("20240330180819_Initial")]
    partial class Initial
    {
        /// <inheritdoc />
        protected override void BuildTargetModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder
                .HasDefaultSchema("encounters")
                .HasAnnotation("ProductVersion", "7.0.5")
                .HasAnnotation("Relational:MaxIdentifierLength", 63);

            NpgsqlModelBuilderExtensions.UseIdentityByDefaultColumns(modelBuilder);

            modelBuilder.Entity("Explorer.Encounters.Core.Domain.Encounter", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<int>("ApprovalStatus")
                        .HasColumnType("integer");

                    b.Property<string>("Description")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Image")
                        .HasColumnType("text");

                    b.Property<double?>("ImageLatitude")
                        .HasColumnType("double precision");

                    b.Property<double?>("ImageLongitude")
                        .HasColumnType("double precision");

                    b.Property<double>("Latitude")
                        .HasColumnType("double precision");

                    b.Property<double>("Longitude")
                        .HasColumnType("double precision");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int?>("PeopleCount")
                        .HasColumnType("integer");

                    b.Property<double>("Range")
                        .HasColumnType("double precision");

                    b.Property<int>("Status")
                        .HasColumnType("integer");

                    b.Property<int>("Type")
                        .HasColumnType("integer");

                    b.Property<int>("UserId")
                        .HasColumnType("integer");

                    b.Property<int>("Xp")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.ToTable("Encounters", "encounters");
                });

            modelBuilder.Entity("Explorer.Encounters.Core.Domain.EncounterCompletion", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("EncounterId")
                        .HasColumnType("bigint");

                    b.Property<DateTime>("LastUpdatedAt")
                        .HasColumnType("timestamp with time zone");

                    b.Property<int>("Status")
                        .HasColumnType("integer");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.Property<int>("Xp")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("EncounterId");

                    b.ToTable("EncounterCompletions", "encounters");
                });

            modelBuilder.Entity("Explorer.Encounters.Core.Domain.KeypointEncounter", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("EncounterId")
                        .HasColumnType("bigint");

                    b.Property<bool>("IsRequired")
                        .HasColumnType("boolean");

                    b.Property<long>("KeyPointId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("EncounterId")
                        .IsUnique();

                    b.ToTable("KeypointEncounters", "encounters");
                });

            modelBuilder.Entity("Explorer.Encounters.Core.Domain.EncounterCompletion", b =>
                {
                    b.HasOne("Explorer.Encounters.Core.Domain.Encounter", "Encounter")
                        .WithMany()
                        .HasForeignKey("EncounterId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Encounter");
                });

            modelBuilder.Entity("Explorer.Encounters.Core.Domain.KeypointEncounter", b =>
                {
                    b.HasOne("Explorer.Encounters.Core.Domain.Encounter", "Encounter")
                        .WithOne()
                        .HasForeignKey("Explorer.Encounters.Core.Domain.KeypointEncounter", "EncounterId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Encounter");
                });
#pragma warning restore 612, 618
        }
    }
}
