# InvocationsPerPixelNV(3) Manual Page

## Name

InvocationsPerPixelNV - Number of fragment shader invocations for the
current pixel



## <a href="#_description" class="anchor"></a>Description

`InvocationsPerPixelNV`  
Decorating a variable with the `InvocationsPerPixelNV` built-in
decoration will make that variable contain the maximum number of
fragment shader invocations per pixel, as derived from the effective
shading rate for the fragment. If a primitive does not fully cover a
pixel, the number of fragment shader invocations for that pixel **may**
be less than the value of `InvocationsPerPixelNV`. If the shading rate
indicates a fragment covering multiple pixels, then
`InvocationsPerPixelNV` will be one.

Valid Usage

- <a href="#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04260"
  id="VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04260"></a>
  VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04260  
  The `InvocationsPerPixelNV` decoration **must** be used only within
  the `Fragment` `Execution` `Model`

- <a href="#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04261"
  id="VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04261"></a>
  VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04261  
  The variable decorated with `InvocationsPerPixelNV` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04262"
  id="VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04262"></a>
  VUID-InvocationsPerPixelNV-InvocationsPerPixelNV-04262  
  The variable decorated with `InvocationsPerPixelNV` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InvocationsPerPixelNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
