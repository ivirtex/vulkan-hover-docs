# VkAccelerationStructureMotionInstanceDataNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceDataNV - Union specifying an acceleration structure motion instance data for building into an acceleration structure geometry



## [](#_c_specification)C Specification

Acceleration structure motion instance is defined by the union:

```c++
// Provided by VK_NV_ray_tracing_motion_blur
typedef union VkAccelerationStructureMotionInstanceDataNV {
    VkAccelerationStructureInstanceKHR               staticInstance;
    VkAccelerationStructureMatrixMotionInstanceNV    matrixMotionInstance;
    VkAccelerationStructureSRTMotionInstanceNV       srtMotionInstance;
} VkAccelerationStructureMotionInstanceDataNV;
```

## [](#_members)Members

- `staticInstance` is a [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) structure containing data for a static instance.
- `matrixMotionInstance` is a [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html) structure containing data for a matrix motion instance.
- `srtMotionInstance` is a [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureSRTMotionInstanceNV.html) structure containing data for an SRT motion instance.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html), [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html), [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html), [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html), [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureMotionInstanceDataNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0