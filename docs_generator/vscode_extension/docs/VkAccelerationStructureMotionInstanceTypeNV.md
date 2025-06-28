# VkAccelerationStructureMotionInstanceTypeNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceTypeNV - Enum specifying a type of acceleration structure motion instance data for building into an acceleration structure geometry



## [](#_c_specification)C Specification

The `VkAccelerationStructureMotionInstanceTypeNV` enumeration is defined as:

```c++
// Provided by VK_NV_ray_tracing_motion_blur
typedef enum VkAccelerationStructureMotionInstanceTypeNV {
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV = 0,
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV = 1,
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV = 2,
} VkAccelerationStructureMotionInstanceTypeNV;
```

## [](#_description)Description

- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV` specifies that the instance is a static instance with no instance motion.
- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV` specifies that the instance is a motion instance with motion specified by interpolation between two matrices.
- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV` specifies that the instance is a motion instance with motion specified by interpolation in the SRT decomposition.

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html), [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureMotionInstanceTypeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0