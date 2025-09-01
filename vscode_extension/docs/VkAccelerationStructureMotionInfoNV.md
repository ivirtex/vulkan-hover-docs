# VkAccelerationStructureMotionInfoNV(3) Manual Page

## Name

VkAccelerationStructureMotionInfoNV - Structure specifying the parameters of a newly created acceleration structure object



## [](#_c_specification)C Specification

The `VkAccelerationStructureMotionInfoNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureMotionInfoNV {
    VkStructureType                             sType;
    const void*                                 pNext;
    uint32_t                                    maxInstances;
    VkAccelerationStructureMotionInfoFlagsNV    flags;
} VkAccelerationStructureMotionInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxInstances` is the maximum number of instances that **may** be used in the motion top-level acceleration structure.
- `flags` is 0 and reserved for future use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureMotionInfoNV-sType-sType)VUID-VkAccelerationStructureMotionInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MOTION_INFO_NV`
- [](#VUID-VkAccelerationStructureMotionInfoNV-flags-zerobitmask)VUID-VkAccelerationStructureMotionInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html), [VkAccelerationStructureMotionInfoFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInfoFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureMotionInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0