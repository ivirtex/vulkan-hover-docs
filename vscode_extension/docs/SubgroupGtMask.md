# SubgroupGtMask(3) Manual Page

## Name

SubgroupGtMask - Mask of shader invocations in a subgroup with a higher subgroup local invocation ID



## [](#_description)Description

`SubgroupGtMask`

Decorating a variable with the `SubgroupGtMask` builtin decoration will make that variable contain the *subgroup mask* of the current subgroup invocation. The bits corresponding to the invocations greater than `SubgroupLocalInvocationId` through `SubgroupSize`-1 are set in the variable decorated with `SubgroupGtMask`. All other bits are set to zero.

`SubgroupGtMaskKHR` is an alias of `SubgroupGtMask`.

Valid Usage

- [](#VUID-SubgroupGtMask-SubgroupGtMask-04374)VUID-SubgroupGtMask-SubgroupGtMask-04374  
  The variable decorated with `SubgroupGtMask` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupGtMask-SubgroupGtMask-04375)VUID-SubgroupGtMask-SubgroupGtMask-04375  
  The variable decorated with `SubgroupGtMask` **must** be declared as a four-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupGtMask).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0