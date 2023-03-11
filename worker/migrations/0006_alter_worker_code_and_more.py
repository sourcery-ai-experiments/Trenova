# Generated by Django 4.1.7 on 2023-03-11 09:23

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("worker", "0005_alter_workerprofile_license_number"),
    ]

    operations = [
        migrations.AlterField(
            model_name="worker",
            name="code",
            field=models.CharField(
                editable=False,
                help_text="The code of the worker. This field is required and must be unique.",
                max_length=10,
                verbose_name="Code",
            ),
        ),
        migrations.AddConstraint(
            model_name="worker",
            constraint=models.UniqueConstraint(
                fields=("code", "organization"), name="unique_worker_code_organization"
            ),
        ),
    ]
