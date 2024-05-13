# SubgroupId(3) Manual Page

## Name

SubgroupId - Subgroup ID



## <a href="#_description" class="anchor"></a>Description

`SubgroupId`  
Decorating a variable with the `SubgroupId` built-in decoration will
make that variable contain the index of the subgroup within the local
workgroup. This variable is in range \[0, `NumSubgroups`-1\].

Valid Usage

- <a href="#VUID-SubgroupId-SubgroupId-04367"
  id="VUID-SubgroupId-SubgroupId-04367"></a>
  VUID-SubgroupId-SubgroupId-04367  
  The `SubgroupId` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-SubgroupId-SubgroupId-04368"
  id="VUID-SubgroupId-SubgroupId-04368"></a>
  VUID-SubgroupId-SubgroupId-04368  
  The variable decorated with `SubgroupId` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-SubgroupId-SubgroupId-04369"
  id="VUID-SubgroupId-SubgroupId-04369"></a>
  VUID-SubgroupId-SubgroupId-04369  
  The variable decorated with `SubgroupId` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
