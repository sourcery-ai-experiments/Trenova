# Generated by Django 4.1.3 on 2022-11-22 05:38

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0007_alter_dispatchcontrol_distance_method_and_more"),
        ("order", "0038_ordercomment_entered_by"),
    ]

    operations = [
        migrations.AlterField(
            model_name="serviceincident",
            name="delay_code",
            field=models.ForeignKey(
                blank=True,
                null=True,
                on_delete=django.db.models.deletion.PROTECT,
                related_name="service_incidents",
                related_query_name="service_incident",
                to="dispatch.delaycode",
                verbose_name="Delay Code",
            ),
        ),
    ]
