# Generated by Django 4.1.3 on 2022-11-25 17:51

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ("accounts", "0015_user_organization"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="user",
            name="organization",
        ),
    ]
