# TaskCountNV(3) Manual Page

## Name

TaskCountNV - Number of mesh shader workgroups that will be generated



## <a href="#_description" class="anchor"></a>Description

`TaskCountNV`  
Decorating a variable with the `TaskCountNV` decoration will make that
variable contain the task count. The task count specifies the number of
subsequent mesh shader workgroups that get generated upon completion of
the task shader.

Valid Usage

- <a href="#VUID-TaskCountNV-TaskCountNV-04384"
  id="VUID-TaskCountNV-TaskCountNV-04384"></a>
  VUID-TaskCountNV-TaskCountNV-04384  
  The `TaskCountNV` decoration **must** be used only within the `TaskNV`
  `Execution` `Model`

- <a href="#VUID-TaskCountNV-TaskCountNV-04385"
  id="VUID-TaskCountNV-TaskCountNV-04385"></a>
  VUID-TaskCountNV-TaskCountNV-04385  
  The variable decorated with `TaskCountNV` **must** be declared using
  the `Output` `Storage` `Class`

- <a href="#VUID-TaskCountNV-TaskCountNV-04386"
  id="VUID-TaskCountNV-TaskCountNV-04386"></a>
  VUID-TaskCountNV-TaskCountNV-04386  
  The variable decorated with `TaskCountNV` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#TaskCountNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
