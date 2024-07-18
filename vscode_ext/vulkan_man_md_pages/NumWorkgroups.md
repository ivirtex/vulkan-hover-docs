# NumWorkgroups(3) Manual Page

## Name

NumWorkgroups - Number of workgroups in a dispatch



## <a href="#_description" class="anchor"></a>Description

`NumWorkgroups`  
Decorating a variable with the `NumWorkgroups` built-in decoration will
make that variable contain the number of local workgroups that are part
of the dispatch that the invocation belongs to. Each component is equal
to the values of the workgroup count parameters passed into the
dispatching commands.

Valid Usage

- <a href="#VUID-NumWorkgroups-NumWorkgroups-04296"
  id="VUID-NumWorkgroups-NumWorkgroups-04296"></a>
  VUID-NumWorkgroups-NumWorkgroups-04296  
  The `NumWorkgroups` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, or `TaskEXT` `Execution` `Model`

- <a href="#VUID-NumWorkgroups-NumWorkgroups-04297"
  id="VUID-NumWorkgroups-NumWorkgroups-04297"></a>
  VUID-NumWorkgroups-NumWorkgroups-04297  
  The variable decorated with `NumWorkgroups` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-NumWorkgroups-NumWorkgroups-04298"
  id="VUID-NumWorkgroups-NumWorkgroups-04298"></a>
  VUID-NumWorkgroups-NumWorkgroups-04298  
  The variable decorated with `NumWorkgroups` **must** be declared as a
  three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#NumWorkgroups"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
