# VkLatencySleepModeInfoNV(3) Manual Page

## Name

VkLatencySleepModeInfoNV - Structure to set low latency mode



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkLatencySleepModeInfoNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkLatencySleepModeInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           lowLatencyMode;
    VkBool32           lowLatencyBoost;
    uint32_t           minimumIntervalUs;
} VkLatencySleepModeInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `lowLatencyMode` is the toggle to enable or disable low latency mode.

- `lowLatencyBoost` allows an application to hint to the GPU to increase
  performance to provide additional latency savings at a cost of
  increased power consumption.

- `minimumIntervalUs` is the microseconds between
  [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) calls for a given
  swapchain that [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkLatencySleepNV.html) will enforce.

## <a href="#_description" class="anchor"></a>Description

If `lowLatencyMode` is set to `VK_FALSE`, `lowLatencyBoost` will still
hint to the GPU to increase its power state and `vkLatencySleepNV` will
still enforce `minimumIntervalUs` between `vkQueuePresentKHR` calls.

Valid Usage (Implicit)

- <a href="#VUID-VkLatencySleepModeInfoNV-sType-sType"
  id="VUID-VkLatencySleepModeInfoNV-sType-sType"></a>
  VUID-VkLatencySleepModeInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SLEEP_MODE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencySleepModeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLatencySleepModeInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
