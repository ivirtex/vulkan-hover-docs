# VkOpticalFlowImageFormatPropertiesNV(3) Manual Page

## Name

VkOpticalFlowImageFormatPropertiesNV - Structure describing properties of an optical flow image format



## [](#_c_specification)C Specification

The [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatPropertiesNV.html) structure is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowImageFormatPropertiesNV {
    VkStructureType    sType;
    const void*        pNext;
    VkFormat           format;
} VkOpticalFlowImageFormatPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that specifies the format that can be used with the specified optical flow image usages.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkOpticalFlowImageFormatPropertiesNV-sType-sType)VUID-VkOpticalFlowImageFormatPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_PROPERTIES_NV`
- [](#VUID-VkOpticalFlowImageFormatPropertiesNV-pNext-pNext)VUID-VkOpticalFlowImageFormatPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowImageFormatPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0