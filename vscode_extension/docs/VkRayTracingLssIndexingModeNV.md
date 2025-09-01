# VkRayTracingLssIndexingModeNV(3) Manual Page

## Name

VkRayTracingLssIndexingModeNV - LSS indexing mode



## [](#_c_specification)C Specification

Chaining LSS primitives **can** be achieved by specifying an index buffer in [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`indexData` and setting [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`indexingMode` to one of `VkRayTracingLssIndexingModeNV` values:

```c++
// Provided by VK_NV_ray_tracing_linear_swept_spheres
typedef enum VkRayTracingLssIndexingModeNV {
    VK_RAY_TRACING_LSS_INDEXING_MODE_LIST_NV = 0,
    VK_RAY_TRACING_LSS_INDEXING_MODE_SUCCESSIVE_NV = 1,
} VkRayTracingLssIndexingModeNV;
```

## [](#_description)Description

- `VK_RAY_TRACING_LSS_INDEXING_MODE_LIST_NV` specifies that a list of indices is provided where each consecutive pair of indices define a LSS primitive.
- `VK_RAY_TRACING_LSS_INDEXING_MODE_SUCCESSIVE_NV` specifies a successive implicit indexing format, in which each LSS primitive is defined by two successive positions and radii, (k, k + 1), where k is a single index provided in the index buffer. In this indexing scheme, there is a 1:1 mapping between the index buffer and primitive index within the geometry.

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_linear_swept_spheres.html), [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingLssIndexingModeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0