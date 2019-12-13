import React from "react";
import FormLabel from "@material-ui/core/FormLabel";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Switch from "@material-ui/core/Switch";
import shallowEqual from "../utils/shallowEqual";
import FormHelperText from "@material-ui/core/FormHelperText";

const isChecked = (currentValues, value) => {
  if (Array.isArray(currentValues)) {
    return currentValues.indexOf(value) < 0 ? false : true;
  }
  return false;
};

export const MySwitch = React.memo(
  props => {
    const { label, options, handleBlur, handleChange, mutate } = props;
    const { error, touched, value, name, type } = mutate;
    let switches;
    if (Array.isArray(options)) {
      switches = options.map((currentCheckBox, index) => (
        <FormControlLabel
          key={`${index}-${currentCheckBox.value}`}
          control={
            <Switch
              type={type}
              name={name}
              value={currentCheckBox.value || ""}
              checked={isChecked(value, currentCheckBox.value)}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          }
          label={currentCheckBox.label}
        />
      ));
    }
    const refCount = React.useRef(0);
    return (
      <>
        <FormLabel error={!!error} component="legend">
          {label}
        </FormLabel>
        <FormGroup row={true}>{switches}</FormGroup>
        <FormHelperText error={touched && !!error}>{error}</FormHelperText>
        <div>RenderCount:{refCount.current++}</div>
      </>
    );
  },
  (prevProps, nextProps) => {
    if (
      !shallowEqual(prevProps.mutate, nextProps.mutate) ||
      prevProps.label !== nextProps.label
    ) {
      return false;
    } else {
      return true;
    }
  }
);
