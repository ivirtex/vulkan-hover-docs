# ObjectRayDirectionKHR(3) Manual Page

## Name

ObjectRayDirectionKHR - Ray direction in object space



## <a href="#_description" class="anchor"></a>Description

`ObjectRayDirectionKHR`  
A variable decorated with the `ObjectRayDirectionKHR` decoration will
specify the direction of the ray being processed, in object space.

Valid Usage

- <a href="#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04299"
  id="VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04299"></a>
  VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04299  
  The `ObjectRayDirectionKHR` decoration **must** be used only within
  the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution`
  `Model`

- <a href="#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04300"
  id="VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04300"></a>
  VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04300  
  The variable decorated with `ObjectRayDirectionKHR` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04301"
  id="VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04301"></a>
  VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04301  
  The variable decorated with `ObjectRayDirectionKHR` **must** be
  declared as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ObjectRayDirectionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
