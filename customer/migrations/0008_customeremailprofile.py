# Generated by Django 4.1.2 on 2022-11-10 16:36

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("customer", "0007_customerbillingprofile"),
    ]

    operations = [
        migrations.CreateModel(
            name="CustomerEmailProfile",
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
                    "name",
                    models.CharField(
                        help_text="Name",
                        max_length=50,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
                (
                    "subject",
                    models.CharField(
                        blank=True,
                        help_text="Subject",
                        max_length=100,
                        verbose_name="Subject",
                    ),
                ),
                (
                    "comment",
                    models.CharField(
                        blank=True,
                        help_text="Comment",
                        max_length=100,
                        verbose_name="Comment",
                    ),
                ),
                (
                    "from_address",
                    models.CharField(
                        blank=True,
                        help_text="From Address",
                        max_length=255,
                        verbose_name="From Address",
                    ),
                ),
                (
                    "blind_copy",
                    models.CharField(
                        blank=True,
                        help_text="Blind Copy",
                        max_length=255,
                        verbose_name="Blind Copy",
                    ),
                ),
                (
                    "read_receipt",
                    models.BooleanField(
                        default=False,
                        help_text="Read Receipt",
                        verbose_name="Read Receipt",
                    ),
                ),
                (
                    "read_receipt_to",
                    models.CharField(
                        blank=True,
                        help_text="Read Receipt To",
                        max_length=255,
                        verbose_name="Read Receipt To",
                    ),
                ),
                (
                    "attachment_name",
                    models.CharField(
                        blank=True,
                        help_text="Attachment Name",
                        max_length=255,
                        verbose_name="Attachment Name",
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
                "verbose_name": "Customer Email Profile",
                "verbose_name_plural": "Customer Email Profiles",
                "ordering": ["-name"],
            },
        ),
    ]
