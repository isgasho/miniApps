import React from "react";
import { useFormik, getIn, FormikProvider, FieldArray } from "formik";
import {
  MyTextField,
  MyCheckbox,
  MySelectDependent,
  MySelectStatic,
  MyRadio,
  MyKeyboardDatePicker,
  MySlider,
  MySwitch
} from "./components/index";
import DateFnsUtils from "@date-io/date-fns";
import { MuiPickersUtilsProvider } from "@material-ui/pickers";
import * as Yup from "yup";
import invariant from "tiny-warning";

const FieldBag = (formikBag, type, name, watch) => {
  invariant(typeof formikBag === "object", "Invalid object passed");
  invariant(!!name, "name is required");
  invariant(!!type, "type is required");
  const fieldBag = {
    mutate: {
      value: getIn(formikBag.values, name),
      error: getIn(formikBag.errors, name),
      touched: getIn(formikBag.touched, name),
      name: name,
      type: type
    },
    handleBlur: formikBag.handleBlur,
    handleChange: formikBag.handleChange
  };
  if (!!watch) {
    fieldBag.mutate.watch = getIn(formikBag.values, watch);
  }
  return fieldBag;
};

export const Form = () => {
  const formikBag = useFormik({
    initialValues: {},
    validationSchema: Yup.object().shape({
      name: Yup.string()
        .required("required")
        .min(2, "Too Short!")
        .max(50, "Too Long!"),
      age: Yup.number("Should be a number type")
        .required("required")
        .min(18, "Too Short!")
        .max(5000, "Too Long!")
    })
  });
  const [asyncError, setAsyncError] = React.useState({});

  return (
    <>
      <MuiPickersUtilsProvider utils={DateFnsUtils}>
        <FormikProvider value={formikBag}>
          <MyTextField
            variant="outlined"
            label="Name"
            {...FieldBag(formikBag, "text", "name")}
          />
          <MyTextField label="Age" {...FieldBag(formikBag, "text", "age")} />
          <MyCheckbox
            options={[
              {
                value: "apple",
                label: "Apple"
              },
              {
                value: "banana",
                label: "Banana"
              }
            ]}
            label="Fruits"
            {...FieldBag(formikBag, "checkbox", "fruits")}
          />
          <MySwitch
            options={[
              {
                value: "green",
                label: "Green"
              },
              {
                value: "blue",
                label: "Blue"
              }
            ]}
            label="Lights"
            {...FieldBag(formikBag, "checkbox", "lights")}
          />
          <MyKeyboardDatePicker
            label="Date of Birth"
            {...FieldBag(formikBag, "datetime", "dob")}
          />
          <MySlider
            label="Marks"
            {...FieldBag(formikBag, "slider", "marks")}
            min={0}
            max={100}
            step={10}
          />
          <MyRadio
            options={[
              {
                value: "male",
                label: "Male"
              },
              {
                value: "female",
                label: "Female"
              }
            ]}
            label="Gender"
            {...FieldBag(formikBag, "radio", "gender")}
          />
          <MySelectStatic
            options={[
              {
                value: "devarsh",
                label: "devarsh"
              },
              {
                value: "harsh",
                label: "harsh"
              },
              {
                value: "dvija",
                label: "dvija"
              },
              {
                value: "nirali",
                label: "nirali"
              }
            ]}
            label="Cousin"
            {...FieldBag(formikBag, "select", "cousin")}
          />
          <FieldArray
            name="address"
            render={({ push, remove, insert }) => (
              <div>
                {formikBag.values.address &&
                formikBag.values.address.length > 0 ? (
                  formikBag.values.address.map((_, index) => (
                    <div key={index}>
                      <MyTextField
                        label="Street1"
                        {...FieldBag(
                          formikBag,
                          "text",
                          `address[${index}].street1`
                        )}
                      />
                      <MyTextField
                        label="Street2"
                        {...FieldBag(
                          formikBag,
                          "text",
                          `address[${index}].street2`
                        )}
                      />
                      <button type="button" onClick={() => remove(index)}>
                        -
                      </button>
                      <button
                        type="button"
                        onClick={() => push({ street1: "", street2: "" })}
                      >
                        +
                      </button>
                    </div>
                  ))
                ) : (
                  <button onClick={() => push({ street1: "", street2: "" })}>
                    Add Address
                  </button>
                )}
              </div>
            )}
          />
          <pre>{JSON.stringify(formikBag, null, 2)}</pre>
        </FormikProvider>
      </MuiPickersUtilsProvider>
    </>
  );
};
