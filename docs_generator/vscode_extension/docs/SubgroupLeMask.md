# SubgroupLeMask(3) Manual Page

## Name

SubgroupLeMask - Mask of shader invocations in a subgroup with the same or lower subgroup local invocation ID



## [](#_description)Description

`SubgroupLeMask`

Decorating a variable with the `SubgroupLeMask` builtin decoration will make that variable contain the *subgroup mask* of the current subgroup invocation. The bits corresponding to the invocations less than or equal to `SubgroupLocalInvocationId` are set in the variable decorated with `SubgroupLeMask`. All other bits are set to zero.

`SubgroupLeMaskKHR` is an alias of `SubgroupLeMask`.

Valid Usage

- [](#VUID-SubgroupLeMask-SubgroupLeMask-04376)VUID-SubgroupLeMask-SubgroupLeMask-04376  
  The variable decorated with `SubgroupLeMask` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupLeMask-SubgroupLeMask-04377)VUID-SubgroupLeMask-SubgroupLeMask-04377  
  The variable decorated with `SubgroupLeMask` **must** be declared as a four-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupLeMask)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0