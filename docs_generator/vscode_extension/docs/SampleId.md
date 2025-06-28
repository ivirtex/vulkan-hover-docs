# SampleId(3) Manual Page

## Name

SampleId - Sample ID within a fragment



## [](#_description)Description

`SampleId`

Decorating a variable with the `SampleId` built-in decoration will make that variable contain the [coverage index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) for the current fragment shader invocation. `SampleId` ranges from zero to the number of samples in the framebuffer minus one. If a fragment shader entry point’s interface includes an input variable decorated with `SampleId`, [Sample Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading) is considered enabled with a `minSampleShading` value of 1.0.

Valid Usage

- [](#VUID-SampleId-SampleId-04354)VUID-SampleId-SampleId-04354  
  The `SampleId` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-SampleId-SampleId-04355)VUID-SampleId-SampleId-04355  
  The variable decorated with `SampleId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SampleId-SampleId-04356)VUID-SampleId-SampleId-04356  
  The variable decorated with `SampleId` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SampleId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0