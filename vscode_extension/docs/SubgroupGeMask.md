# SubgroupGeMask(3) Manual Page

## Name

SubgroupGeMask - Mask of shader invocations in a subgroup with the same
or higher subgroup local invocation ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupGeMask`  
Decorating a variable with the `SubgroupGeMask` builtin decoration will
make that variable contain the *subgroup mask* of the current subgroup
invocation. The bits corresponding to the invocations greater than or
equal to `SubgroupLocalInvocationId` through `SubgroupSize`-1 are set in
the variable decorated with `SubgroupGeMask`. All other bits are set to
zero.

`SubgroupGeMaskKHR` is an alias of `SubgroupGeMask`.

Valid Usage

- <a href="#VUID-SubgroupGeMask-SubgroupGeMask-04372"
  id="VUID-SubgroupGeMask-SubgroupGeMask-04372"></a>
  VUID-SubgroupGeMask-SubgroupGeMask-04372  
  The variable decorated with `SubgroupGeMask` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupGeMask-SubgroupGeMask-04373"
  id="VUID-SubgroupGeMask-SubgroupGeMask-04373"></a>
  VUID-SubgroupGeMask-SubgroupGeMask-04373  
  The variable decorated with `SubgroupGeMask` **must** be declared as a
  four-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupGeMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
