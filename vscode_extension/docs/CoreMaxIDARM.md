# CoreMaxIDARM(3) Manual Page

## Name

CoreMaxIDARM - Max core ID that can be observed on the device running the invovation reading CoreMaxIDARM



## [](#_description)Description

`CoreMaxIDARM`

Decorating a variable with the `CoreMaxIDARM` built-in decoration will make that variable contain the max ID of any shader core on the device on which the current shader invocation is running.

Valid Usage

- [](#VUID-CoreMaxIDARM-CoreMaxIDARM-07597)VUID-CoreMaxIDARM-CoreMaxIDARM-07597  
  The variable decorated with `CoreMaxIDARM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-CoreMaxIDARM-CoreMaxIDARM-07598)VUID-CoreMaxIDARM-CoreMaxIDARM-07598  
  The variable decorated with `CoreMaxIDARM` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CoreMaxIDARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0