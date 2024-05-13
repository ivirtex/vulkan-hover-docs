# VkAccelerationStructureMotionInstanceDataNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceDataNV - Union specifying an
acceleration structure motion instance data for building into an
acceleration structure geometry



## <a href="#_c_specification" class="anchor"></a>C Specification

Acceleration structure motion instance is defined by the union:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef union VkAccelerationStructureMotionInstanceDataNV {
    VkAccelerationStructureInstanceKHR               staticInstance;
    VkAccelerationStructureMatrixMotionInstanceNV    matrixMotionInstance;
    VkAccelerationStructureSRTMotionInstanceNV       srtMotionInstance;
} VkAccelerationStructureMotionInstanceDataNV;
```

## <a href="#_members" class="anchor"></a>Members

- `staticInstance` is a
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  structure containing data for a static instance.

- `matrixMotionInstance` is a
  [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)
  structure containing data for a matrix motion instance.

- `srtMotionInstance` is a
  [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)
  structure containing data for an SRT motion instance.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html),
[VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html),
[VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html),
[VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMotionInstanceDataNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
