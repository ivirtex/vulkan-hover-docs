# VkSamplerBlockMatchWindowCreateInfoQCOM(3) Manual Page

## Name

VkSamplerBlockMatchWindowCreateInfoQCOM - Structure specifying the block match window parameters



## [](#_c_specification)C Specification

The `VkSamplerBlockMatchWindowCreateInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_image_processing2
typedef struct VkSamplerBlockMatchWindowCreateInfoQCOM {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExtent2D                           windowExtent;
    VkBlockMatchWindowCompareModeQCOM    windowCompareMode;
} VkSamplerBlockMatchWindowCreateInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `windowExtent` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) specifying a the width and height of the block match window.
- `windowCompareMode` is a [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html) specifying the compare mode.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-WindowExtent-09210)VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-WindowExtent-09210  
  `WindowExtent` **must** not be larger than [VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html)::`maxBlockMatchWindow`

Valid Usage (Implicit)

- [](#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-sType-sType)VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_BLOCK_MATCH_WINDOW_CREATE_INFO_QCOM`
- [](#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-windowCompareMode-parameter)VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-windowCompareMode-parameter  
  `windowCompareMode` **must** be a valid [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html) value

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing2.html), [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerBlockMatchWindowCreateInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0