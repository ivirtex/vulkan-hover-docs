# VkAccelerationStructureMotionInstanceTypeNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceTypeNV - Enum specifying a type of
acceleration structure motion instance data for building into an
acceleration structure geometry



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureMotionInstanceTypeNV` enumeration is defined
as:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef enum VkAccelerationStructureMotionInstanceTypeNV {
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV = 0,
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV = 1,
    VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV = 2,
} VkAccelerationStructureMotionInstanceTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV` specifies
  that the instance is a static instance with no instance motion.

- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV`
  specifies that the instance is a motion instance with motion specified
  by interpolation between two matrices.

- `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV`
  specifies that the instance is a motion instance with motion specified
  by interpolation in the SRT decomposition.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMotionInstanceTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
