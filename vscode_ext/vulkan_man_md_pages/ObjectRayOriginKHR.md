# ObjectRayOriginKHR(3) Manual Page

## Name

ObjectRayOriginKHR - Ray origin in object space



## <a href="#_description" class="anchor"></a>Description

`ObjectRayOriginKHR`  
A variable decorated with the `ObjectRayOriginKHR` decoration will
specify the origin of the ray being processed, in object space.

Valid Usage

- <a href="#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04302"
  id="VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04302"></a>
  VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04302  
  The `ObjectRayOriginKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`

- <a href="#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04303"
  id="VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04303"></a>
  VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04303  
  The variable decorated with `ObjectRayOriginKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04304"
  id="VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04304"></a>
  VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04304  
  The variable decorated with `ObjectRayOriginKHR` **must** be declared
  as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ObjectRayOriginKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
