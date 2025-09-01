# VkLatencySleepModeInfoNV(3) Manual Page

## Name

VkLatencySleepModeInfoNV - Structure to set low latency mode



## [](#_c_specification)C Specification

The `VkLatencySleepModeInfoNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkLatencySleepModeInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           lowLatencyMode;
    VkBool32           lowLatencyBoost;
    uint32_t           minimumIntervalUs;
} VkLatencySleepModeInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `lowLatencyMode` is the toggle to enable or disable low latency mode.
- `lowLatencyBoost` allows an application to hint to the GPU to increase performance to provide additional latency savings at a cost of increased power consumption.
- `minimumIntervalUs` is the microseconds between [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) calls for a given swapchain that [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkLatencySleepNV.html) will enforce.

## [](#_description)Description

If `lowLatencyMode` is `VK_FALSE`, `lowLatencyBoost` will still hint to the GPU to increase its power state and `vkLatencySleepNV` will still enforce `minimumIntervalUs` between `vkQueuePresentKHR` calls.

Valid Usage (Implicit)

- [](#VUID-VkLatencySleepModeInfoNV-sType-sType)VUID-VkLatencySleepModeInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SLEEP_MODE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencySleepModeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencySleepModeInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0