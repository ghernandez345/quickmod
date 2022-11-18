import React from "react";

interface ITestModProps {}

const baseClass = "test-mod"

const TestMod = ({}: ITestModProps) => {
  return <div classname={baseClass}></div>;
};

export default TestMod;
