# VkAccelerationStructureMotionInfoNV(3) Manual Page

## Name

VkAccelerationStructureMotionInfoNV - Structure specifying the
parameters of a newly created acceleration structure object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureMotionInfoNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureMotionInfoNV {
    VkStructureType                             sType;
    const void*                                 pNext;
    uint32_t                                    maxInstances;
    VkAccelerationStructureMotionInfoFlagsNV    flags;
} VkAccelerationStructureMotionInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxInstances` is the maximum number of instances that **may** be used
  in the motion top-level acceleration structure.

- `flags` is 0 and reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureMotionInfoNV-sType-sType"
  id="VUID-VkAccelerationStructureMotionInfoNV-sType-sType"></a>
  VUID-VkAccelerationStructureMotionInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MOTION_INFO_NV`

- <a href="#VUID-VkAccelerationStructureMotionInfoNV-flags-zerobitmask"
  id="VUID-VkAccelerationStructureMotionInfoNV-flags-zerobitmask"></a>
  VUID-VkAccelerationStructureMotionInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureMotionInfoFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMotionInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
