# MeshViewIndicesNV(3) Manual Page

## Name

MeshViewIndicesNV - Indices of views processed by a mesh or task shader



## <a href="#_description" class="anchor"></a>Description

`MeshViewIndicesNV`  
Decorating a variable with the `MeshViewIndicesNV` built-in decoration
will make that variable contain the mesh view indices. The mesh view
indices is an array of values where each element holds the view number
of one of the views being processed by the current mesh or task shader
invocations. The values of array elements with indices greater than or
equal to `MeshViewCountNV` are undefined. If the value of
`MeshViewIndicesNV`\[i\] is j, then any outputs decorated with
`PerViewNV` will take on the value of array element i when processing
primitives for view index j.

Valid Usage

- <a href="#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04290"
  id="VUID-MeshViewIndicesNV-MeshViewIndicesNV-04290"></a>
  VUID-MeshViewIndicesNV-MeshViewIndicesNV-04290  
  The `MeshViewIndicesNV` decoration **must** be used only within the
  `MeshNV` or `TaskNV` `Execution` `Model`

- <a href="#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04291"
  id="VUID-MeshViewIndicesNV-MeshViewIndicesNV-04291"></a>
  VUID-MeshViewIndicesNV-MeshViewIndicesNV-04291  
  The variable decorated with `MeshViewIndicesNV` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04292"
  id="VUID-MeshViewIndicesNV-MeshViewIndicesNV-04292"></a>
  VUID-MeshViewIndicesNV-MeshViewIndicesNV-04292  
  The variable decorated with `MeshViewIndicesNV` **must** be declared
  as an array of scalar 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#MeshViewIndicesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
