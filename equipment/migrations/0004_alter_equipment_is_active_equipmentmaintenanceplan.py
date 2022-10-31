# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-10-31 05:36

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0004_alter_organization_language_and_more"),
        ("equipment", "0003_alter_equipment_manufacturer"),
    ]

    operations = [
        migrations.AlterField(
            model_name="equipment",
            name="is_active",
            field=models.BooleanField(
                default=True,
                help_text="Whether the Equipment is active or not.",
                verbose_name="Active",
            ),
        ),
        migrations.CreateModel(
            name="EquipmentMaintenancePlan",
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
                    "id",
                    models.CharField(
                        help_text="ID of the equipment maintenance plan.",
                        max_length=50,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "description",
                    models.TextField(
                        blank=True,
                        help_text="Description of the equipment maintenance plan.",
                        verbose_name="Description",
                    ),
                ),
                (
                    "by_distance",
                    models.BooleanField(
                        default=False,
                        help_text="Maintenance plan is by distance.",
                        verbose_name="By Distance",
                    ),
                ),
                (
                    "by_time",
                    models.BooleanField(
                        default=False,
                        help_text="Maintenance plan is by time.",
                        verbose_name="By Time",
                    ),
                ),
                (
                    "by_engine_hours",
                    models.BooleanField(
                        default=False,
                        help_text="Maintenance plan is by engine hours.",
                        verbose_name="By Engine Hours",
                    ),
                ),
                (
                    "miles",
                    models.PositiveIntegerField(
                        default=0,
                        help_text="Miles of the equipment maintenance plan.",
                        verbose_name="Miles",
                    ),
                ),
                (
                    "months",
                    models.PositiveIntegerField(
                        default=0,
                        help_text="Months of the equipment maintenance plan.",
                        verbose_name="Months",
                    ),
                ),
                (
                    "engine_hours",
                    models.PositiveIntegerField(
                        default=0,
                        help_text="Engine hours of the equipment maintenance plan.",
                        verbose_name="Engine Hours",
                    ),
                ),
                (
                    "equipment_types",
                    models.ManyToManyField(
                        related_name="maintenance_plans",
                        related_query_name="maintenance_plan",
                        to="equipment.equipmenttype",
                        verbose_name="Equipment Types",
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
                "verbose_name": "Equipment Maintenance Plan",
                "verbose_name_plural": "Equipment Maintenance Plans",
                "ordering": ["id"],
            },
        ),
    ]
