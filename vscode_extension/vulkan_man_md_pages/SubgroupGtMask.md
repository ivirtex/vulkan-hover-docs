# SubgroupGtMask(3) Manual Page

## Name

SubgroupGtMask - Mask of shader invocations in a subgroup with a higher
subgroup local invocation ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupGtMask`  
Decorating a variable with the `SubgroupGtMask` builtin decoration will
make that variable contain the *subgroup mask* of the current subgroup
invocation. The bits corresponding to the invocations greater than
`SubgroupLocalInvocationId` through `SubgroupSize`-1 are set in the
variable decorated with `SubgroupGtMask`. All other bits are set to
zero.

`SubgroupGtMaskKHR` is an alias of `SubgroupGtMask`.

Valid Usage

- <a href="#VUID-SubgroupGtMask-SubgroupGtMask-04374"
  id="VUID-SubgroupGtMask-SubgroupGtMask-04374"></a>
  VUID-SubgroupGtMask-SubgroupGtMask-04374  
  The variable decorated with `SubgroupGtMask` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupGtMask-SubgroupGtMask-04375"
  id="VUID-SubgroupGtMask-SubgroupGtMask-04375"></a>
  VUID-SubgroupGtMask-SubgroupGtMask-04375  
  The variable decorated with `SubgroupGtMask` **must** be declared as a
  four-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupGtMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
