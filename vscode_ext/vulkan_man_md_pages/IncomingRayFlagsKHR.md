# IncomingRayFlagsKHR(3) Manual Page

## Name

IncomingRayFlagsKHR - Flags used to trace a ray



## <a href="#_description" class="anchor"></a>Description

`IncomingRayFlagsKHR`  
A variable with the `IncomingRayFlagsKHR` decoration will contain the
ray flags passed in to the trace call that invoked this particular
shader. Setting pipeline flags on the ray tracing pipeline **must** not
cause any corresponding flags to be set in variables with this
decoration.

Valid Usage

- <a href="#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04248"
  id="VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04248"></a>
  VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04248  
  The `IncomingRayFlagsKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04249"
  id="VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04249"></a>
  VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04249  
  The variable decorated with `IncomingRayFlagsKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04250"
  id="VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04250"></a>
  VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04250  
  The variable decorated with `IncomingRayFlagsKHR` **must** be declared
  as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#IncomingRayFlagsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
