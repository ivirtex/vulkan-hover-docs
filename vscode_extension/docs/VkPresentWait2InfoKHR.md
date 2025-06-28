# VkPresentWait2InfoKHR(3) Manual Page

## Name

VkPresentWait2InfoKHR - Structure describing parameters of a presentation wait



## [](#_c_specification)C Specification

The `VkPresentWait2InfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_present_wait2
typedef struct VkPresentWait2InfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           presentId;
    uint64_t           timeout;
} VkPresentWait2InfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentId` is the presentation presentId to wait for.
- `timeout` is the timeout period in units of nanoseconds. `timeout` is adjusted to the closest value allowed by the implementation-dependent timeout accuracy, which **may** be substantially longer than one nanosecond, and **may** be longer than the requested period.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPresentWait2InfoKHR-sType-sType)VUID-VkPresentWait2InfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_WAIT_2_INFO_KHR`
- [](#VUID-VkPresentWait2InfoKHR-pNext-pNext)VUID-VkPresentWait2InfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_present\_wait2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_wait2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkWaitForPresent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresent2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentWait2InfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0