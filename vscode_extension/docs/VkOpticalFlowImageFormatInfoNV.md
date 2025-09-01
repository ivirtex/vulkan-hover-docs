# VkOpticalFlowImageFormatInfoNV(3) Manual Page

## Name

VkOpticalFlowImageFormatInfoNV - Structure describing optical flow image format info



## [](#_c_specification)C Specification

The [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowImageFormatInfoNV {
    VkStructureType              sType;
    const void*                  pNext;
    VkOpticalFlowUsageFlagsNV    usage;
} VkOpticalFlowImageFormatInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`usage` is a bitmask of [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagBitsNV.html) describing the intended optical flow usage of the image.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkOpticalFlowImageFormatInfoNV-sType-sType)VUID-VkOpticalFlowImageFormatInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_INFO_NV`
- [](#VUID-VkOpticalFlowImageFormatInfoNV-usage-parameter)VUID-VkOpticalFlowImageFormatInfoNV-usage-parameter  
  `usage` **must** be a valid combination of [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagBitsNV.html) values
- [](#VUID-VkOpticalFlowImageFormatInfoNV-usage-requiredbitmask)VUID-VkOpticalFlowImageFormatInfoNV-usage-requiredbitmask  
  `usage` **must** not be `0`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowImageFormatInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0