# VkAccelerationStructureGeometrySpheresDataNV(3) Manual Page

## Name

VkAccelerationStructureGeometrySpheresDataNV - Structure specifying a sphere geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

If `VkAccelerationStructureGeometrySpheresDataNV` is included in the `pNext` chain of a [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html) structure, then that structures defines the sphere’s geometry data.

The `VkAccelerationStructureGeometrySpheresDataNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_linear_swept_spheres
typedef struct VkAccelerationStructureGeometrySpheresDataNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkFormat                         vertexFormat;
    VkDeviceOrHostAddressConstKHR    vertexData;
    VkDeviceSize                     vertexStride;
    VkFormat                         radiusFormat;
    VkDeviceOrHostAddressConstKHR    radiusData;
    VkDeviceSize                     radiusStride;
    VkIndexType                      indexType;
    VkDeviceOrHostAddressConstKHR    indexData;
    VkDeviceSize                     indexStride;
} VkAccelerationStructureGeometrySpheresDataNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each sphere’s vertex element.
- `vertexData` is a device or host address to memory containing vertex data in form of pairs of centers of spheres that define all sphere geometry.
- `vertexStride` is the stride in bytes between each vertex element.
- `radiusFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each sphere’s radius.
- `radiusData` is a device or host address to memory containing sphere’s radius data value.
- `radiusStride` is the stride in bytes between each radius value.
- `indexType` is the [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) of each index element.
- `indexData` is `NULL` or a device or host address to memory containing index data for vertex and radius buffers for this geometry.
- `indexStride` is the stride in bytes between each index element.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-None-10429)VUID-VkAccelerationStructureGeometrySpheresDataNV-None-10429  
  The [spheres](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-spheres) feature **must** be enabled
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexData-10430)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexData-10430  
  The memory address in `vertexData` **must** not be `0` or \`NULL'
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexStride-10431)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexStride-10431  
  `vertexStride` **must** be a multiple of:
  
  - the [size of the format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) specified in `vertexFormat` if that format is a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
  - the smallest [component size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) specified in `vertexFormat` if that format is not a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexStride-10432)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexStride-10432  
  `vertexStride` and `radiusStride` **must** be less than or equal to 232-1
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-10433)VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-10433  
  The memory address in `radiusData` **must** not be `0` or \`NULL'
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexFormat-10434)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexFormat-10434  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `vertexFormat` **must** contain `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusFormat-10435)VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusFormat-10435  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `radiusFormat` **must** contain `VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_RADIUS_BUFFER_BIT_NV`
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-10436)VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-10436  
  All values referenced in `radiusData` **must** be greater than or equal to `0`
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-indexData-10437)VUID-VkAccelerationStructureGeometrySpheresDataNV-indexData-10437  
  If `indexData` is not `NULL`, `indexType` **must** be one of `VK_INDEX_TYPE_UINT16` or `VK_INDEX_TYPE_UINT32`

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-sType-sType)VUID-VkAccelerationStructureGeometrySpheresDataNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_SPHERES_DATA_NV`
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexFormat-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexData-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-vertexData-parameter  
  `vertexData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusFormat-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusFormat-parameter  
  `radiusFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-radiusData-parameter  
  `radiusData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-indexType-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-VkAccelerationStructureGeometrySpheresDataNV-indexData-parameter)VUID-VkAccelerationStructureGeometrySpheresDataNV-indexData-parameter  
  `indexData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_linear_swept_spheres.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometrySpheresDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0