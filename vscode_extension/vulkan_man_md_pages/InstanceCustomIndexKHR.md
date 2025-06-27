# InstanceCustomIndexKHR(3) Manual Page

## Name

InstanceCustomIndexKHR - Custom index associated with an intersected
instance



## <a href="#_description" class="anchor"></a>Description

`InstanceCustomIndexKHR`  
A variable decorated with the `InstanceCustomIndexKHR` decoration will
contain the application-defined value of the instance that intersects
the current ray. This variable contains the value that was specified in
[VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)::`instanceCustomIndex`
for the current acceleration structure instance in the lower 24 bits and
the upper 8 bits will be zero.

Valid Usage

- <a href="#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04251"
  id="VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04251"></a>
  VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04251  
  The `InstanceCustomIndexKHR` decoration **must** be used only within
  the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution`
  `Model`

- <a href="#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04252"
  id="VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04252"></a>
  VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04252  
  The variable decorated with `InstanceCustomIndexKHR` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04253"
  id="VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04253"></a>
  VUID-InstanceCustomIndexKHR-InstanceCustomIndexKHR-04253  
  The variable decorated with `InstanceCustomIndexKHR` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InstanceCustomIndexKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
