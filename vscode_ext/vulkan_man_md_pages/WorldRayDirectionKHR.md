# WorldRayDirectionKHR(3) Manual Page

## Name

WorldRayDirectionKHR - Ray direction in world space



## <a href="#_description" class="anchor"></a>Description

`WorldRayDirectionKHR`  
A variable decorated with the `WorldRayDirectionKHR` decoration will
specify the direction of the ray being processed, in world space. The
value is the parameter passed into the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
target="_blank" rel="noopener">pipeline trace ray</a> instruction.

Valid Usage

- <a href="#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04428"
  id="VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04428"></a>
  VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04428  
  The `WorldRayDirectionKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04429"
  id="VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04429"></a>
  VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04429  
  The variable decorated with `WorldRayDirectionKHR` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04430"
  id="VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04430"></a>
  VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04430  
  The variable decorated with `WorldRayDirectionKHR` **must** be
  declared as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WorldRayDirectionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
