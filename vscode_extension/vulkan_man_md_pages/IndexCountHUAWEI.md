# IndexCountHUAWEI(3) Manual Page

## Name

IndexCountHUAWEI - cluster culling shader output variable



## <a href="#_description" class="anchor"></a>Description

`IndexCountHUAWEI`  
The `IndexCountHUAWEI` decoration can be used to decorate a cluster
culling shader output variable,this indexed mode specific variable will
contain an integer value that specifies the number of indexed vertices
in a cluster to draw.

Valid Usage

- <a href="#VUID-IndexCountHUAWEI-IndexCountHUAWEI-07805"
  id="VUID-IndexCountHUAWEI-IndexCountHUAWEI-07805"></a>
  VUID-IndexCountHUAWEI-IndexCountHUAWEI-07805  
  The `IndexCountHUAWEI` decoration **must** be used only within the
  `ClusterCullingHUAWEI` `Execution` `Model`

- <a href="#VUID-IndexCountHUAWEI-IndexCountHUAWEI-07806"
  id="VUID-IndexCountHUAWEI-IndexCountHUAWEI-07806"></a>
  VUID-IndexCountHUAWEI-IndexCountHUAWEI-07806  
  The variable decorated with `IndexCountHUAWEI` **must** be declared as
  a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#IndexCountHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
