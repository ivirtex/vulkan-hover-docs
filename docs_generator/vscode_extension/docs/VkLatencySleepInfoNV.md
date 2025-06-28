# VkLatencySleepInfoNV(3) Manual Page

## Name

VkLatencySleepInfoNV - Structure specifying the parameters of vkLatencySleepNV



## [](#_c_specification)C Specification

The `VkLatencySleepInfoNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkLatencySleepInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphore        signalSemaphore;
    uint64_t           value;
} VkLatencySleepInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `signalSemaphore` is a semaphore that is signaled to indicate that the application **should** resume input sampling work.
- `value` is the value that `signalSemaphore` is set to for resuming sampling work.

## [](#_description)Description

Valid Usage

- [](#VUID-VkLatencySleepInfoNV-signalSemaphore-09361)VUID-VkLatencySleepInfoNV-signalSemaphore-09361  
  `signalSemaphore` **must** be a timeline semaphore

Valid Usage (Implicit)

- [](#VUID-VkLatencySleepInfoNV-sType-sType)VUID-VkLatencySleepInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SLEEP_INFO_NV`
- [](#VUID-VkLatencySleepInfoNV-signalSemaphore-parameter)VUID-VkLatencySleepInfoNV-signalSemaphore-parameter  
  `signalSemaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkLatencySleepNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencySleepInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0