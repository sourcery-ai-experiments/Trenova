# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-10-30 01:18

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="Integration",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=False,
                        help_text="Is the integration active?",
                        verbose_name="Is Active",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        choices=[
                            ("google_maps", "Google Maps"),
                            ("google_places", "Google Places"),
                        ],
                        help_text="Name of the integration",
                        max_length=100,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
                (
                    "auth_type",
                    models.CharField(
                        choices=[
                            ("no_auth", "No Auth"),
                            ("api_key", "API Key"),
                            ("oauth", "OAuth"),
                            ("oauth2", "OAuth 2.0"),
                            ("bearer_token", "Bearer Token"),
                            ("basic_auth", "Basic Auth"),
                            ("digest_auth", "Digest Auth"),
                            ("hawk_auth", "Hawk Auth"),
                            ("aws_sig4", "AWS Sig4"),
                            ("ntlm_auth", "NTLM Auth"),
                            ("akamai_edgegrid", "Akamai EdgeGrid"),
                        ],
                        default="no_auth",
                        help_text="Authentication type for the integration",
                        max_length=100,
                        verbose_name="Auth Type",
                    ),
                ),
                (
                    "login_url",
                    models.URLField(
                        blank=True,
                        help_text="Login URL for the integration",
                        max_length=255,
                        null=True,
                        verbose_name="Login URL",
                    ),
                ),
                (
                    "auth_token",
                    models.CharField(
                        blank=True,
                        help_text="API Key for the specified integration",
                        max_length=255,
                        null=True,
                        verbose_name="API Key",
                    ),
                ),
                (
                    "username",
                    models.CharField(
                        blank=True,
                        help_text="Username for the specified integration",
                        max_length=255,
                        null=True,
                        verbose_name="Username",
                    ),
                ),
                (
                    "password",
                    models.CharField(
                        blank=True,
                        help_text="Password for the specified integration",
                        max_length=255,
                        null=True,
                        verbose_name="Password",
                    ),
                ),
                (
                    "client_id",
                    models.CharField(
                        blank=True,
                        max_length=255,
                        null=True,
                        verbose_name="Client ID for the specified integration",
                    ),
                ),
                (
                    "client_secret",
                    models.CharField(
                        blank=True,
                        max_length=255,
                        null=True,
                        verbose_name="Client Secret for the specified integration",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="integrations",
                        related_query_name="integration",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Integration",
                "verbose_name_plural": "Integrations",
                "ordering": ["name"],
            },
        ),
        migrations.AddIndex(
            model_name="integration",
            index=models.Index(fields=["name"], name="integration_name_8bca70_idx"),
        ),
    ]
