# VkAccelerationStructureGeometryKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryKHR - Structure specifying geometries to be built into an acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureGeometryKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    VkGeometryTypeKHR                         geometryType;
    VkAccelerationStructureGeometryDataKHR    geometry;
    VkGeometryFlagsKHR                        flags;
} VkAccelerationStructureGeometryKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `geometryType` describes which type of geometry this `VkAccelerationStructureGeometryKHR` refers to.
- `geometry` is a [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryDataKHR.html) union describing the geometry data for the relevant geometry type.
- `flags` is a bitmask of [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsKHR.html) values describing additional properties of how the geometry should be built.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryKHR-sType-sType)VUID-VkAccelerationStructureGeometryKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR`
- [](#VUID-VkAccelerationStructureGeometryKHR-pNext-pNext)VUID-VkAccelerationStructureGeometryKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX.html), [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html), or [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)
- [](#VUID-VkAccelerationStructureGeometryKHR-sType-unique)VUID-VkAccelerationStructureGeometryKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkAccelerationStructureGeometryKHR-geometryType-parameter)VUID-VkAccelerationStructureGeometryKHR-geometryType-parameter  
  `geometryType` **must** be a valid [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html) value
- [](#VUID-VkAccelerationStructureGeometryKHR-triangles-parameter)VUID-VkAccelerationStructureGeometryKHR-triangles-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the `triangles` member of `geometry` **must** be a valid [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html) structure
- [](#VUID-VkAccelerationStructureGeometryKHR-aabbs-parameter)VUID-VkAccelerationStructureGeometryKHR-aabbs-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_AABBS_KHR`, the `aabbs` member of `geometry` **must** be a valid [VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html) structure
- [](#VUID-VkAccelerationStructureGeometryKHR-instances-parameter)VUID-VkAccelerationStructureGeometryKHR-instances-parameter  
  If `geometryType` is `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the `instances` member of `geometry` **must** be a valid [VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html) structure
- [](#VUID-VkAccelerationStructureGeometryKHR-flags-parameter)VUID-VkAccelerationStructureGeometryKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryDataKHR.html), [VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagsKHR.html), [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0