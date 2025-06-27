# VkCalibratedTimestampInfoKHR(3) Manual Page

## Name

VkCalibratedTimestampInfoKHR - Structure specifying the input parameters
of a calibrated timestamp query



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCalibratedTimestampInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_calibrated_timestamps
typedef struct VkCalibratedTimestampInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkTimeDomainKHR    timeDomain;
} VkCalibratedTimestampInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_calibrated_timestamps
typedef VkCalibratedTimestampInfoKHR VkCalibratedTimestampInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `timeDomain` is a [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html) value
  specifying the time domain from which the calibrated timestamp value
  should be returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCalibratedTimestampInfoEXT-timeDomain-02354"
  id="VUID-VkCalibratedTimestampInfoEXT-timeDomain-02354"></a>
  VUID-VkCalibratedTimestampInfoEXT-timeDomain-02354  
  `timeDomain` **must** be one of the
  [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html) values returned by
  [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

Valid Usage (Implicit)

- <a href="#VUID-VkCalibratedTimestampInfoKHR-sType-sType"
  id="VUID-VkCalibratedTimestampInfoKHR-sType-sType"></a>
  VUID-VkCalibratedTimestampInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_KHR`

- <a href="#VUID-VkCalibratedTimestampInfoKHR-pNext-pNext"
  id="VUID-VkCalibratedTimestampInfoKHR-pNext-pNext"></a>
  VUID-VkCalibratedTimestampInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCalibratedTimestampInfoKHR-timeDomain-parameter"
  id="VUID-VkCalibratedTimestampInfoKHR-timeDomain-parameter"></a>
  VUID-VkCalibratedTimestampInfoKHR-timeDomain-parameter  
  `timeDomain` **must** be a valid
  [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_calibrated_timestamps.html),
[VK_KHR_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html),
[vkGetCalibratedTimestampsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCalibratedTimestampsEXT.html),
[vkGetCalibratedTimestampsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCalibratedTimestampsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCalibratedTimestampInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
