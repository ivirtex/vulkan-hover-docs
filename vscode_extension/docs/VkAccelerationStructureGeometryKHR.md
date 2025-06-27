# VkAccelerationStructureGeometryKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryKHR - Structure specifying geometries to
be built into an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureGeometryKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    VkGeometryTypeKHR                         geometryType;
    VkAccelerationStructureGeometryDataKHR    geometry;
    VkGeometryFlagsKHR                        flags;
} VkAccelerationStructureGeometryKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `geometryType` describes which type of geometry this
  `VkAccelerationStructureGeometryKHR` refers to.

- `geometry` is a
  [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html)
  union describing the geometry data for the relevant geometry type.

- `flags` is a bitmask of
  [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsKHR.html) values describing
  additional properties of how the geometry should be built.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureGeometryKHR-sType-sType"
  id="VUID-VkAccelerationStructureGeometryKHR-sType-sType"></a>
  VUID-VkAccelerationStructureGeometryKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR`

- <a href="#VUID-VkAccelerationStructureGeometryKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureGeometryKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureGeometryKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureGeometryKHR-geometryType-parameter"
  id="VUID-VkAccelerationStructureGeometryKHR-geometryType-parameter"></a>
  VUID-VkAccelerationStructureGeometryKHR-geometryType-parameter  
  `geometryType` **must** be a valid
  [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html) value

- <a href="#VUID-VkAccelerationStructureGeometryKHR-triangles-parameter"
  id="VUID-VkAccelerationStructureGeometryKHR-triangles-parameter"></a>
  VUID-VkAccelerationStructureGeometryKHR-triangles-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the `triangles`
  member of `geometry` **must** be a valid
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)
  structure

- <a href="#VUID-VkAccelerationStructureGeometryKHR-aabbs-parameter"
  id="VUID-VkAccelerationStructureGeometryKHR-aabbs-parameter"></a>
  VUID-VkAccelerationStructureGeometryKHR-aabbs-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_AABBS_KHR`, the `aabbs` member
  of `geometry` **must** be a valid
  [VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html)
  structure

- <a href="#VUID-VkAccelerationStructureGeometryKHR-instances-parameter"
  id="VUID-VkAccelerationStructureGeometryKHR-instances-parameter"></a>
  VUID-VkAccelerationStructureGeometryKHR-instances-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the `instances`
  member of `geometry` **must** be a valid
  [VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)
  structure

- <a href="#VUID-VkAccelerationStructureGeometryKHR-flags-parameter"
  id="VUID-VkAccelerationStructureGeometryKHR-flags-parameter"></a>
  VUID-VkAccelerationStructureGeometryKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsKHR.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html),
[VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html),
[VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagsKHR.html),
[VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureGeometryKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
