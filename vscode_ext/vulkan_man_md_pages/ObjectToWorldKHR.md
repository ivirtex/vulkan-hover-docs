# ObjectToWorldKHR(3) Manual Page

## Name

ObjectToWorldKHR - Transformation matrix from object to world space



## <a href="#_description" class="anchor"></a>Description

`ObjectToWorldKHR`  
A variable decorated with the `ObjectToWorldKHR` decoration will contain
the current object-to-world transformation matrix, which is determined
by the instance of the current intersection.

Valid Usage

- <a href="#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04305"
  id="VUID-ObjectToWorldKHR-ObjectToWorldKHR-04305"></a>
  VUID-ObjectToWorldKHR-ObjectToWorldKHR-04305  
  The `ObjectToWorldKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`

- <a href="#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04306"
  id="VUID-ObjectToWorldKHR-ObjectToWorldKHR-04306"></a>
  VUID-ObjectToWorldKHR-ObjectToWorldKHR-04306  
  The variable decorated with `ObjectToWorldKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04307"
  id="VUID-ObjectToWorldKHR-ObjectToWorldKHR-04307"></a>
  VUID-ObjectToWorldKHR-ObjectToWorldKHR-04307  
  The variable decorated with `ObjectToWorldKHR` **must** be declared as
  a matrix with four columns of three-component vectors of 32-bit
  floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ObjectToWorldKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
