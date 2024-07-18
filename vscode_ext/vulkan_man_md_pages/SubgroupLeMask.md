# SubgroupLeMask(3) Manual Page

## Name

SubgroupLeMask - Mask of shader invocations in a subgroup with the same
or lower subgroup local invocation ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupLeMask`  
Decorating a variable with the `SubgroupLeMask` builtin decoration will
make that variable contain the *subgroup mask* of the current subgroup
invocation. The bits corresponding to the invocations less than or equal
to `SubgroupLocalInvocationId` are set in the variable decorated with
`SubgroupLeMask`. All other bits are set to zero.

`SubgroupLeMaskKHR` is an alias of `SubgroupLeMask`.

Valid Usage

- <a href="#VUID-SubgroupLeMask-SubgroupLeMask-04376"
  id="VUID-SubgroupLeMask-SubgroupLeMask-04376"></a>
  VUID-SubgroupLeMask-SubgroupLeMask-04376  
  The variable decorated with `SubgroupLeMask` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupLeMask-SubgroupLeMask-04377"
  id="VUID-SubgroupLeMask-SubgroupLeMask-04377"></a>
  VUID-SubgroupLeMask-SubgroupLeMask-04377  
  The variable decorated with `SubgroupLeMask` **must** be declared as a
  four-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupLeMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
