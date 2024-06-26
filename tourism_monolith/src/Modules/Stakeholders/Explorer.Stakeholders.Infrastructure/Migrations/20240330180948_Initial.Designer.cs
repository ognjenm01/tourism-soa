﻿// <auto-generated />
using System;
using Explorer.Stakeholders.Infrastructure.Database;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Migrations;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Explorer.Stakeholders.Infrastructure.Migrations
{
    [DbContext(typeof(StakeholdersContext))]
    [Migration("20240330180948_Initial")]
    partial class Initial
    {
        /// <inheritdoc />
        protected override void BuildTargetModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder
                .HasDefaultSchema("stakeholders")
                .HasAnnotation("ProductVersion", "7.0.5")
                .HasAnnotation("Relational:MaxIdentifierLength", 63);

            NpgsqlModelBuilderExtensions.UseIdentityByDefaultColumns(modelBuilder);

            modelBuilder.Entity("AchievementClub", b =>
                {
                    b.Property<long>("AchievementsId")
                        .HasColumnType("bigint");

                    b.Property<long>("ClubId")
                        .HasColumnType("bigint");

                    b.HasKey("AchievementsId", "ClubId");

                    b.HasIndex("ClubId");

                    b.ToTable("AchievementClub", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Achievement", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Description")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int>("Type")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.ToTable("Achievements", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ApplicationRating", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Comment")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<bool>("IsRated")
                        .HasColumnType("boolean");

                    b.Property<DateTime>("LastModified")
                        .HasColumnType("timestamp with time zone");

                    b.Property<int>("Rating")
                        .HasColumnType("integer");

                    b.Property<int>("UserId")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.ToTable("ApplicationRatings", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ChatMessage", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Content")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<DateTime>("CreationDateTime")
                        .HasColumnType("timestamp with time zone");

                    b.Property<bool>("IsRead")
                        .HasColumnType("boolean");

                    b.Property<long>("ReceiverId")
                        .HasColumnType("bigint");

                    b.Property<long>("SenderId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("ReceiverId");

                    b.HasIndex("SenderId");

                    b.ToTable("ChatMessages", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Club", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Description")
                        .HasColumnType("text");

                    b.Property<int>("FightsWon")
                        .HasColumnType("integer");

                    b.Property<string>("Image")
                        .HasColumnType("text");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<long>("OwnerId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("OwnerId");

                    b.ToTable("Clubs", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubChallengeRequest", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("ChallengedId")
                        .HasColumnType("bigint");

                    b.Property<long>("ChallengerId")
                        .HasColumnType("bigint");

                    b.Property<int>("Status")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("ChallengedId");

                    b.HasIndex("ChallengerId");

                    b.ToTable("ClubChallengeRequests", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubFight", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("Club1Id")
                        .HasColumnType("bigint");

                    b.Property<long>("Club2Id")
                        .HasColumnType("bigint");

                    b.Property<DateTime>("EndOfFight")
                        .HasColumnType("timestamp with time zone");

                    b.Property<bool>("IsInProgress")
                        .HasColumnType("boolean");

                    b.Property<DateTime>("StartOfFight")
                        .HasColumnType("timestamp with time zone");

                    b.Property<int?>("WinnerId")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("Club1Id");

                    b.HasIndex("Club2Id");

                    b.ToTable("ClubFights", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubInvitation", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("ClubId")
                        .HasColumnType("bigint");

                    b.Property<int>("Status")
                        .HasColumnType("integer");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.ToTable("ClubInvitations", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubJoinRequest", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("ClubId")
                        .HasColumnType("bigint");

                    b.Property<int>("Status")
                        .HasColumnType("integer");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.ToTable("ClubJoinRequests", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.NewsletterPreference", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<long>("Frequency")
                        .HasColumnType("bigint");

                    b.Property<DateTime>("LastSent")
                        .HasColumnType("timestamp with time zone");

                    b.Property<long>("UserID")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("UserID")
                        .IsUnique();

                    b.ToTable("NewsletterPreferences", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Notification", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("ActionURL")
                        .HasColumnType("text");

                    b.Property<string>("Content")
                        .HasColumnType("text");

                    b.Property<DateTime>("CreationDateTime")
                        .HasColumnType("timestamp with time zone");

                    b.Property<bool>("IsRead")
                        .HasColumnType("boolean");

                    b.Property<int>("Type")
                        .HasColumnType("integer");

                    b.Property<int>("UserId")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.ToTable("Notifications", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Person", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Biography")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<long?>("ClubId")
                        .HasColumnType("bigint");

                    b.Property<string>("Email")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int>("Level")
                        .HasColumnType("integer");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("ProfileImage")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Quote")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Surname")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.Property<int>("XP")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("ClubId");

                    b.HasIndex("UserId")
                        .IsUnique();

                    b.ToTable("People", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.TourIssue", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Category")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<DateTime>("CreationDateTime")
                        .HasColumnType("timestamp with time zone");

                    b.Property<string>("Description")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<bool>("IsResolved")
                        .HasColumnType("boolean");

                    b.Property<long>("Priority")
                        .HasColumnType("bigint");

                    b.Property<DateTime?>("ResolveDateTime")
                        .HasColumnType("timestamp with time zone");

                    b.Property<long>("TourId")
                        .HasColumnType("bigint");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("UserId");

                    b.ToTable("TourIssue", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.TourIssueComment", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Comment")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<DateTime>("CreationDateTime")
                        .HasColumnType("timestamp with time zone");

                    b.Property<long>("TourIssueId")
                        .HasColumnType("bigint");

                    b.Property<long>("UserId")
                        .HasColumnType("bigint");

                    b.HasKey("Id");

                    b.HasIndex("TourIssueId");

                    b.HasIndex("UserId");

                    b.ToTable("TourIssueComments", "stakeholders");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.User", b =>
                {
                    b.Property<long>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("bigint");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<long>("Id"));

                    b.Property<string>("Email")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<bool>("IsActive")
                        .HasColumnType("boolean");

                    b.Property<bool>("IsBlocked")
                        .HasColumnType("boolean");

                    b.Property<bool>("IsEnabled")
                        .HasColumnType("boolean");

                    b.Property<string>("Password")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int>("Role")
                        .HasColumnType("integer");

                    b.Property<string>("Username")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("VerificationToken")
                        .IsRequired()
                        .HasColumnType("text");

                    b.HasKey("Id");

                    b.HasIndex("Username")
                        .IsUnique();

                    b.ToTable("Users", "stakeholders");
                });

            modelBuilder.Entity("PersonPerson", b =>
                {
                    b.Property<long>("FollowersId")
                        .HasColumnType("bigint");

                    b.Property<long>("FollowingId")
                        .HasColumnType("bigint");

                    b.HasKey("FollowersId", "FollowingId");

                    b.HasIndex("FollowingId");

                    b.ToTable("PersonPerson", "stakeholders");
                });

            modelBuilder.Entity("AchievementClub", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Achievement", null)
                        .WithMany()
                        .HasForeignKey("AchievementsId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", null)
                        .WithMany()
                        .HasForeignKey("ClubId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ChatMessage", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Person", "Receiver")
                        .WithMany()
                        .HasForeignKey("ReceiverId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.Person", "Sender")
                        .WithMany()
                        .HasForeignKey("SenderId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Receiver");

                    b.Navigation("Sender");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Club", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Person", "Owner")
                        .WithMany()
                        .HasForeignKey("OwnerId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Owner");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubChallengeRequest", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", "Challenged")
                        .WithMany()
                        .HasForeignKey("ChallengedId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", "Challenger")
                        .WithMany()
                        .HasForeignKey("ChallengerId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Challenged");

                    b.Navigation("Challenger");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.ClubFight", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", "Club1")
                        .WithMany()
                        .HasForeignKey("Club1Id")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", "Club2")
                        .WithMany()
                        .HasForeignKey("Club2Id")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Club1");

                    b.Navigation("Club2");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.NewsletterPreference", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.User", "User")
                        .WithOne("NewsletterPreference")
                        .HasForeignKey("Explorer.Stakeholders.Core.Domain.NewsletterPreference", "UserID")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("User");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Person", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Club", "Club")
                        .WithMany("Members")
                        .HasForeignKey("ClubId")
                        .OnDelete(DeleteBehavior.SetNull);

                    b.HasOne("Explorer.Stakeholders.Core.Domain.User", null)
                        .WithOne()
                        .HasForeignKey("Explorer.Stakeholders.Core.Domain.Person", "UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Club");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.TourIssue", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.User", null)
                        .WithMany("Issues")
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.TourIssueComment", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.TourIssue", null)
                        .WithMany("Comments")
                        .HasForeignKey("TourIssueId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.User", null)
                        .WithMany("IssueComments")
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("PersonPerson", b =>
                {
                    b.HasOne("Explorer.Stakeholders.Core.Domain.Person", null)
                        .WithMany()
                        .HasForeignKey("FollowersId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("Explorer.Stakeholders.Core.Domain.Person", null)
                        .WithMany()
                        .HasForeignKey("FollowingId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.Club", b =>
                {
                    b.Navigation("Members");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.TourIssue", b =>
                {
                    b.Navigation("Comments");
                });

            modelBuilder.Entity("Explorer.Stakeholders.Core.Domain.User", b =>
                {
                    b.Navigation("IssueComments");

                    b.Navigation("Issues");

                    b.Navigation("NewsletterPreference");
                });
#pragma warning restore 612, 618
        }
    }
}
