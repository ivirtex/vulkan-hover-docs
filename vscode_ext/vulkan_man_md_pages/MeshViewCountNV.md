# MeshViewCountNV(3) Manual Page

## Name

MeshViewCountNV - Number of views processed by a mesh or task shader



## <a href="#_description" class="anchor"></a>Description

`MeshViewCountNV`  
Decorating a variable with the `MeshViewCountNV` built-in decoration
will make that variable contain the number of views processed by the
current mesh or task shader invocations.

Valid Usage

- <a href="#VUID-MeshViewCountNV-MeshViewCountNV-04287"
  id="VUID-MeshViewCountNV-MeshViewCountNV-04287"></a>
  VUID-MeshViewCountNV-MeshViewCountNV-04287  
  The `MeshViewCountNV` decoration **must** be used only within the
  `MeshNV` or `TaskNV` `Execution` `Model`

- <a href="#VUID-MeshViewCountNV-MeshViewCountNV-04288"
  id="VUID-MeshViewCountNV-MeshViewCountNV-04288"></a>
  VUID-MeshViewCountNV-MeshViewCountNV-04288  
  The variable decorated with `MeshViewCountNV` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-MeshViewCountNV-MeshViewCountNV-04289"
  id="VUID-MeshViewCountNV-MeshViewCountNV-04289"></a>
  VUID-MeshViewCountNV-MeshViewCountNV-04289  
  The variable decorated with `MeshViewCountNV` **must** be declared as
  a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#MeshViewCountNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
