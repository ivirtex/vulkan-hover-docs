# VkCalibratedTimestampInfoKHR(3) Manual Page

## Name

VkCalibratedTimestampInfoKHR - Structure specifying the input parameters of a calibrated timestamp query



## [](#_c_specification)C Specification

The `VkCalibratedTimestampInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_calibrated_timestamps
typedef struct VkCalibratedTimestampInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkTimeDomainKHR    timeDomain;
} VkCalibratedTimestampInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_calibrated_timestamps
typedef VkCalibratedTimestampInfoKHR VkCalibratedTimestampInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `timeDomain` is a [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html) value specifying the time domain from which the calibrated timestamp value should be returned.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCalibratedTimestampInfoKHR-timeDomain-02354)VUID-VkCalibratedTimestampInfoKHR-timeDomain-02354  
  `timeDomain` **must** be one of the [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html) values returned by [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

Valid Usage (Implicit)

- [](#VUID-VkCalibratedTimestampInfoKHR-sType-sType)VUID-VkCalibratedTimestampInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_KHR`
- [](#VUID-VkCalibratedTimestampInfoKHR-pNext-pNext)VUID-VkCalibratedTimestampInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCalibratedTimestampInfoKHR-timeDomain-parameter)VUID-VkCalibratedTimestampInfoKHR-timeDomain-parameter  
  `timeDomain` **must** be a valid [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html) value

## [](#_see_also)See Also

[VK\_EXT\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_calibrated_timestamps.html), [VK\_KHR\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_calibrated_timestamps.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html), [vkGetCalibratedTimestampsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCalibratedTimestampsKHR.html), [vkGetCalibratedTimestampsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCalibratedTimestampsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCalibratedTimestampInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0