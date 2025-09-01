# VkAccelerationStructureGeometryLinearSweptSpheresDataNV(3) Manual Page

## Name

VkAccelerationStructureGeometryLinearSweptSpheresDataNV - Structure specifying a LSS geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

If `VkAccelerationStructureGeometryLinearSweptSpheresDataNV` is included in the `pNext` chain of a [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html) structure, then that structures defines the linear swept sphere’s (LSS) geometry data.

The `VkAccelerationStructureGeometryLinearSweptSpheresDataNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_linear_swept_spheres
typedef struct VkAccelerationStructureGeometryLinearSweptSpheresDataNV {
    VkStructureType                          sType;
    const void*                              pNext;
    VkFormat                                 vertexFormat;
    VkDeviceOrHostAddressConstKHR            vertexData;
    VkDeviceSize                             vertexStride;
    VkFormat                                 radiusFormat;
    VkDeviceOrHostAddressConstKHR            radiusData;
    VkDeviceSize                             radiusStride;
    VkIndexType                              indexType;
    VkDeviceOrHostAddressConstKHR            indexData;
    VkDeviceSize                             indexStride;
    VkRayTracingLssIndexingModeNV            indexingMode;
    VkRayTracingLssPrimitiveEndCapsModeNV    endCapsMode;
} VkAccelerationStructureGeometryLinearSweptSpheresDataNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each LSS vertex element.
- `vertexData` is a device or host address to memory containing vertex data for this geometry.
- `vertexStride` is the stride in bytes between each vertex element.
- `radiusFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each LSS radius.
- `radiusData` is a device or host address to memory containing LSS radius data value.
- `radiusStride` is the stride in bytes between each radius value.
- `indexType` is the [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) of each index element.
- `indexData` is `NULL` or a device or host address to memory containing index data for vertex and radius buffers for this geometry.
- `indexStride` is the stride in bytes between each index element.
- `indexingMode` is a [VkRayTracingLssIndexingModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssIndexingModeNV.html) value specifying the mode of indexing.
- `endCapsMode` is a [VkRayTracingLssPrimitiveEndCapsModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssPrimitiveEndCapsModeNV.html) value specifying the endcaps mode for LSS primitives.

## [](#_description)Description

If an index buffer is not specified in `indexData`, LSS primitives are rendered individually using subsequent pairs of vertices similar to `VK_PRIMITIVE_TOPOLOGY_LINE_LIST`.

Valid Usage

- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-None-10419)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-None-10419  
  The [linearSweptSpheres](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearSweptSpheres) feature **must** be enabled
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexData-10420)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexData-10420  
  The memory address in `vertexData` **must** not be `0` or \`NULL'
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexStride-10421)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexStride-10421  
  `vertexStride` **must** be a multiple of:
  
  - the [size of the format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) specified in `vertexFormat` if that format is a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
  - the [component size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) specified in `vertexFormat` if that format is not a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexStride-10422)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexStride-10422  
  `vertexStride` and `radiusStride` **must** be less than or equal to 232-1
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexFormat-10423)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexFormat-10423  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `vertexFormat` **must** contain `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusFormat-10424)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusFormat-10424  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `radiusFormat` **must** contain `VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_RADIUS_BUFFER_BIT_NV`
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-10425)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-10425  
  The memory address in `radiusData` **must** not be `0` or \`NULL'
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-10426)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-10426  
  All values referenced in `radiusData` **must** be greater than or equal to `0`
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexingMode-10427)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexingMode-10427  
  If `indexingMode` is `VK_RAY_TRACING_LSS_INDEXING_MODE_SUCCESSIVE_NV`, `indexData` **must** not be `NULL`
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexData-10428)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexData-10428  
  If `indexData` is not `NULL`, `indexType` **must** be one of `VK_INDEX_TYPE_UINT16` or `VK_INDEX_TYPE_UINT32`

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-sType-sType)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_LINEAR_SWEPT_SPHERES_DATA_NV`
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexFormat-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexData-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-vertexData-parameter  
  `vertexData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusFormat-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusFormat-parameter  
  `radiusFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-radiusData-parameter  
  `radiusData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexType-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexData-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexData-parameter  
  `indexData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexingMode-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-indexingMode-parameter  
  `indexingMode` **must** be a valid [VkRayTracingLssIndexingModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssIndexingModeNV.html) value
- [](#VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-endCapsMode-parameter)VUID-VkAccelerationStructureGeometryLinearSweptSpheresDataNV-endCapsMode-parameter  
  `endCapsMode` **must** be a valid [VkRayTracingLssPrimitiveEndCapsModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssPrimitiveEndCapsModeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_linear_swept_spheres.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkRayTracingLssIndexingModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssIndexingModeNV.html), [VkRayTracingLssPrimitiveEndCapsModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssPrimitiveEndCapsModeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryLinearSweptSpheresDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0