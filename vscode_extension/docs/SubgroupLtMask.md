# SubgroupLtMask(3) Manual Page

## Name

SubgroupLtMask - Mask of shader invocations in a subgroup with a lower subgroup local invocation ID



## [](#_description)Description

`SubgroupLtMask`

Decorating a variable with the `SubgroupLtMask` builtin decoration will make that variable contain the *subgroup mask* of the current subgroup invocation. The bits corresponding to the invocations less than `SubgroupLocalInvocationId` are set in the variable decorated with `SubgroupLtMask`. All other bits are set to zero.

`SubgroupLtMaskKHR` is an alias of `SubgroupLtMask`.

Valid Usage

- [](#VUID-SubgroupLtMask-SubgroupLtMask-04378)VUID-SubgroupLtMask-SubgroupLtMask-04378  
  The variable decorated with `SubgroupLtMask` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupLtMask-SubgroupLtMask-04379)VUID-SubgroupLtMask-SubgroupLtMask-04379  
  The variable decorated with `SubgroupLtMask` **must** be declared as a four-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupLtMask).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0