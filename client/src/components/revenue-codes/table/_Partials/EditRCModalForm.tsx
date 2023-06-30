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

import { RevenueCode } from "@/types/apps/accounting";
import React from "react";
import {
  Box,
  Button,
  createStyles,
  Group,
  rem,
  SimpleGrid,
} from "@mantine/core";
import { SelectInput, SelectItem } from "@/components/ui/fields/SelectInput";
import { ValidatedTextInput } from "@/components/ui/fields/TextInput";
import { ValidatedTextArea } from "@/components/ui/fields/TextArea";
import { useMutation, useQueryClient } from "react-query";
import axios from "@/lib/AxiosConfig";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck, faXmark } from "@fortawesome/pro-solid-svg-icons";
import { APIError } from "@/types/server";
import * as Yup from "yup";
import { useForm, yupResolver } from "@mantine/form";
import {
  divisionCodeTableStore,
  revenueCodeTableStore,
} from "@/stores/AccountingStores";

type Props = {
  revenueCode: RevenueCode;
  selectGlAccountData: SelectItem[];
};

interface EditRevenueCodeFormValues {
  code: string;
  description: string;
  expense_account: string;
  revenue_account: string;
}

const useStyles = createStyles((theme) => {
  const BREAKPOINT = theme.fn.smallerThan("sm");

  return {
    fields: {
      marginTop: rem(10),
    },
    control: {
      [BREAKPOINT]: {
        flex: 1,
      },
    },
    text: {
      color: theme.colorScheme === "dark" ? "white" : "black",
    },
    invalid: {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.fn.rgba(theme.colors.red[8], 0.15)
          : theme.colors.red[0],
    },
    invalidIcon: {
      color: theme.colors.red[theme.colorScheme === "dark" ? 7 : 6],
    },
    div: {
      marginBottom: rem(10),
    },
  };
});

export const EditRCModalForm: React.FC<Props> = ({
  revenueCode,
  selectGlAccountData,
}) => {
  const { classes } = useStyles();
  const [loading, setLoading] = React.useState<boolean>(false);
  const queryClient = useQueryClient();

  const mutation = useMutation(
    (values: EditRevenueCodeFormValues) =>
      axios.put(`/revenue_codes/${revenueCode.id}/`, values),
    {
      onSuccess: () => {
        queryClient
          .invalidateQueries({
            queryKey: ["revenue-code-table-data"],
          })
          .then(() => {
            notifications.show({
              title: "Success",
              message: "Revenue Code updated successfully",
              color: "green",
              withCloseButton: true,
              icon: <FontAwesomeIcon icon={faCheck} />,
            });
            revenueCodeTableStore.set("editModalOpen", false);
          });
      },
      onError: (error: any) => {
        const { data } = error.response;
        if (data.type === "validation_error") {
          data.errors.forEach((error: APIError) => {
            form.setFieldError(error.attr, error.detail);
            if (error.attr === "non_field_errors") {
              notifications.show({
                title: "Error",
                message: error.detail,
                color: "red",
                withCloseButton: true,
                icon: <FontAwesomeIcon icon={faXmark} />,
                autoClose: 10_000, // 10 seconds
              });
            }
          });
        }
      },
      onSettled: () => {
        setLoading(false);
      },
    }
  );

  const editRevenueCodeSchema = Yup.object().shape({
    code: Yup.string()
      .max(4, "Code cannot be longer than 4 characters")
      .required("Code is required"),
    description: Yup.string()
      .max(100, "Description cannot be longer than 100 characters")
      .required("Description is required"),
    expense_account: Yup.string().notRequired(),
    revenue_account: Yup.string().notRequired(),
  });

  const form = useForm<EditRevenueCodeFormValues>({
    validate: yupResolver(editRevenueCodeSchema),
    initialValues: {
      code: revenueCode.code,
      description: revenueCode.description,
      expense_account: revenueCode.expense_account,
      revenue_account: revenueCode.expense_account,
    },
  });

  const submitForm = (values: EditRevenueCodeFormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <>
      <form onSubmit={form.onSubmit((values) => submitForm(values))}>
        <Box className={classes.div}>
          <Box>
            <ValidatedTextInput
              form={form}
              className={classes.fields}
              name="code"
              label="Code"
              placeholder="Code"
              variant="filled"
              withAsterisk
            />
            <ValidatedTextArea
              form={form}
              className={classes.fields}
              name="description"
              label="Description"
              placeholder="Description"
              variant="filled"
              withAsterisk
            />
            <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
              <SelectInput
                form={form}
                data={selectGlAccountData}
                className={classes.fields}
                name="expense_account"
                label="Expense Account"
                placeholder="Expense Account"
                variant="filled"
                onMouseLeave={() => {
                  form.setFieldValue(
                    "expense_account",
                    form.values.expense_account
                  );
                }}
                clearable
              />
              <SelectInput
                form={form}
                data={selectGlAccountData}
                className={classes.fields}
                name="revenue_account"
                label="Revenue Account"
                placeholder="Revenue Account"
                variant="filled"
                onMouseLeave={() => {
                  form.setFieldValue(
                    "revenue_account",
                    form.values.revenue_account
                  );
                }}
                clearable
              />
            </SimpleGrid>
            <Group position="right" mt="md">
              <Button
                color="white"
                type="submit"
                className={classes.control}
                loading={loading}
              >
                Submit
              </Button>
            </Group>
          </Box>
        </Box>
      </form>
    </>
  );
};
