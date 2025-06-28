# SubgroupEqMask(3) Manual Page

## Name

SubgroupEqMask - Mask of shader invocations in a subgroup with the same subgroup local invocation ID



## [](#_description)Description

`SubgroupEqMask`

Decorating a variable with the `SubgroupEqMask` builtin decoration will make that variable contain the *subgroup mask* of the current subgroup invocation. The bit corresponding to the `SubgroupLocalInvocationId` is set in the variable decorated with `SubgroupEqMask`. All other bits are set to zero.

`SubgroupEqMaskKHR` is an alias of `SubgroupEqMask`.

Valid Usage

- [](#VUID-SubgroupEqMask-SubgroupEqMask-04370)VUID-SubgroupEqMask-SubgroupEqMask-04370  
  The variable decorated with `SubgroupEqMask` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupEqMask-SubgroupEqMask-04371)VUID-SubgroupEqMask-SubgroupEqMask-04371  
  The variable decorated with `SubgroupEqMask` **must** be declared as a four-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupEqMask)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0