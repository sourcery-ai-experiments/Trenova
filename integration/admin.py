# -*- coding: utf-8 -*-
"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from django.contrib import admin

from .models import Integration


@admin.register(Integration)
class IntegrationAdmin(admin.ModelAdmin):
    """
    Integration Admin
    """

    list_display: tuple[str, ...] = ("name", "login_url", "is_active")
    list_filter: tuple[str, ...] = ("is_active",)
    search_fields: tuple[str, ...] = ("name", "login_url", "description")
    ordering: tuple[str] = ("name",)
