# FragInvocationCountEXT(3) Manual Page

## Name

FragInvocationCountEXT - Number of fragment shader invocations for a fragment



## [](#_description)Description

`FragInvocationCountEXT`

Decorating a variable with the `FragInvocationCountEXT` built-in decoration will make that variable contain the maximum number of fragment shader invocations for the fragment, as determined by `minSampleShading`.

If [Sample Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading) is not enabled, `FragInvocationCountEXT` will be filled with a value of 1.

Valid Usage

- [](#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04217)VUID-FragInvocationCountEXT-FragInvocationCountEXT-04217  
  The `FragInvocationCountEXT` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04218)VUID-FragInvocationCountEXT-FragInvocationCountEXT-04218  
  The variable decorated with `FragInvocationCountEXT` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FragInvocationCountEXT-FragInvocationCountEXT-04219)VUID-FragInvocationCountEXT-FragInvocationCountEXT-04219  
  The variable decorated with `FragInvocationCountEXT` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragInvocationCountEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0