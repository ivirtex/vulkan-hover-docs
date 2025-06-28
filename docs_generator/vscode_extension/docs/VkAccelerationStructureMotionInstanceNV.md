# VkAccelerationStructureMotionInstanceNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceNV - Structure specifying a single acceleration structure motion instance for building into an acceleration structure geometry



## [](#_c_specification)C Specification

*Acceleration structure motion instances* **can** be built into top-level acceleration structures. Each acceleration structure instance is a separate entry in the top-level acceleration structure which includes all the geometry of a bottom-level acceleration structure at a transformed location including a type of motion and parameters to determine the motion of the instance over time.

An acceleration structure motion instance is defined by the structure:

```c++
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureMotionInstanceNV {
    VkAccelerationStructureMotionInstanceTypeNV     type;
    VkAccelerationStructureMotionInstanceFlagsNV    flags;
    VkAccelerationStructureMotionInstanceDataNV     data;
} VkAccelerationStructureMotionInstanceNV;
```

## [](#_members)Members

- `type` is a [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceTypeNV.html) enumerant identifying which type of motion instance this is and which type of the union is valid.
- `flags` is currently unused, but is required to keep natural alignment of `data`.
- `data` is a [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceDataNV.html) containing motion instance data for this instance.

## [](#_description)Description

Note

If writing this other than with a standard C compiler, note that the final structure should be 152 bytes in size.

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureMotionInstanceNV-type-parameter)VUID-VkAccelerationStructureMotionInstanceNV-type-parameter  
  `type` **must** be a valid [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceTypeNV.html) value
- [](#VUID-VkAccelerationStructureMotionInstanceNV-flags-zerobitmask)VUID-VkAccelerationStructureMotionInstanceNV-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkAccelerationStructureMotionInstanceNV-staticInstance-parameter)VUID-VkAccelerationStructureMotionInstanceNV-staticInstance-parameter  
  If `type` is `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV`, the `staticInstance` member of `data` **must** be a valid [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) structure
- [](#VUID-VkAccelerationStructureMotionInstanceNV-matrixMotionInstance-parameter)VUID-VkAccelerationStructureMotionInstanceNV-matrixMotionInstance-parameter  
  If `type` is `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV`, the `matrixMotionInstance` member of `data` **must** be a valid [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html) structure
- [](#VUID-VkAccelerationStructureMotionInstanceNV-srtMotionInstance-parameter)VUID-VkAccelerationStructureMotionInstanceNV-srtMotionInstance-parameter  
  If `type` is `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV`, the `srtMotionInstance` member of `data` **must** be a valid [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureSRTMotionInstanceNV.html) structure

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html), [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceDataNV.html), [VkAccelerationStructureMotionInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceFlagsNV.html), [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureMotionInstanceNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0