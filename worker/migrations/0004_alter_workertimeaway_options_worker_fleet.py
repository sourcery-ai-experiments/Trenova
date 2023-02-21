# Generated by Django 4.1.7 on 2023-02-20 20:16

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0004_dispatchcontrol_enforce_driver_ta"),
        ("worker", "0003_workertimeaway"),
    ]

    operations = [
        migrations.AlterModelOptions(
            name="workertimeaway",
            options={
                "ordering": ["worker"],
                "verbose_name": "Worker Time Away",
                "verbose_name_plural": "Worker Time Away",
            },
        ),
        migrations.AddField(
            model_name="worker",
            name="fleet",
            field=models.ForeignKey(
                blank=True,
                help_text="The fleet of the worker.",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="worker",
                related_query_name="workers",
                to="dispatch.fleetcode",
                verbose_name="Fleet",
            ),
        ),
    ]
