# -*- coding: utf-8 -*-
# Generated by Django 4.1.2 on 2022-11-08 05:12

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("billing", "0009_alter_chargetype_description_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="customerfuelprofile",
            name="end_date",
            field=models.DateField(
                blank=True, help_text="End Date", null=True, verbose_name="End Date"
            ),
        ),
    ]
