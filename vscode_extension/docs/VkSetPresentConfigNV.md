# VkSetPresentConfigNV(3) Manual Page

## Name

VkSetPresentConfigNV - Structure specifying present metering configuration



## [](#_c_specification)C Specification

Present Metering evenly paces out the next `numFramesPerBatch` [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) presents. This gives smoother pacing between presents in applications with frame generation integrations.

The `VkSetPresentConfigNV` structure is defined as:

```c++
// Provided by VK_NV_present_metering
typedef struct VkSetPresentConfigNV {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           numFramesPerBatch;
    uint32_t           presentConfigFeedback;
} VkSetPresentConfigNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `numFramesPerBatch` is the number of frames to batch
- `presentConfigFeedback` will return the success or error status

## [](#_description)Description

The metering configuration applies to all swapchains in the array in [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html). The configuration specified by `VkSetPresentConfigNV` applies to the next `numFramesPerBatch` calls to [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) and needs to be updated every `numFramesPerBatch` presents.

Valid Usage

- [](#VUID-VkSetPresentConfigNV-numFramesPerBatch-10581)VUID-VkSetPresentConfigNV-numFramesPerBatch-10581  
  `numFramesPerBatch` **must** not be larger than 8

Valid Usage (Implicit)

- [](#VUID-VkSetPresentConfigNV-sType-sType)VUID-VkSetPresentConfigNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SET_PRESENT_CONFIG_NV`

## [](#_see_also)See Also

[VK\_NV\_present\_metering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_present_metering.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSetPresentConfigNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0