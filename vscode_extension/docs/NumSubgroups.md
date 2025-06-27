# NumSubgroups(3) Manual Page

## Name

NumSubgroups - Number of subgroups in a workgroup



## <a href="#_description" class="anchor"></a>Description

`NumSubgroups`  
Decorating a variable with the `NumSubgroups` built-in decoration will
make that variable contain the number of subgroups in the local
workgroup.

Valid Usage

- <a href="#VUID-NumSubgroups-NumSubgroups-04293"
  id="VUID-NumSubgroups-NumSubgroups-04293"></a>
  VUID-NumSubgroups-NumSubgroups-04293  
  The `NumSubgroups` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-NumSubgroups-NumSubgroups-04294"
  id="VUID-NumSubgroups-NumSubgroups-04294"></a>
  VUID-NumSubgroups-NumSubgroups-04294  
  The variable decorated with `NumSubgroups` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-NumSubgroups-NumSubgroups-04295"
  id="VUID-NumSubgroups-NumSubgroups-04295"></a>
  VUID-NumSubgroups-NumSubgroups-04295  
  The variable decorated with `NumSubgroups` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#NumSubgroups"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
