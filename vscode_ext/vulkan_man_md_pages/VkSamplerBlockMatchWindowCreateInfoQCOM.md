# VkSamplerBlockMatchWindowCreateInfoQCOM(3) Manual Page

## Name

VkSamplerBlockMatchWindowCreateInfoQCOM - Structure specifying the block
match window parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSamplerBlockMatchWindowCreateInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_image_processing2
typedef struct VkSamplerBlockMatchWindowCreateInfoQCOM {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExtent2D                           windowExtent;
    VkBlockMatchWindowCompareModeQCOM    windowCompareMode;
} VkSamplerBlockMatchWindowCreateInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `windowExtent` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) specifying a the
  width and height of the block match window.

- `windowCompareMode` is a
  [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html)
  specifying the compare mode.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-WindowExtent-09210"
  id="VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-WindowExtent-09210"></a>
  VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-WindowExtent-09210  
  `WindowExtent` **must** not be larger than
  [VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html)::`maxBlockMatchWindow`

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-sType-sType"
  id="VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-sType-sType"></a>
  VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_BLOCK_MATCH_WINDOW_CREATE_INFO_QCOM`

- <a
  href="#VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-windowCompareMode-parameter"
  id="VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-windowCompareMode-parameter"></a>
  VUID-VkSamplerBlockMatchWindowCreateInfoQCOM-windowCompareMode-parameter  
  `windowCompareMode` **must** be a valid
  [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing2.html),
[VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerBlockMatchWindowCreateInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
