# SubgroupGeMask(3) Manual Page

## Name

SubgroupGeMask - Mask of shader invocations in a subgroup with the same or higher subgroup local invocation ID



## [](#_description)Description

`SubgroupGeMask`

Decorating a variable with the `SubgroupGeMask` builtin decoration will make that variable contain the *subgroup mask* of the current subgroup invocation. The bits corresponding to the invocations greater than or equal to `SubgroupLocalInvocationId` through `SubgroupSize`-1 are set in the variable decorated with `SubgroupGeMask`. All other bits are set to zero.

`SubgroupGeMaskKHR` is an alias of `SubgroupGeMask`.

Valid Usage

- [](#VUID-SubgroupGeMask-SubgroupGeMask-04372)VUID-SubgroupGeMask-SubgroupGeMask-04372  
  The variable decorated with `SubgroupGeMask` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupGeMask-SubgroupGeMask-04373)VUID-SubgroupGeMask-SubgroupGeMask-04373  
  The variable decorated with `SubgroupGeMask` **must** be declared as a four-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupGeMask).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0