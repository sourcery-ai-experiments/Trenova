/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import React from "react";
import { faTruck } from "@fortawesome/pro-duotone-svg-icons";
import {
  LinksGroup,
  LinksGroupProps,
} from "@/components/layout/Navbar/_partials/LinksGroup";

/** Links for Dispatch Navigation Menu */
export const dispatchNavLinks = [
  {
    label: "Dispatch",
    icon: faTruck,
    link: "/",
    permission: "view_dispatch",
    links: [
      {
        label: "Rate Management",
        link: "/dispatch/rate-management/",
      },
      {
        label: "Configuration Files",
        link: "#",
        subLinks: [
          {
            label: "Comment Type",
            link: "/dispatch/comment-types/",
            permission: "view_commenttype",
          },
          {
            label: "Delay Codes",
            link: "/dispatch/delay-codes/",
            permission: "view_delaycode",
          },
          {
            label: "Fleet Codes",
            link: "/dispatch/fleet-codes/",
            permission: "view_fleetcode",
          },
          {
            label: "Locations",
            link: "/dispatch/locations/",
            permission: "view_location",
          },
          {
            label: "Routes",
            link: "/dispatch/routes/",
            permission: "view_route",
          },
          {
            label: "Route Control",
            link: "/admin/control-files#route-controls",
            permission: "view_routecontrol",
          },
          {
            label: "Locations Categories",
            link: "/dispatch/locations-categories/",
            permission: "view_locationcategory",
          },
        ],
      },
    ],
  },
] as LinksGroupProps[];

export function DispatchLinks() {
  const dispatchLinks = dispatchNavLinks.map((item) => (
    <LinksGroup {...item} key={item.label} />
  ));

  return <>{dispatchLinks}</>;
}
