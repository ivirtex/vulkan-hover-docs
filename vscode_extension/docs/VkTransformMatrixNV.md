# VkTransformMatrixKHR(3) Manual Page

## Name

VkTransformMatrixKHR - Structure specifying a 3x4 affine transformation matrix



## [](#_c_specification)C Specification

The `VkTransformMatrixKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkTransformMatrixKHR {
    float    matrix[3][4];
} VkTransformMatrixKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkTransformMatrixKHR VkTransformMatrixNV;
```

## [](#_members)Members

- `matrix` is a 3x4 row-major affine transformation matrix.

## [](#_description)Description

Valid Usage

- [](#VUID-VkTransformMatrixKHR-matrix-03799)VUID-VkTransformMatrixKHR-matrix-03799  
  The first three columns of `matrix` **must** define an invertible 3x3 matrix

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html), [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html), [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTransformMatrixKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0