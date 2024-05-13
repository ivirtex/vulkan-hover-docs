# SubgroupEqMask(3) Manual Page

## Name

SubgroupEqMask - Mask of shader invocations in a subgroup with the same
subgroup local invocation ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupEqMask`  
Decorating a variable with the `SubgroupEqMask` builtin decoration will
make that variable contain the *subgroup mask* of the current subgroup
invocation. The bit corresponding to the `SubgroupLocalInvocationId` is
set in the variable decorated with `SubgroupEqMask`. All other bits are
set to zero.

`SubgroupEqMaskKHR` is an alias of `SubgroupEqMask`.

Valid Usage

- <a href="#VUID-SubgroupEqMask-SubgroupEqMask-04370"
  id="VUID-SubgroupEqMask-SubgroupEqMask-04370"></a>
  VUID-SubgroupEqMask-SubgroupEqMask-04370  
  The variable decorated with `SubgroupEqMask` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupEqMask-SubgroupEqMask-04371"
  id="VUID-SubgroupEqMask-SubgroupEqMask-04371"></a>
  VUID-SubgroupEqMask-SubgroupEqMask-04371  
  The variable decorated with `SubgroupEqMask` **must** be declared as a
  four-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupEqMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
