# Generated by Django 4.1.2 on 2022-11-10 05:49

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import localflavor.us.models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("customer", "0003_delete_customer"),
    ]

    operations = [
        migrations.CreateModel(
            name="Customer",
            fields=[
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
                        default=True,
                        help_text="Designates whether this customer should be treated as active. Unselect this instead of deleting customers.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "code",
                    models.CharField(
                        editable=False,
                        help_text="Customer code",
                        max_length=10,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="Customer name", max_length=150, verbose_name="Name"
                    ),
                ),
                (
                    "address_line_1",
                    models.CharField(
                        help_text="Address line 1",
                        max_length=150,
                        verbose_name="Address Line 1",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="Address line 2",
                        max_length=150,
                        verbose_name="Address Line 2",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="City", max_length=150, verbose_name="City"
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        help_text="State", max_length=2, verbose_name="State"
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        help_text="Zip code", max_length=10, verbose_name="Zip Code"
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Customer",
                "verbose_name_plural": "Customers",
                "ordering": ["code"],
            },
        ),
    ]
