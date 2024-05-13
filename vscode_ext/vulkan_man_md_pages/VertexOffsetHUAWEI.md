# VertexOffsetHUAWEI(3) Manual Page

## Name

VertexOffsetHUAWEI - cluster culling shader output variable



## <a href="#_description" class="anchor"></a>Description

`VertexOffsetHUAWEI`  
The `VertexOffsetHUAWEI` decoration can be used to decorate a cluster
culling shader output variable,this indexed mode specific variable will
contain an integer value that specifies a offset value added to the
vertex index of a cluster before indexing into the vertex buffer.

Valid Usage

- <a href="#VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07811"
  id="VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07811"></a>
  VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07811  
  The `VertexOffsetHUAWEI` decoration **must** be used only within the
  `ClusterCullingHUAWEI` `Execution` `Model`

- <a href="#VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07812"
  id="VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07812"></a>
  VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07812  
  The variable decorated with `VertexOffsetHUAWEI` **must** be declared
  as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VertexOffsetHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
