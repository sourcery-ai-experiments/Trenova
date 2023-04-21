# Generated by Django 4.2 on 2023-04-21 03:52

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('reports', '0011_alter_customreport_table'),
    ]

    operations = [
        migrations.AlterField(
            model_name='customreport',
            name='table',
            field=models.CharField(choices=[('accessorial_charge', 'accessorial_charge'), ('additional_charge', 'additional_charge'), ('auditlog_logentry', 'auditlog_logentry'), ('auth_group_permissions', 'auth_group_permissions'), ('auth_token', 'auth_token'), ('billing_control', 'billing_control'), ('billing_exception', 'billing_exception'), ('billing_history', 'billing_history'), ('billing_queue', 'billing_queue'), ('billing_transfer_log', 'billing_transfer_log'), ('charge_type', 'charge_type'), ('comment_type', 'comment_type'), ('commodity', 'commodity'), ('custom_report', 'custom_report'), ('customer', 'customer'), ('customer_billing_profile', 'customer_billing_profile'), ('customer_contact', 'customer_contact'), ('customer_email_profile', 'customer_email_profile'), ('customer_fuel_profile', 'customer_fuel_profile'), ('customer_fuel_profile_detail', 'customer_fuel_profile_detail'), ('customer_fuel_table', 'customer_fuel_table'), ('customer_rule_profile', 'customer_rule_profile'), ('customer_rule_profile_document_class', 'customer_rule_profile_document_class'), ('delay_code', 'delay_code'), ('department', 'department'), ('depot', 'depot'), ('depot_detail', 'depot_detail'), ('dispatch_control', 'dispatch_control'), ('division_code', 'division_code'), ('django_celery_beat_clockedschedule', 'django_celery_beat_clockedschedule'), ('django_celery_beat_intervalschedule', 'django_celery_beat_intervalschedule'), ('django_celery_beat_periodictasks', 'django_celery_beat_periodictasks'), ('django_celery_results_chordcounter', 'django_celery_results_chordcounter'), ('django_celery_results_taskresult', 'django_celery_results_taskresult'), ('django_migrations', 'django_migrations'), ('document_classification', 'document_classification'), ('email_control', 'email_control'), ('email_log', 'email_log'), ('email_profile', 'email_profile'), ('equipment_maintenance_plan', 'equipment_maintenance_plan'), ('equipment_maintenance_plan_equipment_types', 'equipment_maintenance_plan_equipment_types'), ('equipment_manufacturer', 'equipment_manufacturer'), ('equipment_type', 'equipment_type'), ('equipment_type_detail', 'equipment_type_detail'), ('fleet_code', 'fleet_code'), ('fuel_vendor', 'fuel_vendor'), ('fuel_vendor_fuel_detail', 'fuel_vendor_fuel_detail'), ('general_ledger_account', 'general_ledger_account'), ('google_api', 'google_api'), ('hazardous_material', 'hazardous_material'), ('integration', 'integration'), ('integration_vendor', 'integration_vendor'), ('invoice_control', 'invoice_control'), ('job_title', 'job_title'), ('kube_configuration', 'kube_configuration'), ('location', 'location'), ('location_category', 'location_category'), ('location_comment', 'location_comment'), ('location_contact', 'location_contact'), ('movement', 'movement'), ('order', 'order'), ('order_comment', 'order_comment'), ('order_control', 'order_control'), ('order_documentation', 'order_documentation'), ('order_type', 'order_type'), ('organization', 'organization'), ('qualifier_code', 'qualifier_code'), ('rate', 'rate'), ('rate_billing_table', 'rate_billing_table'), ('reason_code', 'reason_code'), ('report_column', 'report_column'), ('revenue_code', 'revenue_code'), ('route', 'route'), ('route_control', 'route_control'), ('scheduled_report', 'scheduled_report'), ('scheduled_report_day_of_week', 'scheduled_report_day_of_week'), ('service_incident', 'service_incident'), ('silk_profile_queries', 'silk_profile_queries'), ('silk_response', 'silk_response'), ('stop', 'stop'), ('stop_comment', 'stop_comment'), ('table_change_alert', 'table_change_alert'), ('tax_rate', 'tax_rate'), ('tractor', 'tractor'), ('trailer', 'trailer'), ('user', 'user'), ('user_groups', 'user_groups'), ('user_profile', 'user_profile'), ('user_user_permissions', 'user_user_permissions'), ('weekday', 'weekday'), ('worker', 'worker'), ('worker_comment', 'worker_comment'), ('worker_contact', 'worker_contact'), ('worker_hos', 'worker_hos'), ('worker_profile', 'worker_profile'), ('worker_time_away', 'worker_time_away')], help_text='The table that the table change alert is for.', max_length=255, verbose_name='Table'),
        ),
    ]
