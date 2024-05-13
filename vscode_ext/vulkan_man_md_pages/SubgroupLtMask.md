# SubgroupLtMask(3) Manual Page

## Name

SubgroupLtMask - Mask of shader invocations in a subgroup with a lower
subgroup local invocation ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupLtMask`  
Decorating a variable with the `SubgroupLtMask` builtin decoration will
make that variable contain the *subgroup mask* of the current subgroup
invocation. The bits corresponding to the invocations less than
`SubgroupLocalInvocationId` are set in the variable decorated with
`SubgroupLtMask`. All other bits are set to zero.

`SubgroupLtMaskKHR` is an alias of `SubgroupLtMask`.

Valid Usage

- <a href="#VUID-SubgroupLtMask-SubgroupLtMask-04378"
  id="VUID-SubgroupLtMask-SubgroupLtMask-04378"></a>
  VUID-SubgroupLtMask-SubgroupLtMask-04378  
  The variable decorated with `SubgroupLtMask` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupLtMask-SubgroupLtMask-04379"
  id="VUID-SubgroupLtMask-SubgroupLtMask-04379"></a>
  VUID-SubgroupLtMask-SubgroupLtMask-04379  
  The variable decorated with `SubgroupLtMask` **must** be declared as a
  four-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupLtMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
